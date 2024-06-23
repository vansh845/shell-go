package main

import (
	"bufio"
	"os/exec"

	// Uncomment this block to pass the first stage
	"fmt"
	"os"
	"strings"
)

var pathVar = os.Getenv("PATH")

func main() {
	// Uncomment this block to pass the first stage
out:
	for {

		fmt.Fprint(os.Stdout, "$ ")

		// Wait for user input

		inp, _, err := bufio.NewReader(os.Stdin).ReadLine()
		if err != nil {
			panic(err)
		}
		commands := strings.Split(string(inp), " ")

		cmd := commands[0]
	outerSwitch:
		switch cmd {
		case "echo":
			fmt.Fprint(os.Stdout, strings.Join(commands[1:], " ")+"\n")
		case "type":
			temp := commands[1]
		switchCase:
			switch temp {
			case "echo", "type", "exit", "pwd":
				fmt.Printf("%s is a shell builtin\n", temp)
			default:

				dirs := strings.Split(pathVar, ":")
				for _, dir := range dirs {
					dir += fmt.Sprintf("/%s", temp)
					_, err := os.Stat(dir)
					if err != nil {
						if os.IsNotExist(err) {
							continue
						} else {
							panic(err)
						}
					}
					fmt.Printf("%s is %s\n", temp, dir)
					break switchCase

				}
				fmt.Fprintf(os.Stderr, "%s: not found\n", temp)

			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			fmt.Println(pwd)
		case "cd":
			dir := commands[1]
			if dir == "~" {
				os.Chdir(os.Getenv("HOME"))
				break outerSwitch
			}
			if len(commands) > 1 {
				err := os.Chdir(dir)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: No such file or directory\n", dir)
				}
			}
		case "exit":
			break out
		default:

			dirs := strings.Split(pathVar, ":")

			for _, dir := range dirs {
				dir += fmt.Sprintf("/%s", cmd)
				_, err := os.Stat(dir)
				if err != nil {
					if os.IsNotExist(err) {
						continue
					} else {
						panic(err)
					}
				}

				exe := exec.Command(cmd, commands[1:]...)
				buffer, err := exe.Output()
				if err != nil {
					panic(err)
				}
				fmt.Print(string(buffer))
				break outerSwitch

			}
			fmt.Fprintf(os.Stderr, "%s: command not found\n", cmd)

		}

	}

}
