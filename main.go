package main

import (
	"fmt"
	"net/http"

	"github.com/goodnessemmanuel/getting-started/controllers"
)

func main() {
   //invoking the created webserver 
   controllers.RegisterControllers();
   http.ListenAndServe(":3030", nil)  //listen to this port on my local machine and route any request coming to it
    
   fmt.Println("Hello world")

}