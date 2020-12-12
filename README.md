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

`go run reading_writing_json_3.go`
`go run reading_writing_json_4.go`
`go run reading_writing_json_5.go`


### Routing in net/http

Even a simple microservice will need the capabilility to route requests to diferent handlers dependent on the requested path or method. In Go this is handled by the DefaultServerMux method which is an instance of ServerMux. Earlier in this chapter, we briefly covered that when nil is passed to the handler parameter for the ListenAndServe function then the DefaultServeMux method is used. When we call the http.HandleFunc("/helloWorld", helloWorldHandler) package function we are actually just indirectly calling http.DefaultServerMux.HandleFunc(...).

The Go HTTP server does not have a specific router instead any object which implements the http.Handler interface is passed as a top level function to the Listen() function, when a request comes into the server the ServeHTTP method of this handler is called and it is responsible for performing or delegating any work. To facilitate the handling of multiple routes the HTTP package has a special object called ServerMux, which implements the http.Handler interface.

There are two functions to adding handlers to a ServerMux handler:

    func HandlerFunc(pattern string, handler func(ResponseWriter, *Request))
    func Handle(pattern string, handler Handler)

The HandlerFunc function is a convenience function that creates a handler who's ServeHTTP method calls an ordinary function with the func(ResponseWriter, *Request) signature that you pass as a parameter.

The Handle function requires that you pass two parameters, the pattern that you would like to register the handler and an object that implements the Handler interfaces:

    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }


### Paths

We already explained how ServeMux is responsable for routing inbound requests to the regustered handlers, however the way that the router has a very simple routing model it does not support wildcards or regular expressions, with ServeMux you must be expicit about the registered paths.

You can register both fixed rooted paths, such as /images/file.png, or rooted subtrees such as /images/. The trailing slash in the rooted subtree is important as any request that starts with /images/, for example /images/other_file.png, would be routed to the handler associated with /images/.

if we register a path /images/ to the handler foo, and the user makes a request to our service as /images (note no trailing slash), then ServeMux will forward the request to the /images/ handler, appending a trailing slash.

if we also regsiter the path /images (note no trailing slash) to the handler bar the user requests /images then this request will be directed to bar; however, /images/ or /images/file.png will be directed to foo:

    http.Handle("/images/", newFooHandler())
    http.Handle("/images/persiab/", newBarHandler())
    http.Handle("/images", newBuzzHandler())

    /images                 => Buzz
    /images/                => Foo
    /images/cat             => Foo
    /images/cat.png         => Foo
    /images/oersian/cat.png => Bar

Longer paths will always take precedence over shorte ones so it is posible to have an explicit toute that points to a different handler to a catch all route.

We can also specify the hostname, we could register a path souch as search.google.com/ and /ServeMux would forward any requests to http://search.google.com and http://www.google.com to their respective handlers.

If you are used to a framework based application development approach such as using Ruby on Rails or ExpressJS you may find this router incredibly simple and it is, remember that we are not using a framework but the standard packages of Go, the intention is always to provide a basis that can be built upon. In very simple cases the ServeMux approach more than good enough and in fact I personally don't use anything else. Everyone's needs are different however and the beauty and simple to build your own route as all is needed is an object which implements the Handler interface. A quick trawl through google will surface some very good third party routes but my recommendation for you is to learn the limitations of ServeMux first before deciding to choose a third-party package it will greatly help with your decision process as you will know the problem you are trying to solve.

### FileServer

A FileServer function returns a handler that server HTTP requests with the contents of the filesystem. This can be used to serve static files such as images or other content that is stored on the file system:

    func FileServer(root FileSystem) Handler

Take a look at the following code:

    http.Handle("/images", http.FileServer(http.Dir("./images")))

this allows us to map the contents of the file system path ./images to the server route /images, Dir implements a file system which is restricted to a specific directory tree, the FileServer method uses this to be able to serve the assests.

### NotFoundHandler

The NotFoundHandler function returns a simple request handler that replices to each request qith a 404 page not found reply:

    func NotFoundHandler() Handler

    