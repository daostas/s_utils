package s_utils

import (
	"fmt"
	"reflect"
	"strings"
)

// ToStringArray Принимает любой массив, может принимать указатели на массив.
// Собирает все элементы массива в массив строк
func ToStringArray[T any](arr []T) (res []string) {
	for _, value := range arr {
		v := reflect.Indirect(reflect.ValueOf(value))
		res = append(res, fmt.Sprint(v))
	}
	return res
}

// ToStringArrayF Принимает любой массив, может принимать указатели на массив.
// Собирает все элементы массива в массив строк
func ToStringArrayF[T any](arr []T, format string) (res []string) {
	for _, value := range arr {
		v := reflect.Indirect(reflect.ValueOf(value))
		res = append(res, fmt.Sprintf(format, v))
	}
	return res
}

// ToStringArrayByTagName Принимает любой массив структур, может принимать указатели на массив.
// Ищет в структуре поля с указанным json тегом и собирает их в массив строк
func ToStringArrayByTagName[T any](arr []T, tagName string) (res []string) {
	if len(arr) != 0 {
		v := reflect.Indirect(reflect.ValueOf(arr[0]))
		pos := 0
		for i := 0; i < v.NumField(); i++ {
			if tagName == strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0] {
				pos = i
			}
		}
		for _, value := range arr {
			v = reflect.Indirect(reflect.ValueOf(value))
			res = append(res, fmt.Sprint(v.Field(pos)))
		}
		return
	}
	return
}

// ToStringArrayByTagNameF Принимает любой массив структур, может принимать указатели на массив.
// Ищет в структуре поля с указанным json тегом и собирает их в массив строк
func ToStringArrayByTagNameF[T any](arr []T, tagName string, format string) (res []string) {
	if len(arr) != 0 {
		v := reflect.Indirect(reflect.ValueOf(arr[0]))
		pos := 0
		for i := 0; i < v.NumField(); i++ {
			if tagName == strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0] {
				pos = i
			}
		}
		for _, value := range arr {
			v = reflect.Indirect(reflect.ValueOf(value))
			res = append(res, fmt.Sprintf(format, v.Field(pos)))
		}
		return
	}
	return
}
