package main

import (
	"fmt"
	//"net/http"
	//"github.com/goodnessemmanuel/getting-started/controllers"
)

func main() {
	//invoking the created webserver
	//controllers.RegisterControllers();
	// http.ListenAndServe(":3030", nil)  //listen to this port on my local machine and route any request coming to it

	demonstrateLoop1()
	demonstrateLoop2()
    demonstratePanic()

}

// Go only uses a for loop
func demonstrateLoop1() {
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 { //stop looping if this condition  is met
			break
		}
	}
}

// Go only uses a for loop
func demonstrateLoop2() {
	fmt.Println("**** Inside demonstrateLoop2 ******")
	var i int
	for i < 5 {
		fmt.Println(i)
		i++
		if i == 3 { //
			continue
		}
		fmt.Println("Continuing.... \n")
	}

	fmt.Println("Global i after increment: ", i) //casting i to it's string format

	for i := 0; i < 5; i++ { //notice this 'i' is only recognize within this forloop block and doesn't interfer with the global 'i'
		fmt.Println(i)
	}

	//infinite loops
	var j int
	for { //without the inner if, this loop will run forever!
		fmt.Println(j)
		if j == 5 {
			break
		}

		j++
	}

}

// A panic is similar to exception (more like throwing an exception) in other language where an application can no longer proceed
func demonstratePanic(){
        fmt.Println("Starting web  server")
        // do something here

        panic("Something bad just happened! :)")

        fmt.Println("This line of the code will not be reached due to the panic on line 68")

}
