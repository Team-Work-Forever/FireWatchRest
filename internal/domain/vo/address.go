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

func checkStreetNumber(number int) error {
	if number == 0 {
		return exec.ADDRESS_PROVIDE_NUMBER
	}

	if number < 0 {
		return exec.ADDRESS_PROVIDE_AN_VALID_NUMBER
	}

	return nil
}

func NewAddressWithEmptyValues(
	street string,
	number int,
	zipCode ZipCode,
	city string,
) (*Address, error) {
	if street == "" {
		street = "not defined"
	}

	if err := checkStreetNumber(number); err != nil {
		return nil, err
	}

	if city == "" {
		city = "not defined"
	}

	return &Address{
		Street:  street,
		Number:  number,
		ZipCode: zipCode.GetValue(),
		City:    city,
	}, nil
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

	if err := checkStreetNumber(number); err != nil {
		return nil, err
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

func (a *Address) SetStreetNumber(number int) error {
	if err := checkStreetNumber(number); err != nil {
		return err
	}

	a.Number = number
	return nil
}
