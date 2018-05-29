package main

import (
	"bufio"
	"fmt"
	"os"
)

type user struct {
	name     string
	password string
	message  string
}

func main() {
	fmt.Println("Welcome to Locker")
	fmt.Println("*****************")
	users := make([]user, 10)

	for true {
		fmt.Println("1. New User, 2. Existing User, 3. Exit")
		var input string
		fmt.Scan(&input)
		if input == "1" {
			newUser1 := user{}
			getInput(&newUser1)
			users = append(users, newUser1)
		} else if input == "2" {
			findUser(users)
		} else if input == "3" {
			fmt.Println("Thanks for using Locker Application.")
			os.Exit(1)
		}
	}

	fmt.Println(users)
}

func getInput(u *user) {
	var input string

	fmt.Print("Enter name     : ")
	fmt.Scan(&input)
	u.name = input

	fmt.Print("Enter password : ")
	fmt.Scan(&input)
	u.password = input

	fmt.Println("Enter secret Message : ")
	u.message, _ = bufio.NewReader(os.Stdin).ReadString('\n')
}

func findUser(users []user) {
	var name string
	var password string
	found := false
	fmt.Print("Enter name     : ")
	fmt.Scan(&name)

	fmt.Print("Enter password : ")
	fmt.Scan(&password)

	for i := 0; i < len(users); i++ {
		if users[i].name == name && users[i].password == password {
			fmt.Println("Your secret Message : ", users[i].message)
			found = true
			break
		} else if users[i].name == name {
			fmt.Println("Wrong password, Please try again.")
			found = true
			break
		}
	}

	if found == false {
		fmt.Println("User '", name, "' Not found, Please register.")
	}
}

// ➜  map go run main.go
// Welcome to Locker
// *****************
// 1. New User, 2. Existing User, 3. Exit
// 1
// Enter name     : p
// Enter password : k
// Enter secret Message :
// hola
// 1. New User, 2. Existing User, 3. Exit
// 2
// Enter name     : p
// Enter password : k
// Your secret Message :  hola

// 1. New User, 2. Existing User, 3. Exit
// 1
// Enter name     : a
// Enter password : b
// Enter secret Message :
// manzana
// 1. New User, 2. Existing User, 3. Exit
// 2
// Enter name     : a
// Enter password : b
// Your secret Message :  manzana

// 1. New User, 2. Existing User, 3. Exit

// 2
// Enter name     : p
// Enter password : k
// Your secret Message :  hola

// 1. New User, 2. Existing User, 3. Exit
// 3
// Thanks for using Locker Application.
// exit status 1
