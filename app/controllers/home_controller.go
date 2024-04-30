package controllers

import (
	"fmt"
	"net/http"
)



 
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to gotoko home page")
}