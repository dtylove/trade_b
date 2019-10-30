package models

import "github.com/pkg/errors"

func FindList(model interface{}) error {
	switch model.(type) {
	case *[]Product:
		model = model.(*[]Product)
	case *[]User:
		model = model.(*[]User)
	case *[]Order:
		model = model.(*[]Order)
	default:
		return errors.New("invalid model type")
	}

	err := GetDB().Find(model).Error
	return err
}

func Add(model interface{}) error {
	switch model.(type) {
	case *Product:
		model = model.(*Product)
	case *User:
		model = model.(*User)
	case *Order:
		model = model.(*Order)
	default:
		return errors.New("invalid model type")
	}

	err := GetDB().Create(model).Error
	return err
}

func FindOne(model interface{}) error {
	switch model.(type) {
	case *Product:
		model = model.(*Product)
	case *User:
		model = model.(*User)
	case *Order:
		model = model.(*Order)
	default:
		return errors.New("invalid model type")
	}

	return GetDB().Find(model).Error
}
