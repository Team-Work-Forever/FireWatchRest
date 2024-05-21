package vo

import (
	exec "github.com/Team-Work-Forever/FireWatchRest/pkg/exceptions"
)

type Address struct {
	Street  string `gorm:"column:address_street"`
	Number  int    `gorm:"column:address_number;type:int"`
	ZipCode string `gorm:"column:address_zip_code"`
	City    string `gorm:"column:address_city"`
}

func NewAddress(
	street string,
	number int,
	zipCode ZipCode,
	city string,
) (*Address, error) {
	if street == "" {
		return nil, exec.ADDRESS_PROVIDE_STREET
	}

	if number == 0 {
		return nil, exec.ADDRESS_PROVIDE_NUMBER
	}

	if number < 0 {
		return nil, exec.ADDRESS_PROVIDE_AN_VALID_NUMBER
	}

	if city == "" {
		return nil, exec.ADDRESS_PROVIDE_CITY
	}

	return &Address{
		Street:  street,
		Number:  number,
		ZipCode: zipCode.GetValue(),
		City:    city,
	}, nil
}
