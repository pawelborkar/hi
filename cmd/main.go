package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"os/exec"
	"strings"
)

func main() {
	commands := map[string]string{
		"fd":   "fastapi dev",
		"cm":   "mf",
		"fget": "~/github/fastget/bin/main",
		"l3":   "xdg-open http://localhost:3000/",
		"l5":   "xdg-open http://localhost:5000/",
		"l55":  "xdg-open http://localhost:5050/",
		"l8":   "xdg-open http://localhost:8000/",
		"pa":   "pnpm add",
		"pad":  "pnpm add -D",
		"puu":  "pnpm update; pnpm upgrade",
		"pag":  "pnpm add -g ",
		"psu":  "pnpm self-update",
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
		"dsa":  "python manage.py startapp",
		"dsh":  "python manage.py shell",
		"dt":   "python manage.py test",
		"dc":   "python manage.py check",
		"dpm":  "python manage.py",
		"dcu":  "docker compose up",
		"dps":  "docker ps",
		"dimg": "docker images",
		"epd":  "export PKG_CONFIG_PATH=/usr/lib/x86_64-linux-gnu/pkgconfig:$PKG_CONFIG_PATH",
	}

	if len(os.Args) < 2 {
		color.Yellow("Usage: hi <command>")
		os.Exit(1)
	}

	commandKey := os.Args[1]

	fullCommand, exists := commands[commandKey]

	// Give back the lastest version of hi
	if (commandKey == "-v") || (commandKey == "version") || (commandKey == "--version") {
		color.HiCyan("version 0.3.3")
		return
	}

	if commandKey == "-add" {
		reader := bufio.NewReader(os.Stdin)
		color.HiYellow("Enter the command (e.g. 'ga: git add -A'): ")
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

	// Takes arguments from the user
	if len(os.Args) == 3 {
		fullCommand = fullCommand + " " + os.Args[2]
	}
	color.HiGreen("Executing %s\n", fullCommand)
	executeCommand(fullCommand)
}

func executeCommand(command string) {
	// Executes the command in the shell/terminal

	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()

	if err != nil {
		color.Red("Error executing command: %v\n", err)
		os.Exit(1)
	}
}
