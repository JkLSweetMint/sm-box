package details

import (
	"encoding/json"
	"sm-box/pkg/errors/types"
	"sync"
)

type (
	// Details - детали для ошибок.
	Details struct {
		fields Fields

		storage map[string]any
		rwMux   *sync.RWMutex
	}

	Fields []*types.DetailsField
)

// init - инициализация хранилища.
func (ds *Details) init() {
	if ds.storage == nil {
		ds.storage = make(map[string]any)
	}

	if ds.rwMux == nil {
		ds.rwMux = new(sync.RWMutex)
	}

	if ds.fields == nil {
		ds.fields = make([]*types.DetailsField, 0)
	}

	return
}

// Peek - получение данных из хранилища.
func (ds *Details) Peek(k string) (v any) {
	ds.init()

	ds.rwMux.RLock()
	defer ds.rwMux.RUnlock()

	return ds.storage[k]
}

// Set - установить значение в хранилище по ключу.
func (ds *Details) Set(k string, v any) types.Details {
	ds.init()

	ds.rwMux.Lock()
	defer ds.rwMux.Unlock()

	ds.storage[k] = v

	return ds
}

// Reset - сбросить детали.
func (ds *Details) Reset() types.Details {
	ds.init()

	ds.storage = make(map[string]any)

	return ds
}

// PeekFieldMessage - получить сообщение поля.
func (ds *Details) PeekFieldMessage(k string) (m types.DetailsFieldMessage) {
	ds.init()

	for _, f := range ds.fields {
		if k == f.Key.String() {
			m = f.Message
			return
		}
	}

	return
}

// SetField - установить значение поля ошибки.
// В случае пересечения ключей поля будет вставлено новое значение.
func (ds *Details) SetField(k types.DetailsFieldKey, m types.DetailsFieldMessage) types.Details {
	ds.init()

	var f = &types.DetailsField{
		Key:     k,
		Message: m,
	}

	var found bool

	for _, field := range ds.fields {
		if f.Key == field.Key {
			field.Message = f.Message
			found = true
			break
		}
	}

	if !found {
		ds.fields = append(ds.fields, f)
	}

	return ds
}

// SetFields - установить значение полей ошибки.
// В случае пересечения ключей поля будет вставлено новое значение.
func (ds *Details) SetFields(fields ...types.DetailsField) types.Details {
	ds.init()

For:
	for _, newField := range fields {
		for _, field := range ds.fields {
			if newField.Key == field.Key {
				field.Message = newField.Message
				continue For
			}
		}

		ds.fields = append(ds.fields, &newField)
	}

	return ds
}

// ResetFields - сбросить поля.
func (ds *Details) ResetFields() types.Details {
	ds.init()

	ds.fields = make([]*types.DetailsField, 0)

	return ds
}

// Clone - копирование деталей ошибки.
func (ds *Details) Clone() types.Details {
	ds.init()

	var ds_ = &Details{
		fields:  make([]*types.DetailsField, 0),
		storage: make(map[string]any),
		rwMux:   new(sync.RWMutex),
	}

	// storage
	{
		if len(ds.storage) > 0 {
			if data, err := json.Marshal(ds.storage); err == nil {
				if err = json.Unmarshal(data, &ds_.storage); err != nil {
					ds_.storage = make(map[string]any)
				}
			}
		}
	}

	// fields
	{
		for _, f := range ds.fields {
			ds_.SetField(f.Key.Clone(), f.Message.Clone())
		}
	}

	ds_.init()

	return ds_
}
