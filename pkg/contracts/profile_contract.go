package contracts

import (
	"errors"
	"mime/multipart"

	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/entities"
	"github.com/Team-Work-Forever/FireWatchRest/internal/domain/vo"
)

var (
	ErrCannotConvert = errors.New("cannot Convert")
)

type (
	ProfileResponse struct {
		Id       string          `json:"id"`
		Email    string          `json:"email"`
		NIF      string          `json:"nif"`
		Avatar   string          `json:"avatar"`
		Phone    PhoneResponse   `json:"phone"`
		Address  AddressResponse `json:"address"`
		UserType string          `json:"user_type"`
	}

	Profile interface {
		GetProfile() ProfileResponse
	}

	UserProfileResponse struct {
		ProfileResponse
		UserName  string `json:"user_name"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	AutarchyProfileResponse struct {
		ProfileResponse
		Title string `json:"title"`
	}

	PublicProfileResponse struct {
		Email    string        `json:"email"`
		UserName string        `json:"user_name"`
		Avatar   string        `json:"avatar"`
		NIF      string        `json:"nif"`
		Phone    PhoneResponse `json:"phone"`
	}

	WhoamiRequest struct {
		UserId string
	}

	PublicProfileRequest struct {
		Email string
	}

	UpdateProfileResponse struct {
		UserId      string                `swaggerignore:"true"`
		Email       string                `form:"email" binding:"required"`
		UserName    string                `form:"user_name" binding:"required"`
		PhoneCode   string                `form:"phone_code" binding:"required"`
		PhoneNumber string                `form:"phone_number" binding:"required"`
		Street      string                `form:"street" binding:"required"`
		StreetPort  *int                  `form:"street_port" binding:"required"`
		ZipCode     string                `form:"zip_code" binding:"required"`
		City        string                `form:"city" binding:"required"`
		Avatar      *multipart.FileHeader `form:"avatar" binding:"required" swaggerignore:"true"`
	}
)

func GetProfileResponse(auth *entities.Auth, user interface{}) (interface{}, error) {
	switch auth.UserType {
	case int(vo.User), int(vo.Admin):
		userProfile, ok := user.(*entities.User)

		if !ok {
			return nil, ErrCannotConvert
		}

		return &UserProfileResponse{
			ProfileResponse: createProfileResponse(auth, &userProfile.IdentityUser),
			UserName:        userProfile.UserName,
			FirstName:       userProfile.FirstName,
			LastName:        userProfile.LastName,
		}, nil

	case int(vo.Autarchy):
		autarchyProfile, ok := user.(*entities.Autarchy)

		if !ok {
			return nil, ErrCannotConvert
		}

		return AutarchyProfileResponse{
			ProfileResponse: createProfileResponse(auth, &autarchyProfile.IdentityUser),
			Title:           autarchyProfile.Title,
		}, nil
	}

	return nil, ErrCannotConvert
}

func createProfileResponse(auth *entities.Auth, identity *entities.IdentityUser) ProfileResponse {
	return ProfileResponse{
		Id:     auth.ID,
		Email:  auth.Email.GetValue(),
		NIF:    auth.NIF.GetValue(),
		Avatar: identity.Picture,
		Phone: PhoneResponse{
			CountryCode: identity.PhoneNumber.CountryCode,
			Number:      identity.PhoneNumber.Number,
		},
		Address: AddressResponse{
			Street: identity.Address.Street,
			Number: identity.Address.Number,
			ZipCode: ZipCodeResponse{
				Value: identity.Address.ZipCode,
			},
			City: identity.Address.City,
		},
		UserType: auth.GetRole(),
	}
}
