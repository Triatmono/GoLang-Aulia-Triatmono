# Golang-Aulia-Triatmono
 
## Progress up to Switch Stament (1.54.02)
### 0:00 - Introduction to the Course
Learn concepts isolated without context
core concepts such as
- Data Types
- Variables & Constants
- Formatted Output
- User Input
- Pointers
- Scope Rules
- Loops
- If-else & Switch
- Functions
- Packages
- Goroutines

### 02.47 - What is Go ?
Programming language developed by Google in 2007
Go is a statically typed, compiled language known for its simplicity, efficiency, and concurrency capabilities. Developed by Google, it’s designed for high-performance server-side applications.
It focuses on simplicity with minimal syntax and is particularly known for its robust concurrency support, making it ideal for network servers and distributed systems.

### 6.50 - Characteristics of Go and Go Use Cases GO SYNTAX & CONCEPTS
Characteristics of Go such as fast compilation, efficient memory management, built-in garbage collection, and excellent support for concurrent programming. Use Cases: System programming, cloud computing, and large-scale server applications are common areas where Go shines.

### 08:59 - Local Setup - Install Go & Editor
- Download Go from [here](https://go.dev/doc/install)
- Install Go on your local machine
- Install an editor Visual Studio Code
- Install Go extensions for your editor

### 12:54 - Write our First Program & Structure of a Go File
We can run our first program by opening our terminal on Visual Studio Code ```go mod init booking-app```. After this procedure is complete then we can start making our code for example :

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
After the code is finish, then we can compile it by using ```go run main.go```

### 22:02 - Variables & Constants in Go 
We can print variables in Go as followings
```go
package main

import "fmt"

func main(){
    fmt.Println("Welcome to our conference booking application")
    fmt.Println("Get your tickets here to attend")

    var conferenceName = "Go Conference"
    fmt.Println(conferenceName)
}
```
We can also do this instead :
```go
package main

import "fmt"

func main(){
    var conferenceName = "Go Conference"
    fmt.Println("Welcome to", conferenceName, "booking application")
    fmt.Println("Get your tickets here to attend")
}
```
Here are the code for the constants in Go
```go
package main

import "fmt"

func main(){
    var conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets = 50

    fmt.Println("Welcome to", conferenceName, "booking application")
    fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
    fmt.Println("Get your tickets here to attend")
}
```

### 30:43 - Formatted Output - printf
```go 
fmt.printf("some text with %s", myVariable)
```
Here are the codes for the formatted output
```go
package main

import "fmt"

func main(){
    var conferenceName = "Go Conference"
    const conferenceTickets = 50
    var remainingTickets = 50

    fmt.Printf("Welcome to %v booking application", conferenceName)
    fmt.Printf("We have total of %v tickets and %v are still available.", conferenceTickets, remainingTickets)
    fmt.Println("Get your tickets here to attend")
}
```
### 33:43 - Data Types in Go
Primary Data Types:
- String: var greeting string = "Hello"
- Integer: var age int = 30
- Float: var price float64 = 19.99
- Boolean: var isAvailable bool = true

Go is statically typed, meaning each variable type is declared and enforced at compile time.

### 45:18 - Getting User Input
In Go language, we can use ```fmt.Scanln ``` to get input. Here are the example code:
```go
package main

import "fmt"
func main(){
            var firstName string
            var lastName string
            var email string
            var userTickets uint

            
            // ask user for their name
            fmt.Println("Please enter your first name:")
            fmt.Scan(&firstName)

            fmt.Println("Please enter your last name:")
            fmt.Scan(&lastName)

            fmt.Println("Please enter your email:")
            fmt.Scan(&email)

            fmt.Println("Please enter number of tickets:")
            fmt.Scan(&userTickets)
}
```
### 47:19 - What is a Pointer ?
Pointer is a variable that points to the memory address of another variable. Here is the snippet example below:
```go
var firstName string
    fmt.Println("Please enter your first name:")
    fmt.Scan(&firstName)
```
### 53:55 - Book Ticket Logic
Suppose we want the number of tickets is reduced. We can use the following logic to our code.
```go
package main
import "fmt"
func main(){
    var remainingTickets uint = 50
    var userTickets uint

    fmt.Println("Please enter number of tickets:")
		fmt.Scan(&userTickets)

    remainingTickets = remainingTickets - userTickets
}
```
### 57:16 - Arrays & Slices
Array is a fixed size in Go. Here are the following code for implementing array.
```go
var bookings = [50]string{"Nana", "Nicole", "Peter"}
```
Slice is a dynamic size in Go. Here are the following code for implementing slice.
```go
var bookings []string
bookings = append(bookings, firstName + " " + lastName)
```
### 1:11:12 - Loops in Go
In General, a loop statement allows us to execute code multiple times, in a loop. Here are the snippet code to implemented for loop.
```go
for remainingTickets > 0 && len(bookings) < 50 {
		
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		
		// ask user for their name
		fmt.Println("Please enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Please enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Please enter your email:")
		fmt.Scan(&email)

		fmt.Println("Please enter number of tickets:")
		fmt.Scan(&userTickets)
		
		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber{
			
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings,firstName + " " + lastName )

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining{
				fmt.Println("Tickets are sold out. Come back another time.")
				break
			}
		}  else {
			if !isValidName{
				fmt.Println("First name and last name must be at least 2 characters.")
			} 
			if !isValidEmail{
				fmt.Println("Email must be valid.")
			}
			if  !isValidTicketNumber{
				fmt.Println("Number of tickets is invalid.")
			}			
		}
	}
```
### 1:24:24 - Conditionals (if / else) and Boolean Data Type & 1:39:33 - Validate User Input
Conditionals are used to execute different blocks of code based on a condition. The `if` statement is used to execute a block of code if a condition is true.

```go
        isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
        if isValidName && isValidEmail && isValidTicketNumber{
			
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings,firstName + " " + lastName )

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings{
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining{
				fmt.Println("Tickets are sold out. Come back another time.")
				break
			}
		}  else {
			if !isValidName{
				fmt.Println("First name and last name must be at least 2 characters.")
			} 
			if !isValidEmail{
				fmt.Println("Email must be valid.")
			}
			if  !isValidTicketNumber{
				fmt.Println("Number of tickets is invalid.")
			}			
		}
```
From the code above, we can determine that it will check the Name, Email, and Ticket number is valid or not. If the client input is wrong or invalid, then the program will execute the else statement and explain where the client input is wrong. This also can validate the user input so that the info will be correct.

### 1:54:02 - Switch Statement 
In Go, a switch statement is used to execute different blocks of code based on the value of a variable or expression. It provides a more concise way to handle multiple conditions compared to using multiple if-else statements. The switch statement evaluates an expression and matches its value against different cases. For example:
```go
city := "London"

	switch city {
		case "New York":
			//code
		case "Singapore":

		case "London":

		case "Berlin":

		case "Hong Kong":
		
		default:
			fmt.Print("No valid city selected\n")
	}
```
### 1:58:37 - Encapsulate Logic with Functions
```go
func main(){
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

		greetUsers(conferenceName, conferenceTickets, remainingTickets)
}


func greetUsers(confName string, confTickets int, remainingTickets uint){
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("Get your ticket now")
}
```

### 2:22:36 - Organize Code with Go Packages
We can make a new file such as ```helper.go``` to organize the code. Take a look at an example

```go
	helper.go

	package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
```

```go
	main.go


	isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)
```
To run we can use ```go run main.go helper.go``` it would run without any error.

### 2:37:16 - Maps

```go
var userData = make(map[string]string)
	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	strconv.FormatUint(uint64(userTickets), 10)
```

### 2:53:20 - Structs
```go
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
	}
```
### 3:02:15 - Goroutines - Concurrency in Go
```go
package main

import (
	"booking-app/helper"
	"fmt"
	"time"
	"sync"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

		greetUsers()

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email)
			

			firstNames := getFirstnames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			noTicketsRemaining := remainingTickets == 0
			if noTicketsRemaining {
				fmt.Println("Tickets are sold out. Come back another time.")
				//break
			}
		} else {
			if !isValidName {
				fmt.Println("First name and last name must be at least 2 characters.")
			}
			if !isValidEmail {
				fmt.Println("Email must be valid.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets is invalid.")
			}
		}
		wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket now")
}

func getFirstnames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("Please enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData {
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		numberOfTickets:   userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string){
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#########")
	wg.Done()
}

```