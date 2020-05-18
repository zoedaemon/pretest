package routers

import (
	"fmt"
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
	"github.com/zoedaemon/pretest/simplestorage"
)

// GetAllMessage get all message that hold up in simplestorage.Mapper
func GetAllMessage(scope simpleapi.Scope) *simpleapi.Response {

	//new data instance
	data := []Data{}

	fmt.Printf("scope.GetCustomData() = %+v ", scope.GetCustomData())
	//get custom data that hold Mapper
	CData := scope.GetCustomData().(*simplestorage.Mapper)
	DataMaps := CData.GetRefData()

	if len(*DataMaps) == 0 {
		//do not necessarilly send data or error detail...
		//...coz 204 not showing any data response in browser or Postman
		return &simpleapi.Response{
			ResponseCode: http.StatusNoContent,
		}
	}

	for key, val := range *DataMaps {
		singleData := Data{}
		singleData.ID = key
		//check data is must string or []byte
		switch val.(type) {
		case string:
			singleData.Message = val.(string)
			break

		case []byte:
			bytes := val.([]byte)
			singleData.Message = string(bytes)
			break
		default:
			continue
		}
		data = append(data, singleData)
	}

	dataResponse := map[string]interface{}{
		"size": len(*DataMaps),
		"data": data,
	}
	//return message passed and created
	return &simpleapi.Response{
		ResponseCode: http.StatusOK,
		Data:         dataResponse,
	}
}
