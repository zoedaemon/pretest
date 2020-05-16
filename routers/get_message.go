package routers

import (
	"net/http"

	"github.com/zoedaemon/pretest/simpleapi"
	"github.com/zoedaemon/pretest/simplestorage"
)

//GetMessage handle get message by key in query parameters
func GetMessage(scope simpleapi.Scope) *simpleapi.Response {

	//new data instance
	data := &Data{}

	//get request
	req := scope.GetRequest()

	//get key
	//TODO: implement in simpleapi.QueryParm
	Key := req.URL.Query()["key"][0]

	//get custom data that hold Mapper
	CData := scope.GetCustomData().(*simplestorage.Mapper)
	CData.PrintData()
	CData.IsPrintLog = true

	//cek if data exist or not
	Error, isError := CData.Get(Key).(error)
	if isError {
		return &simpleapi.Response{
			Error:        Error,
			ResponseCode: http.StatusNotFound,
		}
	}

	//get data by key and assign for response data
	dataMapped := CData.Get(Key).(string)
	data.ID = Key
	data.Message = dataMapped

	//return message passed and created
	return &simpleapi.Response{
		ResponseCode: http.StatusOK,
		Data:         data,
	}
}
