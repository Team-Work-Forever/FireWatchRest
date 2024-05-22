package vo

type BurnReason int

const (
	SanitaryBurn BurnReason = iota
	AgritoralWasteManagement
	ForestryWasteManagement
	BushManagement
	Others
)

var burnReasonLookup = map[uint16]string{
	0: "sanitaryBurn",
	1: "agritoralWasteManagement",
	2: "forestryWasteManagement",
	3: "bushManagement",
	4: "others",
}

func GetBurnReason(burnReason int) (string, bool) {
	result, ok := burnReasonLookup[uint16(burnReason)]
	return result, ok
}

func GetAllBurnReasons() []string {
	reasons := make([]string, 0, len(burnReasonLookup))

	for _, v := range burnReasonLookup {
		reasons = append(reasons, v)
	}

	return reasons
}

func GetBurnReasonKey(value string) (uint16, bool) {
	for k, v := range burnReasonLookup {
		if v == value {
			return k, true
		}
	}
	return 0, false
}
