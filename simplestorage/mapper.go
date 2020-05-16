package simplestorage

import "errors"

//Mapper hold data map with key is string and data is general interface
type Mapper struct {
	data map[string]interface{}
}

//NewMapper instance
func NewMapper() *Mapper {
	return &Mapper{
		data: make(map[string]interface{}),
	}
}

//Put some data with key reference
func (m *Mapper) Put(key string, data interface{}) error {
	if len(key) <= 0 {
		return errors.New("invalid key")
	}
	m.data[key] = data
	return nil
}

//Get data by key that had set before with Put func
func (m *Mapper) Get(key string) interface{} {

	if len(key) <= 0 {
		return errors.New("invalid key")
	}

	//data not set before
	if m.data[key] == nil {
		return errors.New("empty data")
	}

	return m.data[key]
}
