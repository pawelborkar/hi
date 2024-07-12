package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	commands := map[string]string{
		"pn":   "pnpm",
		"pa":   "pnpm add",
		"pad":  "pnpm add -D",
		"puu":  "pnpm update; pnpm upgrade",
		"pag":  "pnpm add -g ",
		"pagu": "pnpm add -g pnpm",
		"pi":   "pnpm install",
		"pd":   "pnpm run dev",
		"ps":   "pnpm run start",
		"psv":  "pnpm run serve",
		"gi":   "git init",
		"gs":   "git status",
		"gl":   "git log --graph --decorate --all",
		"dasp": "django-admin startproject",
		"drs":  "python manage.py runserver",
		"dmm":  "python manage.py makemigrations",
		"dm":   "python manage.py migrate",
		"dt":   "python manage.py test",
		"dsh":  "python manage.py shell",
		"dc":   "python manage.py check",
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: hi <command>")
		os.Exit(1)
	}

	commandKey := os.Args[1]

	fullCommand, exists := commands[commandKey]

	// Give back the lastest version of hi
	if (commandKey == "-v") || (commandKey == "version") || (commandKey == "--version") {
		fmt.Println("version 0.3.1")
		return
	}

	if commandKey == "-add" {

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the command (e.g. 'ga: git add -A'): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.SplitN(input, ":", 2)

		if len(parts) == 2 {
			alias := strings.TrimSpace(parts[0])
			command := strings.TrimSpace(parts[1])

			if _, exists := commands[alias]; exists {
				fmt.Println("Alias already exists.")
				return
			}

			commands[alias] = command
			fmt.Println("Alias added successfully.")
			return
		} else {
			fmt.Println("Invalid input!")
		}

	}

	if !exists {
		fmt.Printf("Command '%s' not found. Try again or add a new command \n", commandKey)
		os.Exit(1)
	}

	if len(os.Args) == 3 {
		fullCommand = fullCommand + " " + os.Args[2]
		fmt.Printf("Executing %s\n", fullCommand)
		executeCommand(fullCommand)
	} else {
		fmt.Printf("Executing %s\n", fullCommand)
		executeCommand(fullCommand)
	}
}

func executeCommand(command string) {

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
