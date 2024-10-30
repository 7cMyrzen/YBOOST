package Handlers

import (
	"fmt"
	"net/http"
)

func Cocktails(w http.ResponseWriter, r *http.Request) {
	fmt.Println("cocktails")
}
