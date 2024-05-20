package vo

import "errors"

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
		return nil, errors.New("provide an street")
	}

	if number == 0 {
		return nil, errors.New("provide an number")
	}

	if number < 0 {
		return nil, errors.New("provide an valid number")
	}

	if city == "" {
		return nil, errors.New("provide an city")
	}

	return &Address{
		Street:  street,
		Number:  number,
		ZipCode: zipCode.GetValue(),
		City:    city,
	}, nil
}
