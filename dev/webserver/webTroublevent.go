package main

import (

	"fmt"
	"log"
//	"github.com/spf13/viper"
	"net/http"
	"io/ioutil"
	yaml "gopkg.in/yaml.v3"


)


func process(w http.ResponseWriter, r *http.Request) {


	// Error if page is not found.
	if r.URL.Path != "/" {

		http.Error(w, "404 not found.", http.StatusNotFound)
		return

	}

	switch r.Method {

	case "GET":

		http.ServeFile(w, r, "command.html")

	case "POST":

		type Record struct {

			CmdWrite string `yaml:"Cmd"`
			PipeWrite string `yaml:"Pipe"`
			MsgWrite string `yaml:"msg"`
		}

		type Config struct {

			Record Record `yaml:"Command"`

		}

		if err := r.ParseForm(); err != nil {

			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return

		}

		theCmd := r.FormValue("Cmd")
		thePipe := r.FormValue("Pipe")
		theMsg := r.FormValue("msg")

		config := Config{Record: Record{CmdWrite: ""+theCmd+"", PipeWrite: thePipe, MsgWrite: theMsg}}


		data, err := yaml.Marshal(&config)
		if err != nil {

			log.Fatalf("error: %v", err)

		}

		err = ioutil.WriteFile("command.yaml", data, 0644)
		if err != nil {

			log.Fatalf("error: %v", err)

		}


	default:

		fmt.Fprintf(w, "Sorry, Only POST methods are supported.")

	} // end switch statement

} // end process function

func main() {

	http.HandleFunc("/", process)

	fmt.Printf("Starting server at port 8082\n")
	log.Fatal(http.ListenAndServe("127.0.0.1:8082", nil))

}
