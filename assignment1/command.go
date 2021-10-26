package main

import (
	"fmt"
	"io"
	"log"
	"os"
	//"os/exec"
	"github.com/spf13/viper"
	"strings"
	"pipes"

)

func main() {

	// Config file name without the file extension
        viper.SetConfigName("command")
	// Config file type.  The .yaml will be added to the 
	// name provided above
        viper.SetConfigType("yaml")

	// Path for the yaml file.
        viper.AddConfigPath(".")

	// Open config file
        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }

	// Get the triblet
	theCmd := viper.GetString("cmd")

	// Split the command if it has spaces
	args := strings.Fields(theCmd) //or any similar split function
	//exec.Command(args[0], args[1:]...)
	//cmd := exec.Command(args[0], args[1:]...)
	//stdin, err := cmd.StdinPipe()
	cmd, err := pipes.RunString(args)

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

}
