package controller

import "net/http"


func RegisterController(){
    userControllerInstance := newUserController();

    http.Handle("/users", *userControllerInstance);
    http.Handle("/users/", *userControllerInstance);

}