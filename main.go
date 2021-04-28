package main

import (
	"bufio"
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"
)

type smtpServer struct {
	host, port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	//sender data
	from, to, password, message, err := getInput()

	if err != nil {
		fmt.Println(fmt.Errorf("Invalid input"))
		return
	}
	//server configuration
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}

	//authentication
	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	fmt.Println("Sending your email.....")
	go spinner(100 * time.Millisecond)

	//sending the email
	err = smtp.SendMail(smtpServer.Address(), auth, from, to, []byte(message))
	fmt.Printf("\n")
	if err == nil {
		fmt.Println("Email has been sent successfuly !")
	} else {
		fmt.Println("Some error was encounterd !")
		fmt.Println(err)
	}
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

//takes user input for the email
func getInput() (from string, to []string, password string, msg string, err error) {
	err = nil
	line := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("From: ")
	from, err = reader.ReadString('\n')
	from = strings.Replace(from, "\n", "", -1)
	if err != nil {
		return
	}
	fmt.Println("To: ")
	for {
		line, err = reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		if err != nil {
			return
		}
		if line == "{{END}}" {
			break
		}
		to = append(to, line)
	}
	fmt.Println("Password: ")
	password, err = reader.ReadString('\n')
	password = strings.Replace(password, "\n", "", -1)
	if err != nil {
		return
	}
	fmt.Println("Message: ")
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			return
		}
		if line == "{{END}}\n" {
			break
		}
		msg = msg + line
	}
	return
}
