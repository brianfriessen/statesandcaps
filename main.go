package main

import (
	"bufio"
	"fmt"
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


	// Initial setup of the program which consists of initializing a file object and console object along with some variable initiization

	//Open the master file which has the state and capitals  (make this an array later stored with the program rather than a file)
	// os.open returns two values, a file handle and an error
	file, err := os.Open("states.txt")
	defer file.Close()   // defer means wait until the end of time and then close the file
	check(err)  // make sure the file exists and we were able to open it

	//Handle to the terminal console so we can get input from the user
	reader := bufio.NewReader(os.Stdin)

	//Handle to read a text file
	scanner := bufio.NewScanner(file)


	var lineNumber = 1  //  Counter to keep track of how many states we have tested
	var numberCorrect  =0  // Counter to keep track of number of correct answers
	var numberWrong = 0  // Counter to keep track of the number of wrong answers

	// Print a banner one time at the start of the program
	fmt.Println("Welcome to states and capitals.  Type EXIT to quit")

	// now lets read one line at a time from the file
	//scanner.Scan will read one line from the file in each iteration until it reaches the end of the file.
	for scanner.Scan() {
		fmt.Printf("Q %d/50: ", lineNumber)
		lineNumber += 1

		//each line should be of the format  state<tab>capital.  Store the line into variables
		// This code works. Use the strings package Split function to split on the tab <\t>
		thisLine := strings.Split(scanner.Text(), "\t")
		//for debug:  fmt.Printf("The capital of %s is %s.\n", thisLine[0], thisLine[1])

		//thisLine is a slice with first element being the state and second the capital..count from 0
		state := thisLine[0]
		capital := thisLine[1]

		//Prompt the user for his answer
		fmt.Printf("Correct: %2d, Incorrect %2d\n", numberCorrect, numberWrong)
		fmt.Printf("What is the Capital of %s? ", state)

		//Read user input up until enter (Carriage Return/Line Feed  aka CR/LF ) is entered
		text, _ := reader.ReadString('\n')

		//Strip out the carriage return and line feed from the user response so we just have the state name
		text = strings.Replace(text, "\n", "", -1)

		//Check to see if we got EXIT or exit to quit
		if strings.Compare(text,"EXIT") == 1 || strings.Compare(text, "exit") == 1 {
			os.Exit(1)
		}


		fmt.Printf("----> You said: %s, the answer is %s\n", text, capital)
		if strings.Compare(text, capital) == 0 {
			numberCorrect++
			fmt.Println("Correct!")
		} else { numberWrong++
		fmt.Println("Sorry, that is not correct")}

		/*
			This should work but doesn't..need to study up on how regular expressions work in go
			This should match on anything except TAB followed by TAB and then anything byt TAB
			splitStateandCap := regexp.MustCompile(`[^\t]*\t[^\t]*`)

			Above didn't work because Split works based on a single Character
			Verified the below will work OK
				re := regexp.MustCompile("\t")
				result := re.Split(scanner.Txt(), 2)
				State is in result[0] and capital in result[1]
				fmt.Println(result[0])
				fmt.Println(result[1])

			//print another line using our variable names
			fmt.Printf("The capital of %s is %s\n", splitStateandCap.Split(scanner.Text(), -1))

		*/

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
