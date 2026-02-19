/*
Copyright © 2026 Scryv

*/
package cmd

import (
	"encoding/csv"
    "fmt"
    "os"
	"log"
	"github.com/spf13/cobra"
)


var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists Data",
	Long: `This command lists the data that is stored within data.csv`,
	Run: func(cmd *cobra.Command, args []string) {
		listData()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
func listData(){
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
