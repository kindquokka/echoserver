package main

import (
	"log"
	"bufio"
	"os"
	
	"server"
)

const (
	name        = "echoserver"
	description = "My Echo Server"
	port        = ":8080"
)

func ckerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	go func() {
		s := server.NewServer()
		ckerr(s.ListenAndServe(port))
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		switch scanner.Text() {
		case "q":
			return
		}
	}
}






