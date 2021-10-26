package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"github.com/spf13/viper"
	"strings"
	"bytes"

)

func main() {

	// Config file. Don't add the file extension.
        viper.SetConfigName("command")
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")

        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }

	// Get the triblet commands:
	theCmd := viper.GetString("Cmd")
	thePipe := viper.GetString("Pipe")
	theMsg := viper.GetString("msg")


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
	
		// Split the command string
		args := strings.Fields(theCmd)

		// Execute the command.
		cmd := exec.Command(args[0], args[1:]...)

		// Create stdin pipe to send results
		stdin, err := cmd.StdinPipe()

		if err != nil {
	
			log.Fatal(err)
	
		}
	
		go func() {
	
			defer stdin.Close()
	
			io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	
		}()

	        out, err := cmd.CombinedOutput()
	        if err != nil {
	                log.Fatal(err)
	        }
																	        fmt.Printf("%s\n", out)
	
		// Print the message
		fmt.Println("************************************************************")
		fmt.Printf("\n%s\n\n", theMsg)
		fmt.Println("************************************************************")

	}


}
