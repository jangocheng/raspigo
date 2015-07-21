package main

import (
	"github.com/xam090/raspigo"
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Starting server!")
	http.Handle("/", raspigo.GetRaspiStatHandler())
	fmt.Println("Ready!")
	http.ListenAndServe(":10101", nil);
}
