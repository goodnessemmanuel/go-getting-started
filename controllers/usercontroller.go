package controllers

import (
	"net/http"
	"regexp"
)

type UserController  struct {
    userIDPattern *regexp.Regexp
}

//creating a method
func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Hello from the user controller!"))
}

//constructor
func newUserController() *UserController{
    userControllerInstance := UserController{
        userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
    }
    return &userControllerInstance
}