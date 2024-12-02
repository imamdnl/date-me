package value_object

import (
	"date-me/domain/constants"
	"errors"
)

type UserGender struct {
	value uint
}

func NewUserGender(value uint) (*UserGender, error) {
	dataValueObject := &UserGender{value: value}

	if dataValueObject.IsMan() {
		return dataValueObject, nil
	}

	if dataValueObject.IsWoman() {
		return dataValueObject, nil
	}

	if dataValueObject.IsEmpty() {
		return dataValueObject, nil
	}

	return nil, errors.New(constants.USER_GENDER_NOT_FOUND)
}

func NewUserGenderByString(value string) (*UserGender, error) {
	if value == constants.ManString {
		return &UserGender{value: constants.Man}, nil
	}

	if value == constants.WomanString {
		return &UserGender{value: constants.Woman}, nil
	}

	if value == "" {
		return &UserGender{value: 0}, nil
	}

	return nil, errors.New(constants.USER_GENDER_NOT_FOUND)
}

func (value *UserGender) GetStringValue() string {
	var valueString string
	switch {
	case value.IsMan():
		valueString = constants.ManString
	case value.IsWoman():
		valueString = constants.WomanString
	case value.IsEmpty():
		valueString = ""
	}
	return valueString
}

func (value *UserGender) GetValue() uint {
	return value.value
}

func (value *UserGender) IsEmpty() bool {
	return value.value == 0
}

func (value *UserGender) IsMan() bool {
	return value.value == constants.Man
}

func (value *UserGender) IsWoman() bool {
	return value.value == constants.Woman
}
