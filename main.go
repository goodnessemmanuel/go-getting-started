package main

import (
	"errors"
	"fmt"

	"github.com/goodnessemmanuel/getting-started/models"
)

func main() {
   //creating a struct instance 
    user1 := models.User{
        ID: 2,
        FirstName: "Oche",
        LastName: "King",
    }
    fmt.Println(user1)
    
    port := 3030

    //call function with the required arguments
   err := startWebServer(port, 2)
   fmt.Println(err)
   errorCode, err2 := startServerWithMultipleReturns(port, 4)

   fmt.Println(err2, errorCode)


}
//function that accepts two parameters
func startWebServer(port int, numberOftries int) error{
    fmt.Println("Starting server...")
    // do some important things here
    fmt.Println("Server started on port", port)
    fmt.Println("Number of tries ", numberOftries)

    //error string should not end with a punctuation mark or be capitalize
    return errors.New("something went wrong :)");
}

func startServerWithMultipleReturns(port int, numberOftries int) (int, error){
    fmt.Println("Starting server...")
    // do some important things here
    fmt.Println("Server started on port", port)
    fmt.Println("Number of tries ", numberOftries)

    //error string should not end with a punctuation mark or be capitalize
    return 404, errors.New("something went wrong :)");

}    