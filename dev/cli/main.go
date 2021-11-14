// https://www.allhandsontech.com/programming/golang/how-to-use-sqlite-with-go/
package main

import (

	_ "github.com/mattn/go-sqlite3"
	"log"
	"github.com/dixonwille/wmenu/v5"
	"database/sql"
	"os"
	"bufio"
	"strings"
	"fmt"

)

type theTriblet struct {

        id              int
        cmd             string
        aPipe           string
        desc            string
        msg             string
        restart         string
        root            string

}

func main() {

	// Connect to the database
	// Create a database object and open the sqlite db.	
	db, err := sql.Open("sqlite3", "./troublevent.db")
	checkErr(err)

	// Defer close
	defer db.Close()

	menu := wmenu.NewMenu("What would you like to do?")

	// Handles which option is selected.
	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil})

	menu.Option("Create a new triblet", 0, true, nil)
	menu.Option("Find a triblet", 1, false, nil)
	menu.Option("Update a trible", 2, false, nil)
	menu.Option("Delete a trible", 3, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {

		log.Fatal(menuerr)

	}

}

func handleFunc(db *sql.DB, opts []wmenu.Opt) {

	switch opts[0].Value {

		case 0:

			fmt.Println("**************************")
			fmt.Println("Creating a new triblet")
			fmt.Println("**************************")
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter a short description: ")
			description, _ := reader.ReadString('\n')
			if description != "\n" {

				description = strings.TrimSuffix(description, "\n")

			}
			fmt.Print("\nEnter the command to run: ")
			runCmd, _ := reader.ReadString('\n')
			if runCmd != "\n" {

				runCmd = strings.TrimSuffix(runCmd, "\n")

			}
			fmt.Print("\nEnter the message to show to the user.\nIt should provide context for the event they'll be working on: ")
			theMsg, _ := reader.ReadString('\n')
			if theMsg != "\n" {

				theMsg = strings.TrimSuffix(theMsg, "\n")

			}
			fmt.Print("\nEnter command for service to restart.\n[Press enter for none] ")
			needRestart, _ := reader.ReadString('\n')
			if needRestart != "\n" {

				needRestart = strings.TrimSuffix(needRestart, "\n")

			}
			fmt.Print("\nProgram requires root?\n[Press Enter for 'no' or type 'yes'] ")
			needRoot, _ := reader.ReadString('\n')
			if needRoot != "\n" {

				needRoot = strings.TrimSuffix(needRoot, "\n")

			}

			newTriblet := theTriblet {

				desc:		description,
				cmd:		runCmd,
				msg:		theMsg,
				restart:	needRestart,
				root:		needRoot,

			}

			addTriblet(db, newTriblet)

			break

		case 1:
			fmt.Println("Searching for a triblet")
		case 2:
			fmt.Println("Updating a triblet")
		case 3:
			fmt.Println("Deleting a triblet")

	// End switch statement
	}
// end handleFunc
}

func addTriblet(db *sql.DB, newTriblet theTriblet){

		// Insert the new triblet
		// Prepare the statement
	        stmt, _ := db.Prepare("INSERT INTO triblets (id, desc, cmd, restart, root) VALUES (?, ?, ?, ?, ?)")

		// Execute the query statements
	        stmt.Exec(nil, newTriblet.desc, newTriblet.cmd, newTriblet.restart, newTriblet.root)
	        defer stmt.Close()

	        fmt.Printf("Added %v\n", newTriblet.desc)

}

func checkErr(err error) {

	if err != nil {

		log.Fatal(err)

	}

}
