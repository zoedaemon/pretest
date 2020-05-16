package simplestorage

import (
	"strconv"
	"sync"
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

func TestMapperConcurrency(t *testing.T) {
	assert := assert.New(t)

	DataMap := NewMapper()
	Keys := []string{"key1", "key2"} //, "key3", "key4", "key5"}
	Data := "sequenced:"

	//set print log for delayed execution so race condition could be occurs
	DataMap.IsPrintLog = true

	//define waitgroup
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()

			//limiting key by modulus
			index := i % len(Keys)
			Key := Keys[index]
			Data += strconv.Itoa(i)

			//Put data in mapper(concurrent put goroutines can be write at same key)
			DataMap.Put(Key, Data)

			//check data
			assert.Equal(Data, DataMap.Get(Key), "Data should be equal")
		}()

		//wait for all goroutines to finish
		wg.Wait()
	}
}
