package routers

import (
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
	"github.com/zoedaemon/pretest/simplestorage"
)

// GetAllMessage get all message that hold up in simplestorage.Mapper
func GetAllMessage(scope simpleapi.Scope) *simpleapi.Response {

	//new data instance
	data := []Data{}

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
		singleData.Message = val.(string)
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
