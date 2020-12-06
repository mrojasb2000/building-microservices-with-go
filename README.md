# Building microservices with Go


## Preface

Microservice architecture is sweeping the world as the de facto pattern for building web-based application. Golang is a language particularly well suited to building them. Its strong community, encouragement of idiomatic style, and statically-linked binary artifacts make integrating it with other technologies and managing microservices at scale consistent and intuitive. this book will teach you the common patterns and practices, and show you how to apply these using the Go programming language.

It will teach you the fundamental concepts of 


#### Building docker image from microservices
- gcloud builds submit --tag gcr.io/PROJECT ID/rpc

#### Deploy image on Cloud Run (unauthenticated)
- gcloud run deploy --image gcr.io/PROJECT ID/rpc --platform managed --allow-unauthenticated

## Chapter 1

### Building a simple web server with net/http

`go run basic_http_example.go`

### Reading and writing JSON

To encode JSON data, the encoding/json package provides the Marshal function, which has the following signature:

func Marshal(v interface{})([]byte, err)

This function takes one parameter, which is of type interface, so pretty much any object you can think of since interface represents any type in Go. It returns a tuple of ([]byte, error), you will see this return style quite frequently in Go. some languages implement a try catch approach that encorages an error to be thrown when an operation cannot be performed, Go suggests the pattern (return type, error), where the error is nil when an operation succeeds.

In Go, unhandled errors are a bad thing, and whilst the language does implement Panic and Recover, which resemble exception handling in other languages, the situations where you should use these are quite different.
In Go la panic function causes normal execution to stop and all deferred function calls in the Go routine are executed, the program will then crash with a log message. It is generally used for unexpected errors that indicate a bug in the code and good robust Go code will attempt to handle these runtime exceptions and return a detailed error object back to the calling function.

This pattern is exactly what is implemented with the Marshal function. In the instance that Marshal cannot create a JSON encoded byte array from the given object, which could be due to a runtime panic, then this is captured and an error object detailing the problem is returned to the caller.


`go run reading_writing_json_1.go`