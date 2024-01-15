package util

import (
	"errors"

	"github.com/HectorMu/go-rest-api/types"
	"github.com/go-playground/validator/v10"
)

func GetNumber(withErr bool) (int8, error) {

	if withErr {
		return 0, errors.New("what the fuck")
	}

	return 100, nil
}

func FindIndex[T any](list []T, condition func(T) bool) int {
	for i, item := range list {
		if condition(item) {
			return i
		}
	}
	return -1
}

func FilterSlice[T any](list []T, condition func(T) bool) []T {
	tempSlice := []T{}

	for _, item := range list {
		if condition(item) {
			tempSlice = append(tempSlice, item)
		}
	}

	return tempSlice
}

func MapSlice[T, U any](list []T, mapper func(T) U) []U {
	tempSlice := make([]U, len(list))

	for i, item := range list {
		mappedItem := mapper(item)
		tempSlice[i] = mappedItem
	}

	return tempSlice
}

func EverySlice[T any](list []T, condition func(T) bool) bool {
	tempResult := true

	for _, item := range list {
		if condition(item) {
			continue
		} else {
			tempResult = false
			break
		}
	}

	return tempResult
}

var validate = validator.New()

func ValidateUser(user types.User) error {
	if err := validate.Struct(user); err != nil {
		return err
	}
	return nil
}
