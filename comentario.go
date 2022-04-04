package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//going to call the csv file
	//flag will be able to use int, string or boolean options
	//use string to define the format
	//flag. string will define the name, default value and usage string
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question, answer'")
	flag.Parse()
   
	//os is a package that can work with every systems (win, linux, macos)
	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Failed to open csv file: %s\n, *csvFilename")
		os.Exit(1)
		//if the file doesn't open, it will show the message failed to open
		//the pointer is also use to csv file
	}
	r := csv.Newreader(file)
	//reader
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file")
		//ALL this part will be working as a reader to the csv file
	}
	problems := parseLines(lines)
	correct := 0
	fmt.Println(problems)
//parse will separate and read lines
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Printf("You score %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
     ret := make([]problem, len(lines))
     for i, line := range lines{
		 ret[i] = problem{
			 q:line[0]
			 a: string.TrimSpace(line[1])
		 }
	 }
	 return ret 
}
type problem struct {
	q string
	a string
}
func exit (msg string){
	fmt.Println((msg))
	os.Exit(1)
}	