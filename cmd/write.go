/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
    "fmt"
    "os"
	"log"
	"bufio"
	"strings"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Lets you write data",
	Long: `This command lets u parse data in the data.csv file`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
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
			run()
		default:
			fmt.Println("Not a valid option ( Y - Yes or N - No)")
			run()
	}
}
func run() {
	reader := bufio.NewReader(os.Stdin)
	record := getUserInfo(reader)
	confirm(record, reader)
}
