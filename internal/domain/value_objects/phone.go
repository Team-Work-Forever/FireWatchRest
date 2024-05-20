package valueobjects

type Phone struct {
	PhoneCode   string `gorm:"column:phone_code"`
	PhoneNumber string `gorm:"column:phone_number"`
}
