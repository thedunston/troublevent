package main

import (
	"io/ioutil"
	"io"
	"bytes"
	"log"
	"os"
	"os/exec"
	"fmt"
	"github.com/spf13/viper"
)

// Renames the original file with .troublevent extension
func changeFilename(filename string, newFilename string) {

	// Open original file
	originalFile, err := os.Open(filename)
	if err != nil {

		log.Fatal(err)

	}

	defer originalFile.Close()

	// Create new file
	newFile, err := os.Create(newFilename)
	if err != nil {

		log.Fatal(err)

	}

	defer newFile.Close()

	// Copy the bytes to destination from sources
	_, err = io.Copy(newFile, originalFile)
	if err != nil {

		log.Fatal(err)

	}
	//log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {

		log.Fatal(err)

	}

	return

// end changeFilename function
}

// Function to check if file exists
func fileExists(filename string) {

        // Stat returns file info. It will return
        // nil if the file does not exist.
        _, err := os.Stat(filename + ".troublevent")

        if err == nil {

		fmt.Printf("\nBackup file exists. Renaming to the original file => " + filename + "\n\n")

		// Delete the .troublevent file
		err = os.Remove(filename)

		if err != nil {

			log.Fatal(err)

		}
		// Rename the backup file to the original filename
		changeFilename(filename + ".troublevent", filename)

        }

        return

}

// Function to rename from backup to original file.
func renameOrigFile(filename string, newPath string) {

	// Rename the file 
	err := os.Rename(filename, newPath)

	if err != nil {

		log.Fatal(err)

	}

	return

}

func main() {

	// Config
	viper.SetConfigName("replace") // config file name without extension
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
	toSearch := viper.GetString("toSearch")
	toReplaceWith := viper.GetString("toReplaceWith")
	toRestart := viper.GetString("toRestart")
	restartService := viper.GetString("theService")
	theMsg := viper.GetString("theMsg")

	// File to edit.
	input, err := ioutil.ReadFile(theFile)
	if err != nil {

		log.Fatalln(err)
	}

	// Check to see if the user wants to undo changes
	if os.Args[1] == "undo" {

		// Remove the file that is broken
		os.Remove(theFile)

		if err != nil {

			log.Fatal(err)

		}

		// Replace with the original file
		changeFilename(theFile + ".troublevent",theFile)

		// Delete the triblet backup	
		os.Remove(theFile + ".troublevent")
		if err != nil {

			log.Fatal(err)

		}

		fmt.Printf("\nOriginal file restored => " + theFile + "\n")

		os.Exit(0)

	}

	// Check if file exists
	fileExists(theFile)

	// Create a backup of the original file
	changeFilename(theFile, theFile + ".troublevent")

	// Replace the content
	output := bytes.Replace(input, []byte(toSearch), []byte(toReplaceWith), -1)

	// Check for errors
	err = ioutil.WriteFile(theFile, []byte(output), 0644)
	if err != nil {

	        log.Fatalln(err)

	}

	fmt.Println(toRestart)
	fmt.Println(restartService)
	if toRestart == "yes" {

		cmd := exec.Command(restartService)

		_, err := cmd.StdinPipe()
		if err != nil {

			log.Fatal(err)

		}

	}

	// If no errors, print the message
	fmt.Println("\n**************************\n")
	fmt.Println(theMsg)
	fmt.Println("\n**************************\n")

}
