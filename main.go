package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// this function will help with checking for errors.  If e is anything other than empty print the error code and exit
// func format is <function name>(<var name> <var type>) { }
func check(e error) {
	if e != nil {
		fmt.Println("Huston we have a problem..")
		log.Fatal(e)
		panic(e)
	}
}

func main() {
	// read the entire contents of the
	dat, err := ioutil.ReadFile("states.txt")
	check(err)
	fmt.Print(string(dat))
	fmt.Println("finished printing entire contents of file")

	// now lets read one line at a time from the file

	fmt.Println("Now printing each line, one at a time.")

	// os.open returns two values, a file handle and an error
	file, err := os.Open("states.txt")
	check(err)

	scanner := bufio.NewScanner(file)

	//scanner.Scan will read one line from the file in each iteration until it reaches the end of the file.
	var lineNumber = 0
	for scanner.Scan() {
		fmt.Printf("line number %d: ", lineNumber)
		fmt.Println(scanner.Text())
		lineNumber += 1
		//each line should be of the format  state<tab>capital.  Store the line into variables
		thisLine := strings.Fields(scanner.Text())

		// have some bug here need to find a way to split based on <TAB> not just whitespace characters, else messed up
		state := thisLine[0]
		//Capitals with more than one world (like Salt Lake City) require more than one slice
		capital := thisLine[1:]

		//print another line using our variable names
		fmt.Printf("\nThe capital of %s is %s\n", state, capital)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	defer file.Close()

}
