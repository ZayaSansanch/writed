package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func comands() {
	fmt.Println("Comands:")
	fmt.Println("	h")
	// fmt.Println("	open", "- open[файл]")
	fmt.Println("	open")
	fmt.Println("	save")
	fmt.Println("	write")
	// fmt.Println("	ls")
	fmt.Println("	remove")
	fmt.Println("")
}

func open() {
	var file1 string
	fmt.Println("Wat file?")
	fmt.Scan(&file1)

	file, err := os.Open(file1)

	fmt.Println("")

	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

	fmt.Println("")
}

func write() {
	var name string

	fmt.Println("Wat name?")
	fmt.Scan(&name)

	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var text string
	fmt.Println("Wat text?")
	fmt.Scan(&text)

	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}
}

func delet() {
	var file string

	fmt.Println("Wat file?")
	fmt.Scan(&file)

	os.Remove(file)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("File delete successfully.")
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())
}

func save() {
	// runCommand("git", "status")
	runCommand("git", "add", "-A")
	runCommand("git", "commit", "-m", `"SistemCommit"`)
	runCommand("git", "push")
	// runCommand("git", "status")
}

func main() {
	var comand string

	fmt.Println("Its writed!")
	for comand != "save" {
		fmt.Scan(&comand)
		fmt.Println("")
		if comand == "h" {
			comands()
		} else if comand == "open" {
			open()
		} else if comand == "save" {
			save()
			break
		} else if comand == "write" {
			write()
		} else if comand == "remove" {
			delet()
		} else {
			fmt.Println("Command is not valit")
		}
	}
}
