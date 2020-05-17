# Simple API
## Build / Run
Make sure you have installed go correctly
```
$ go build
$ ./pretest
``` 
Or
```
$ go run main.go
```
## Test & Coverage
```
$ go test ./... --coverprofile cover.out
```
then 
```
$ go tool cover -html cover.out 
```
go tool automatically open up browser for you to show coverage of test

## Endpoints
### Send Message
POST ```http://localhost:8080/messages/send```
with payload ```{"message": "zoed"}```
and get response like this
```
{
    "response-code": 201,
    "error": null,
    "data": {
        "detail": {
            "id": "whTPIoVfymn8QGa969PaEYvXAwc=",
            "message": "zoed"
        },
        "info": "success send/store message"
    }
}
```
### Get Message By Key
GET ```http://localhost:8080/messages/send?key=whTPIoVfymn8QGa969PaEYvXAwc=```
and the response suppose to be
```
{
    "response-code": 200,
    "error": null,
    "data": {
        "id": "whTPIoVfymn8QGa969PaEYvXAwc=",
        "message": "zoed"
    }
}
```
### Get All Messages
GET ```http://localhost:8080/messages```
for getting all messages that have been post with ```/messages/send``` 
```
{
    "response-code": 200,
    "error": null,
    "data": {
        "id": "whTPIoVfymn8QGa969PaEYvXAwc=",
        "message": "zoed"
    }
}
```
### Websocket Message
access ```ws://127.0.0.1:8080/ws/messages/echo``` Postwoman or similliar tools, and try to connect and send message.
If you want see the test result for websocket try this command
```
$ go test github.com/zoedaemon/pretest/routers -run TestWebsocketMessage -v 
```
Output suppose to be
```
=== RUN   TestWebsocketMessage
2020/05/17 08:21:56 Websocket connection started...
2020/05/17 08:21:56 some message has been receive : cDiXJwJ
2020/05/17 08:21:56 some message has been receive : HXciBrG
2020/05/17 08:21:56 some message has been receive : PcoUsGw
2020/05/17 08:21:56 some message has been receive : IYwrUIP
2020/05/17 08:21:56 some message has been receive : lfEDEbP
2020/05/17 08:21:56 some message has been receive : NyUxGdl
2020/05/17 08:21:56 some message has been receive : hiweNTs
2020/05/17 08:21:56 some message has been receive : TlFMnRx
2020/05/17 08:21:56 some message has been receive : iLyxsDm
2020/05/17 08:21:56 some message has been receive : lmSlfBc
--- PASS: TestWebsocketMessage (0.01s)
PASS
2020/05/17 08:21:56 Websocket disconnected...bye...
ok      github.com/zoedaemon/pretest/routers
```

***author: Roy Adventus***