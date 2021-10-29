/**

This will replace text within a configuration file.

First, a backup of the original file is created with the extension ".troublevent"

Then the edit is made to the configuration file.

Passing the argument:

undo

after the script will replace the original configuration file and delete the file with the
.troublevent extension.

./replace undo

*/

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
	"strings"
	"embed"
)

/* The below function checks if a regular file (not directory) with a given filepath exist */
 
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

//go:embed "replace.yaml"
var g embed.FS

func main() {

	// Get embedded yaml file.
        data, err := g.ReadFile("replace.yaml")
        check(err)

	var _emcommand = []byte(data)

        viper.SetConfigType("yaml")
        viper.ReadConfig(bytes.NewBuffer(_emcommand))

	// Get the triblets:
	theFile := viper.Get("theFile").(string)
	toSearch := viper.Get("toSearch").(string)
	toReplaceWith := viper.Get("toReplaceWith").(string)
	toRestart := viper.Get("toRestart").(string)
	restartService := viper.Get("theService").(string)
	theMsg := viper.Get("theMsg").(string)

	// File to edit.
	input, err := ioutil.ReadFile(theFile)
	if err != nil {

		log.Fatalln(err)
	}

	// Check to see if the user wants to undo changes

	if len(os.Args) > 1 {

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

	if toRestart == "yes" {

		// Split the command using spaces to restart the service
                args := strings.Fields(restartService)

		cmd := exec.Command(args[0], args[1:]...)

		// .Run allows restarting network systemd scripts
		err := cmd.Run()
		if err != nil {

			log.Fatal(err)

		}

	}

	// If no errors, print the message
	fmt.Println("\n**************************\n")
	fmt.Println(theMsg)
	fmt.Println("\n**************************\n")

}
