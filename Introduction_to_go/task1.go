// This program is a student grade calculator.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var n int

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Hello Student!\nWhat is your name?\nName: ")
	name, err := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if err != nil {
		fmt.Println("Error while Reading name!")
		return
	}

	fmt.Printf("How many subjects Have you taken?\nSubject Count: ")
	fmt.Scanln(&n)

	fmt.Printf("N: %v, Name: %v \n", n, name)

	// Create a map for holding subjects and their corresponding grades
	subjects := make(map[string]float64, n)

	for i := 0; i < n; i++ {
		// Accept subject details

		fmt.Printf("Enter name of subject %v: ", i+1)
		subj_name, err := reader.ReadString('\n')
		subj_name = strings.TrimSpace(subj_name)

		var grade float64
		
		if err != nil {
			fmt.Println("Error Reading subject name!")
			return
		}

		fmt.Printf("Enter your grade in %v : ", subj_name)
		fmt.Scanln(&grade)
		
		if grade > 4.0 || grade < 0.0 {
			fmt.Println("Invalid grade entered!")
			return
		}

		subjects[subj_name] = grade
	}
	var avg float64
	fmt.Printf("Here is your grade report %s!\n\nSubject\t\tGrade\n", name)
	for subj, grd := range subjects {
		fmt.Printf("%s\t\t%v\n", subj, grd)
		avg += grd
	}

	avg /= float64(n)

	fmt.Printf("\nCalculated GPA: %v\n", avg)
}