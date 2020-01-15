package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var name string = "Mavro!"

type helloWorldRequest struct {
	Name string `json:"name"`
}

// HelloWorldResponse http response type
type HelloWorldResponse struct {
	Message string `json:"message"`
}

func main() {
	port := 8080

	http.Handle("/helloworld",
		NewGzipHandler(http.HandlerFunc(helloWorldHandler)),
	)

	log.Printf("Server staring on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(rw http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "Hello " + name}
	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}

// NewGzipHandler new handler
func NewGzipHandler(next http.Handler) http.Handler {
	return &GZipHandler{next}
}

// GZipHandler handler type
type GZipHandler struct {
	next http.Handler
}

func (h *GZipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encodings := r.Header.Get("Accept-Encoding")

	if strings.Contains(encodings, "gzip") {
		h.serveGzipped(w, r)
	} else if strings.Contains(encodings, "deflate") {
		panic("Deflate not implemented")
	} else {
		h.servePlain(w, r)
	}
}

func (h *GZipHandler) serveGzipped(w http.ResponseWriter, r *http.Request) {
	gzw := gzip.NewWriter(w)
	defer gzw.Close()

	w.Header().Set("Content-Encoding", "gzip")
	h.next.ServeHTTP(GzipResponseWriter{gzw, w}, r)
}

func (h *GZipHandler) servePlain(w http.ResponseWriter, r *http.Request) {
	h.next.ServeHTTP(w, r)
}

// GzipResponseWriter my own ResponseWriter that embeds http.ResponseWriter
type GzipResponseWriter struct {
	gw *gzip.Writer
	http.ResponseWriter
}

// The core method for this is the implementation of ther Write method
func (w GzipResponseWriter) Write(b []byte) (int, error) {
	if _, ok := w.Header()["Content-Type"]; !ok {
		// If content type is not set, infer it from the uncompressed body.
		w.Header().Set("Content-Type", http.DetectContentType(b))
	}
	return w.gw.Write(b)
}

// Flush method for response writer
func (w GzipResponseWriter) Flush() {
	w.gw.Flush()
	if fw, ok := w.ResponseWriter.(http.Flusher); ok {
		fw.Flush()
	}
}
