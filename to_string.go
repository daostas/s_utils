package s_utils

import (
	"fmt"
	"reflect"
	"strings"
)

const DefaultSeparator = ", "

func getSeparator(separator ...string) string {
	if len(separator) != 1 {
		return DefaultSeparator
	} else {
		return separator[0]
	}
}

// ToString Принимает любой массив, может принимать указатели на массив.
// Собирает все элементы массива в строку с разделителем `, `
func ToString[T any](arr []T, separator ...string) (str string) {
	sep := getSeparator(separator...)
	for i, value := range arr {
		v := reflect.Indirect(reflect.ValueOf(value))
		str += fmt.Sprint(v)
		if i < len(arr)-1 {
			str += sep
		}
	}
	return
}

// ToStringF Принимает любой массив, может принимать указатели на массив.
// Собирает все элементы массива в строку с разделителем `, `
func ToStringF[T any](arr []T, format string, separator ...string) (str string) {
	sep := getSeparator(separator...)
	for i, value := range arr {
		v := reflect.Indirect(reflect.ValueOf(value))
		str += fmt.Sprintf(format, v)
		if i < len(arr)-1 {
			str += sep
		}
	}
	return
}

// ToStringByTagName Принимает любой массив структур, может принимать указатели на массив.
// Ищет в структуре поля с указанным json тегом и собирает их в строку с раделителем `, `
func ToStringByTagName[T any](arr []T, tagName string, separator ...string) (str string) {
	sep := getSeparator(separator...)
	if len(arr) != 0 {
		v := reflect.Indirect(reflect.ValueOf(arr[0]))
		pos := 0
		for i := 0; i < v.NumField(); i++ {
			if tagName == strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0] {
				pos = i
				break
			}
		}

		for i, value := range arr {
			v = reflect.Indirect(reflect.ValueOf(value))
			str += fmt.Sprint(v.Field(pos))
			if i < len(arr)-1 {
				str += sep
			}
		}
		return
	}
	return
}

// ToStringByTagNameF Принимает любой массив структур, может принимать указатели на массив.
// Ищет в структуре поля с указанным json тегом и собирает их в строку с раделителем `, `
func ToStringByTagNameF[T any](arr []T, tagName string, format string, separator ...string) (str string) {
	sep := getSeparator(separator...)
	if len(arr) != 0 {
		v := reflect.Indirect(reflect.ValueOf(arr[0]))
		pos := 0
		for i := 0; i < v.NumField(); i++ {
			if tagName == strings.Split(v.Type().Field(i).Tag.Get("json"), ",")[0] {
				pos = i
				break
			}
		}

		for i, value := range arr {
			v = reflect.Indirect(reflect.ValueOf(value))
			str += fmt.Sprintf(format, v.Field(pos))
			if i < len(arr)-1 {
				str += sep
			}
		}
		return
	}
	return
}
