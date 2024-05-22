package vo

type BurnType int

const (
	Burn BurnType = iota
)

var burnTypeLookup = map[uint16]string{
	0: "burn",
}

func GetBurnType(burnType int) (string, bool) {
	result, ok := burnTypeLookup[uint16(burnType)]
	return result, ok
}

func GetBurnTypeKey(value string) (uint16, bool) {
	for k, v := range burnTypeLookup {
		if v == value {
			return k, true
		}
	}
	return 0, false
}
