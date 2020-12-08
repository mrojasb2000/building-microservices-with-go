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

Using the strut field's tags, we can have greater control of how the output will look. In the preceding example. when we marshal this struct the output from our server would be:

{"message": "Hello World!"}

`go run reading_writing_json_2.go`


## Unmarshalling JSON to Go structs

Now we have learnd how we can send JSON back to the client, what if we need to read input before returning the output? We could use URL parameters and we will see what that is all about in the next chapter, but commonly you will need more complex data structures that involve the service to accept JSON as part of an HTTP POST request.

Appliying similar techniques that we learned in the previous section to write JSON, reading JSON is just as easy. To decode JSON into a struct field the encoding/json package provides us with the Unmarshal function:

func Unmarshal(data []byte, v interface{}) error

The Unmarshal function works in the opposite way to Marshal; it allocates maps, slices, and pointers as required. Incoming object keys are matched using either the struct field name or its tag and will work a case-insensitive match; however, an exact match is preferred. Like Marshal, Unmarshal will only set exported struct fields, those that start with an upper-case letter.

We start by adding a new struct field to represent the request, whilst Unmarshal can decode the JSON into an interface{}, which would be of map[string]interface{} // for JSON objects type or: []interface{} // for JSON arrays, depending if out JSON is an object or an array.

In my opinion it is much clearer to the readers of our code if we explicity state what we are expecting as a request. We can also save ourselves work by not having to manually cast the data when we come to use it.

Remember two things:

* You do not write code for the compiler, you write code for humans to understand
* You will spend more time reading code than you do writing it

