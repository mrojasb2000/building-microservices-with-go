# Building microservices with Go


## Preface

Microservice architecture is sweeping the world as the de facto pattern for building web-based application. Golang is a language particularly well suited to building them. Its strong community, encouragement of idiomatic style, and statically-linked binary artifacts make integrating it with other technologies and managing microservices at scale consistent and intuitive. this book will teach you the common patterns and practices, and show you how to apply these using the Go programming language.

It will teach you the fundamental concepts of 


#### Building docker image from microservices
- gcloud builds submit --tag gcr.io/PROJECT ID/rpc

#### Deploy image on Cloud Run (unauthenticated)
- gcloud run deploy --image gcr.io/PROJECT ID/rpc --platform managed --allow-unauthenticated