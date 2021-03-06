package simplestorage

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

//Mapper hold data map with key is string and data is general interface
type Mapper struct {
	data       map[string]interface{}
	IsPrintLog bool
	lock       sync.RWMutex
}

//NewMapper instance
func NewMapper() *Mapper {
	return &Mapper{
		data:       make(map[string]interface{}),
		IsPrintLog: false,
	}
}

//Put some data with key reference
func (m *Mapper) Put(key string, data interface{}) error {
	if len(key) <= 0 {
		return errors.New("invalid key")
	}

	//need lock for concurrency; CONCERN: try use go channel
	m.lock.Lock()
	m.data[key] = data
	m.lock.Unlock()

	return nil
}

//Get data by key that had set before with Put func
func (m *Mapper) Get(key string) interface{} {

	if len(key) <= 0 {
		return errors.New("invalid key")
	}

	if m.IsPrintLog {
		log.Println(key, " : ", m.data[key])
	}

	//data not set before
	if m.data[key] == nil {
		return errors.New("empty data")
	}

	return m.data[key]
}

//PrintData show standard fmt output for current data
func (m *Mapper) PrintData() {
	fmt.Printf("data : %+v", m.data)
}

//GetRefData get reference of data; WARNING can manipulate data externally
func (m *Mapper) GetRefData() *map[string]interface{} {
	return &m.data
}
