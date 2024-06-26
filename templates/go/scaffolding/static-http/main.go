package main

import (
	"fmt"
	"os"

	"knative.dev/func-go/http"

	f "f"
)

func main() {
	if err := http.Start(http.DefaultHandler{Handler: f.Handle}); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
