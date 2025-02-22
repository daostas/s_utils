package s_utils

import (
	"fmt"
	"strings"
)

type List[T any] []T

// Add Добавление в массив значений любых структур, по дефолту unique = true
func (l *List[T]) Add(value T, unique ...bool) {
	if len(*l) == 0 {
		*l = append(*l, value)
		return
	}

	if len(unique) != 0 {
		if unique[0] {
			for _, v := range *l {
				if fmt.Sprint(v) == fmt.Sprint(value) {
					return
				}
			}
		}
	} else {
		for _, v := range *l {
			if fmt.Sprint(v) == fmt.Sprint(value) {
				return
			}
		}
	}

	*l = append(*l, value)
}

func (l *List[T]) AddArray(values []T, unique ...bool) {
	for _, v := range values {
		l.Add(v, unique...)
	}
}

type ListString []string

func (list *ListString) Add(str string, unique bool) {

	if len(*list) == 0 {
		*list = append(*list, str)
		return
	}

	if unique {
		for _, l := range *list {
			if strings.EqualFold(l, str) {
				return
			}
		}
	}

	*list = append(*list, str)
}

func (list *ListString) AddSlice(arr []string, unique bool) {
	for _, v := range arr {
		list.Add(v, unique)
	}
}
