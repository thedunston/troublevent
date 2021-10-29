/**

This will read in a script from the script.yaml file and execute it.

- TODO:
Add a check to see if the user is admin, if required.
Need to test on Windows with Powershell script and batch files

*/

package main

import (

	"os"
	"os/exec"
	"os/signal"
	"fmt"
	"github.com/spf13/viper"
	_ "embed"
	"embed"
	"bytes"
	"syscall"

)

func check(e error) {

	if e != nil {

		panic(e)
	}

}

/* Embeds the yaml file so the binary can be portable. */
//go:embed "script.yaml"
var g embed.FS

func main() {


	// Get embedded yaml file.
	data, err := g.ReadFile("script.yaml")
	check(err)

	var _emscript = []byte(data)

        viper.SetConfigType("yaml")
        viper.ReadConfig(bytes.NewBuffer(_emscript))

	// Get the triblets
	// Convert values to a string
	theFile := viper.Get("theFile").(string)
	theShell := viper.Get("theShell").(string)
	theScript := viper.Get("theScript").(string)
	theMsg := viper.Get("theMsg").(string)

	// Message to the user so the don't get impatient.
	// if the script runs a while.
	fmt.Println("Please wait, this may take time to run...")
	//err = os.Chmod(theFile, 0700)
	check(err)

	// Create file
	f, err := os.Create(theFile)
	check(err)

	// Write the script to the file.
        _, err2 := f.Write([]byte(theScript))

        check(err2)
        defer f.Close()

        // Message to the user so the don't get impatient.
        // if the script runs a while.
        fmt.Println("Please wait, this may take time to run...")
        //err = os.Chmod(theFile, 0700)
        check(err)

	// Execute the command
	cmd := exec.Command(theShell, theFile)
	err = cmd.Run()
	check(err)
	fmt.Printf("Command finished with error: %v", err)

	// If no errors, print the message
	fmt.Println("\n**************************\n")
	fmt.Println(theMsg)
	fmt.Println("\n**************************\n")


}
