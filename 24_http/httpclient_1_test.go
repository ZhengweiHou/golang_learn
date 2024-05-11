package http

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestHttpClient1(t *testing.T) {
	client := &http.Client{}
	client.Timeout = time.Second * 2

	httpreq, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		t.Fatal(err)
	}

	rep, err := client.Do(httpreq)
	if err != nil {
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			// t.Logf("Request timed out: %v", err)
			fmt.Printf("Request timed out: %v", err)
		} else {
			t.Fatalf("HTTP request failed: %v", err)
		}
		return
	}
	data, err := io.ReadAll(rep.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("rep:%s", data)
}
