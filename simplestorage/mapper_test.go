package simplestorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapper(t *testing.T) {
	assert := assert.New(t)

	DataMap := NewMapper()
	Key := "IAmAKey"
	Data := "IAmData"

	//Put data in mapper
	DataMap.Put(Key, Data)

	// assert error Put invalid key
	err, _ := DataMap.Put("", "").(error)
	assert.Equal("invalid key", err.Error(), "must be return 'invalid key'")

	// assert equality
	assert.Equal(Data, DataMap.Get(Key), "Data should be equal")

	// assert error Get invalid key
	err, _ = DataMap.Get("").(error)
	assert.Equal("invalid key", err.Error(), "must be return 'invalid key'")

	// assert error Get empty data
	err, _ = DataMap.Get("fakekey :)").(error)
	assert.Equal("empty data", err.Error(), "suppose to be 'empty data'")
}
