package main

import (
	"net/http"
	"testing"
)

func BenchmarkLlamadasApi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := http.Get("http://localhost:8080/myml/2")
		if err != nil {
			panic(err)
		}
	}
}
