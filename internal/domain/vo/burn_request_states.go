package vo

type BurnRequestStates int

const (
	Scheduled BurnRequestStates = iota
	Active
	Completed
	Rejected
)

var burnRequestStateLookUp = map[uint16]string{
	0: "scheduled",
	1: "active",
	2: "completed",
	3: "rejected",
}

func GetBurnRequestState(burnRequestState uint16) (string, bool) {
	result, ok := burnRequestStateLookUp[uint16(burnRequestState)]
	return result, ok
}

func GetAllBurnStates() []string {
	states := make([]string, 0, len(burnRequestStateLookUp))

	for _, v := range burnRequestStateLookUp {
		states = append(states, v)
	}

	return states
}

func GetBurnRequestStateKey(value string) (uint16, bool) {
	for k, v := range burnRequestStateLookUp {
		if v == value {
			return k, true
		}
	}

	return 0, false
}
