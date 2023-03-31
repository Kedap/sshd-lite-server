package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	sshd "github.com/Kedap/sshd-lite-server"
)

func main() {
	c := &sshd.Config{
		Host:       "0.0.0.0",
		Port:       "7890",
		Shell:      os.Getenv("SHELL"),
		KeyFile:    "",
		KeySeed:    "",
		AuthType:   "",
		KeepAlive:  60,
		IgnoreEnv:  false,
		LogVerbose: false,
	}
	c.AuthType = "john:doe"
	s, err := sshd.NewServer(c)
	if err != nil {
		fmt.Println("Error!")
		os.Exit(1)
	}
	q := make(chan bool)
	go s.Start(q)
	var stop string
	for {
		fmt.Scan(&stop)
		if strings.Contains(stop, "stop") {
			fmt.Println("stoping server")
			q <- true
			break
		}
	}
	fmt.Println("Wait 20s")
	time.Sleep(time.Second * 20)
	fmt.Println("Main is close")
}
