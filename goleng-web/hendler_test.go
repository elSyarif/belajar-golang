package goleng_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Images ")
	})

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestReques(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
	}

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
