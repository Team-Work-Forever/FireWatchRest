package vo

type UserType int

const (
	User UserType = iota
	Autarchy
	Admin
)

var userTypeLookup = map[uint16]string{
	0: "user",
	1: "autarchy",
	2: "admin",
}

func GetUserType(userType int) (string, bool) {
	result, ok := userTypeLookup[uint16(userType)]
	return result, ok
}
