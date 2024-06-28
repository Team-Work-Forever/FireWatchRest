package usecases

import (
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/daos"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/repositories"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/date"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/services"
	"github.com/Team-Work-Forever/FireWatchRest/internal/infrastructure/upload"
	"github.com/Team-Work-Forever/FireWatchRest/pkg/contracts"
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type CreateBurnUseCase struct {
	burnRepo     *repositories.BurnRepository
	autarchyRepo *repositories.AutarchyRepository
	fileService  *upload.BlobService
}

func NewCreateBurnUseCase(
	burnRepo *repositories.BurnRepository,
	autarchyRepo *repositories.AutarchyRepository,
	fileService *upload.BlobService,
) *CreateBurnUseCase {
	return &CreateBurnUseCase{
		burnRepo:     burnRepo,
		autarchyRepo: autarchyRepo,
		fileService:  fileService,
	}
}

func (uc *CreateBurnUseCase) Handler(request contracts.CreateBurnRequest) (*contracts.BurnActionResponse, error) {
	var state vo.BurnRequestStates = vo.Scheduled

	address, err := services.GetAddress(request.Lat, request.Lon)

	if err != nil {
		return nil, err
	}

	initDate, err := date.ParseString(request.InitDate)

	if err != nil {
		return nil, err
	}

	dateNow, err := date.Now()

	if err != nil {
		return nil, err
	}

	if *initDate == *dateNow {
		state = vo.Active
	}

	reason, ok := vo.GetBurnReasonKey(request.Reason)

	if !ok {
		return nil, exec.BURN_PROVIDE_NOT_EXISTING_REASON
	}

	burnType, ok := vo.GetBurnTypeKey(request.Type)

	if !ok {
		return nil, exec.BURN_PROVIDE_NOT_EXISTING_TYPE
	}

	foundAutarchy, err := uc.autarchyRepo.GetAutarchyByCity(address.City)

	if err != nil {
		return nil, err
	}

	if ok := services.CheckICFNIndex(request.Lat, request.Lon, request.HasBackUpTeam, request.Ignore); !ok {
		state = vo.Rejected
	}

	burn, err := entities.NewBurn(
		request.Title,
		reason,
		burnType,
		*address,
		*vo.NewCoordinate(request.Lat, request.Lon),
		*initDate,
	)

	if err != nil {
		return nil, err
	}

	if request.MapPicture != nil {
		file, err := request.MapPicture.Open()

		if err != nil {
			return nil, err
		}

		defer file.Close()
		url, err := uc.fileService.UploadFile(&upload.UploadFile{
			Bucket:   upload.ClientBucket,
			FileName: request.MapPicture.Filename,
			FileId:   burn.GetId(),
			FileBody: file,
		})

		if err != nil {
			return nil, err
		}

		burn.SetMapPicture(url)
	}

	burnRequest, err := uc.burnRepo.CreateBurn(daos.CreateBurnDao{
		AuthId:         request.UserId,
		AutarchyId:     foundAutarchy.ID,
		Burn:           burn,
		InitialPropose: request.InitialProprose,
		State:          state,
	})

	if err != nil {
		return nil, err
	}

	return &contracts.BurnActionResponse{
		BurnId: burnRequest.BurnId,
		State:  burnRequest.State.GetState(),
	}, nil
}
