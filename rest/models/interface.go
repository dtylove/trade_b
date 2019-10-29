package models

import "github.com/pkg/errors"

func FindProductList() (pList []Product, error error) {
	error = FindList(&pList)
	return
}

func FindById(model interface{}, id uint) error {
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

	return GetDB().Find(model, id).Error
}

func FindList(model interface{}) error {
	switch model.(type) {
	case *[]Product:
		model = model.(*[]Product)
	case *User:
		model = model.(*User)
	case *Order:
		model = model.(*Order)
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