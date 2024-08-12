package main

import (
	"bufio"
	"log"
	"os/exec"
)

func main() {
	log.Println("Running command /bin/ls -l")
	cmd := exec.Command("/bin/ls", "-l")
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Println(err)
	}

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		log.Println(err)
	}

	// read command's stdout line by line
	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf(in.Text()) // write each line to your log, or anything you need
	}
	if err := in.Err(); err != nil {
		log.Printf("error: %s", err)
	}

}
