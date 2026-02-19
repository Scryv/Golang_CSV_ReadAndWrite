package main

import (
	"encoding/csv"
    "fmt"
    "os"
	"log"
	"bufio"
	"strings"
)
//read
func getUserInfo(reader *bufio.Reader) []string{  //so i can rerun it (ik there are better ways leave me alone)
	firstName, _ := getInput("First name: ", reader)
	lastName, _ := getInput("Last name: ", reader)
	email, _ := getInput("Email: ", reader)
	country, _ := getInput("Country: ", reader)
    record := []string{firstName, lastName, email, country}
	return record
}
func getInput(prompt string, r *bufio.Reader) (string, error){ //input func
	fmt.Print(prompt)
	input, error := r.ReadString('\n') //after enter read
	return strings.TrimSpace(input), error 
}

func readCSV(){
	file, err := os.Open("data.csv") //opens data.csv
	if err != nil { //checks for erroorrs
		log.Fatal(err) //if error it shall log
	}
	defer file.Close() //close DONT FORGET DEFER MEANS END
	reader := csv.NewReader(file) // creates reader
	records, err := reader.ReadAll() // reads all and stores in records var
	if err != nil { //SHALL LOGG IF NOT NIL
		log.Fatal(err)
	}
	for _, bomba := range records { //_ index no value yes
		fmt.Println(bomba)
	}
}
//write
func writeCSV(record []string,){
	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644) //Opens file in append mode u can also just delete everything after , and just Create
	if err != nil { 
		log.Fatal(err) 
	}
	defer file.Close()
	writer := csv.NewWriter(file) //creates writer
	defer writer.Flush()  //flushes that shit
	if err := writer.Write(record); 
	err != nil {
		log.Fatal(err)
	}
}
func confirm(record []string, reader *bufio.Reader) { //imports record from main and reader
	fmt.Println(record) //shows ur slice
	Confirm, _ := getInput("your sure you want to append: ", reader)
	switch Confirm { //da switch if yess shall append if no shall abort Default just if sm wrong was typed
		case "y", "Y", "yes", "Yes":
			writeCSV(record)
		case "n", "N", "no", "No":
			fmt.Println("Aborted")
			main()
		default:
			fmt.Println("Not a valid option ( Y - Yes or N - No)")
			main()
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	record := getUserInfo(reader)
	confirm(record, reader)
}
