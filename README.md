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
## Endpoints
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
***author: Roy Adventus***