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
	// read the entire contents of the
	//dat, err := ioutil.ReadFile("states.txt")
	//check(err)
	//fmt.Print(string(dat))
	//fmt.Println("finished printing entire contents of file")

	// now lets read one line at a time from the file

	//fmt.Println("Now printing each line, one at a time.")

	// os.open returns two values, a file handle and an error
	file, err := os.Open("states.txt")
	check(err)

	//Handle to read a text file
	scanner := bufio.NewScanner(file)

	//Handle so we can get input from the user
	reader := bufio.NewReader(os.Stdin)

	//scanner.Scan will read one line from the file in each iteration until it reaches the end of the file.
	var lineNumber = 1
	for scanner.Scan() {
		fmt.Printf("Q %d/50: ", lineNumber)
		//fmt.Println(scanner.Text())
		lineNumber += 1
		//each line should be of the format  state<tab>capital.  Store the line into variables

		// This code works. Use the strings package Split function to split on the tab <\t>
		thisLine := strings.Split(scanner.Text(), "\t")
		//fmt.Printf("The capital of %s is %s.\n", thisLine[0], thisLine[1])

		state := thisLine[0]
		capital := thisLine[1]

		//Prompt the user for his answer
		fmt.Printf("What is the Capital of %s? ", state)
		text, _ := reader.ReadString('\n')

		//Strip out the carriage return and line feed from the user response so we just have the state name
		text = strings.Replace(text, "\n", "", -1)

		fmt.Printf("You said: %s, the answer is %s\n", text, capital)
		if strings.Compare(text, capital) == 0 {
			fmt.Println("Correct!")
		}
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

	defer file.Close()

}
