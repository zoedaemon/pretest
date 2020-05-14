package simpleapi

import "net/http"

type (

	//HandlerFunc SimpleAPI implementation
	HandlerFunc func(*Scope) error

	//Scope SimpleAPI all attributes / field
	Scope struct {
		//TODO: should be private attributes, but need implement setter and getter,
		//		for moment keep it simple
		Request    *http.Request
		Response   http.ResponseWriter
		CustomData interface{} //e.g. map objects of files config, database connection, redis, etc
	}
)

//New SimpleAPI
func New() *Scope {
	return &Scope{}
}

//GetMethod register handler for Get method
func (s *Scope) GetMethod(path string, h HandlerFunc) {

}

//PostMethod register handler for Get method
func (s *Scope) PostMethod(path string, h HandlerFunc) {

}
