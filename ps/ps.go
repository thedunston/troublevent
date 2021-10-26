package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"github.com/spf13/viper"
	"strings"

)

func main() {

	// Config
        viper.SetConfigName("ps") // config file name without extension
        viper.SetConfigType("yaml")
        viper.AddConfigPath(".")
        viper.AutomaticEnv()             // read value ENV variable

        err := viper.ReadInConfig()
        if err != nil {
                fmt.Println("fatal error config file: default \n", err)
                os.Exit(1)
        }
	// Get the triblet:
	theCmd := viper.GetString("cmd")


	args := strings.Fields(theCmd) //or any similar split function
	log.Println(theCmd)
	//exec.Command(args[0], args[1:]...)
	cmd := exec.Command(args[0], args[1:]...)
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

}
