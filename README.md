# troublevent

Troublevent is a program to create scenarios to help learn troubleshooting and interpreting system admin tasks. It is designed for schools, training providers, etc.


## When to use troublevent?

First, each program created is a called a triblet (I needed a cool name) and is a compiled binary executable.

If you are an educator teaching a system administration course, you can edit a configuration file with a typo and restart the service and the student has to figure out the problem.

You may want users to be able to ensure they are typing the right command to give them the expected output.  You can add a command that prints out specific output and the student has to explain which command was executed.

## Setup

This program uses Golang and the Viper module.

You will need to download Go.

Download the code for the project and cd into the directory

`git clone https://github.com/thedunston/troublevent.git`

`cd troublevent`

`go get github.com/spf13/viper`


then run:

`go mod init goViper`

`go tidy`
 

## Creating a triblet

You can create a triblet by copying the fileediting directory or run a command by copying the newcommand directory template.  It is suggested to rename it so it is descriptive.  Read the comments and example in the YAML file. The YAML file contains the actions to perform and a message to the user.  It is recommended to test the binary before building the executable.

`go run command.go`

and then after testing it build it to an executable:

`go build command.go`

It is recommended to provide a name so it is easier to know which binary it is.

`go build -o LabNameAndNumberGoesHere.exe`

# Sample Triblets:

### Edit the /etc/dhcp/dhcpd.conf file and change "netmask" to netmasks."  It is best to have the configuration change as real as possible so it simulates a real-world mistake.

`cp -rp fileedit Week2Lab3`

`cd dhcpd`

`nano replace.yaml`


     ---
      #Filename to search and replace
      theFile: "/etc/dhcp/dhcpd.conf"
      #Text to search for
      toSearch: "netmask"
      #Text to replace.  For real-world context, make it something the user is likely
      #to mistype such as "netmaks" instead of "netmask."
      #You can also leave off a semicolon or a curly brace.
      toReplaceWith: "netmasks"
      #Whether or not a service needs to be restarted
      #yes or no
      toRestart: "yes"
      #Service to restart command
      theService: "systemctl restart named"

      #Or just ask the user what command was used to generate the output
      theMsg: "You just edited the /etc/dhcp/dhcpd.conf file and restarted the service. Check if users can get a DHCP address, if not fix the issue." 
 

save the file

`go run replace.go`

Fix any issues and then build.

`go build -o Week2Lab3.bin`

### Create a binary that runs ps -ef and the student has to determine the command that printed the output.

`cp -rp newcommand Question1Week8Lab`

`cd ps`

`nano command.yaml`

     ---
       Cmd: "ps -ef"

Save the file and run:

`go run command.go`

Fix any issues and then build.

`go build -o Question1Week8Lab.bin`

### On Windows Create a binary that runs the Get-Process cmdlet.
You can also use the pipe and where-object (where) cmdlet.  
Add single quotes instead of double quotes

`xcopy /E /I newcommand Question1Week9Lab`

`cd ps`

`notepad command.yaml`

     ---
       # Invoke powershell and run the get-process command
       Cmd: "powershell get-Process "
       # A pipe can be used.
       # Cmd: powershell get-process | where { $_.ProcessName -eq 'chrome' }

Save the file and run:

`go run command.go`

Fix any issues and then build.

`go build -o Question1Week9Lab.exe`
