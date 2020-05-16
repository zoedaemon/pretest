package simpleapi

import "net/http"

//Context hold unique request and response for manipulating directly from user defined func
type Context struct {
	//DONE: should be private attributes, but need implement setter and getter,
	//		for moment keep it simple
	request  *http.Request
	response http.ResponseWriter
}

/*SetContext instance of Context
* writer : stream response that can be written on
* request: incoming http request
**/
func (c *Context) SetContext(writer http.ResponseWriter, request *http.Request) {
	c.request = request
	c.response = writer
}

//GetRequest current context
func (c *Context) GetRequest() *http.Request {
	return c.request
}

//GetResponse for direct writing to output response stream
func (c *Context) GetResponse() http.ResponseWriter {
	return c.response
}
