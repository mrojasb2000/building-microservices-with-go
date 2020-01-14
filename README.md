# building-microservices-with-go
Building Microservices with Go

# Building docker image from microservices
- gcloud builds submit --tag gcr.io/PROJECT ID/rpc

# Deploy image on Cloud Run (unauthenticated)
- gcloud run deploy --image gcr.io/PROJECT ID/rpc --platform managed --allow-unauthenticated