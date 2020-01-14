package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestFechGoogle(t *testing.T) {
	r, _ := http.NewRequest("GET", "https://google.com", nil)
	timeoutRequest, cancelFunc := context.WithTimeout(r.Context(), 1*time.Millisecond)
	defer cancelFunc()

	r = r.WithContext(timeoutRequest)

	_, err := http.DefaultClient.Do(r)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
