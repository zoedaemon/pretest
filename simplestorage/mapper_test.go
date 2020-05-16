package simplestorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Put data to mapper
func TestMapper(t *testing.T) {
	assert := assert.New(t)

	DataMap := NewMapper()
	Key := "IAmAKey"
	Data := "IAmData"

	DataMap.Put(Key, Data)
	// assert equality
	assert.Equal(DataMap.Get(Key), Data, "Data Not Equal")

}
