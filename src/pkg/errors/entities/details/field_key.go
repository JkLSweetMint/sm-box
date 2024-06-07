package details

import (
	"fmt"
	"sm-box/pkg/errors/types"
	"strings"
)

type (
	// FieldKey - ключ поля.
	FieldKey struct {
		path []string
	}
)

// init - инициализация ключа поля.
func (fk *FieldKey) init() {
	if fk.path == nil {
		fk.path = make([]string, 0)
	}
}

// Add - добавления элемента в путь к ключу.
func (fk *FieldKey) Add(path ...string) types.DetailsFieldKey {
	fk.init()

	fk.path = append(fk.path, path...)
	return fk
}

// AddArray - добавления массива в путь к ключу.
func (fk *FieldKey) AddArray(name string, index int) types.DetailsFieldKey {
	fk.init()

	fk.path = append(fk.path, fmt.Sprintf("%s[%d]", name, index))
	return fk
}

// AddMap - добавления карты в путь к ключу.
func (fk *FieldKey) AddMap(name string, key any) types.DetailsFieldKey {
	fk.init()

	fk.path = append(fk.path, fmt.Sprintf("%s[%v]", name, key))
	return fk
}

// String - получение строкового представления путя к ключу.
func (fk *FieldKey) String() (str string) {
	return strings.Join(fk.path, ".")
}

// Clone - копирование ключа.
func (fk *FieldKey) Clone() types.DetailsFieldKey {
	fk.init()

	var fk_ = new(FieldKey)
	fk_.init()

	for _, p := range fk.path {
		fk_.path = append(fk_.path, strings.Clone(p))
	}

	return fk_
}
