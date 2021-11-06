/**

A command will be executed without the Pipe or with one.

The yaml variable thePipe is checked to see if it is empty.  If it is then it will execute a single command.

Otherwise. it will invoke a pipe using Go.  Note that the command to run should not have the pipe in it.

Pipe code reference: https://gist.github.com/ochinchina/9e409a88e77c3cfd94c3

*/

package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"github.com/spf13/viper"
	"strings"
	"bytes"
	"embed"

)

/* The below function checks if a regular file (not directory) with a
   given filepath exist 
   https://www.algotree.org/algorithms/snippets/go_check_if_file_exists/ */
func configFileExists (filepath string) bool {

	fileinfo, err := os.Stat(filepath)

	if os.IsNotExist(err) {
		return false
					               }
	// Return false if the fileinfo says the file path is a directory.
	return !fileinfo.IsDir()
}

func check(e error) {

	if e != nil {

		panic(e)

        }

}

//go:embed "command.yaml"
var g embed.FS

func main() {

	// Get embedded yaml file.
        data, err := g.ReadFile("command.yaml")
	check(err)

	var _emcommand = []byte(data)

        viper.SetConfigType("yaml")
        viper.ReadConfig(bytes.NewBuffer(_emcommand))

	// Get the triblet commands:
	theCmd := viper.GetString("Command.Cmd")
	thePipe := viper.GetString("Command.Pipe")
	theMsg := viper.GetString("Command.msg")

	fmt.Println(len(thePipe))
	// If there is a command to pipe
	if (len(thePipe) != 0) {

		// Split the command using spaces
		args := strings.Fields(theCmd) 
	
		// Split the command using spaces
		args2 := strings.Fields(thePipe) 
	
		// Execute the commands
		// args[0] runs the main executable
		// args[1:]... gets the options and switches
		cmd1 := exec.Command(args[0], args[1:]...)
		cmd2 := exec.Command(args2[0], args2[1:]...)
	
		// Create a pipe.
	        reader, writer := io.Pipe()
	
	        // First command
	        cmd1.Stdout = writer
	
	        // second part of command
	        cmd2.Stdin = reader
	
	        // prepare a buffer to capture the output
	        // after second command finished executing
	        var buff bytes.Buffer
	
		// Stores the output of the command from Pipe when it finishes.
	        cmd2.Stdout = &buff
	
		// Run the first command: Cmd
	        cmd1.Start()
	
		// Run the second command: Pipe
	        cmd2.Start()
	
		// Waits for the first command to finish running
	        cmd1.Wait()
	
		// Close the writer
	        writer.Close()

		// Waits for the second command to complete
	        cmd2.Wait()

		// Close the reader
		reader.Close()

		// Converting the command results to a string
	        output := buff.String()

		// Print the results
		fmt.Printf("%s", output)
		fmt.Println("\n\n************************************************************")
		fmt.Printf("\n%s\n\n", theMsg)
		fmt.Println("************************************************************\n\n")

	// Otherwise run the one command.
	} else {

		if (strings.Contains(theCmd, " ")) {

		// Split the command string
		args := strings.Fields(theCmd)
		// Execute the command.
		cmd := exec.Command(args[0], args[1:]...)

		// Get command output
		stdout, err := cmd.Output()
		check(err)

		fmt.Printf("%s\n", stdout)

	} else {


		// Execute the command.
		cmd := exec.Command(theCmd)
		
		// Get command output
		stdout, err := cmd.Output()
		check(err)

		fmt.Printf("%s\n", stdout)
	
	}

		// Print the message
		fmt.Println("************************************************************")
		fmt.Printf("\n%s\n\n", theMsg)
		fmt.Println("************************************************************")

	}

}
