package controllers

import "net/http"


func RegisterControllers(){
    userControllerInstance := newUserController();

    http.Handle("/users", *userControllerInstance);
    http.Handle("/users/", *userControllerInstance);

}