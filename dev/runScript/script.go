/**

This will read in a script from the script.yaml file and execute it.

- TODO:
Add a check to see if the user is admin, if required.
Need to test on Windows with Powershell script and batch files

*/

package main

import (

	"io"
	"log"
	"os"
	"os/exec"
	"fmt"
	"github.com/spf13/viper"

)

func check(e error) {

	if e != nil {
		panic(e)
	}

}

func main() {

	// Config
	viper.SetConfigName("script") // config file name without extension
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()             // read value ENV variable

	err := viper.ReadInConfig()

	if err != nil {

		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	// Get the triblets:
	theFile := viper.GetString("theFile")
	theShell := viper.GetString("theShell")
	theScript := viper.GetString("theScript")
	theMsg := viper.GetString("theMsg")

	check(err)
	// Create the script and make it executable
	f, err := os.OpenFile(theFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0700,)

	check(err)

	defer f.Close()

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

	// If no errors, print the message
	fmt.Println("\n**************************\n")
	fmt.Println(theMsg)
	fmt.Println("\n**************************\n")

}
