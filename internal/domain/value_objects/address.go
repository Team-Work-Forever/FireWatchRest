package valueobjects

type Address struct {
	Street  string `gorm:"column:address_street"`
	Number  int    `gorm:"column:address_number;type:int"`
	ZipCode string `gorm:"column:address_zip_code"`
	City    string `gorm:"column:address_city"`
}
