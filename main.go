package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
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

func filename(name string) string {
	if strings.HasPrefix(name, "./") || strings.HasPrefix(name, "../") {
		log.Fatalf("Kirill uhodi!")
	}
	return "data/" + name + ".txt"
}

func open() {
	var file1 string
	fmt.Println("What file?")
	fmt.Scan(&file1)

	file, err := os.Open(filename(file1))

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

	fmt.Println("What name?")
	fmt.Scan(&name)

	f, err := os.OpenFile(filename(name), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var text string
	fmt.Println("What text?")

	for {
		text = ""
		fmt.Scanln(&text)

		// fmt.Printf(">%s<\n", text)

		if text == "" {
			break
		}

		_, err = f.WriteString(text + "\n")
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("")
}

func remove() {
	var file string

	fmt.Println("What file?")
	fmt.Scan(&file)

	os.Remove(filename(file))

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	fmt.Println("File delete successfully.")
}

func runCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)

	// var out bytes.Buffer
	// cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(out.String())
}

func save() {
	// runCommand("git", "status")
	runCommand("git", "add", "-A")
	runCommand("git", "commit", "-m", `"SistemCommit"`)
	runCommand("git", "push")
	// runCommand("git", "status")
	fmt.Println("Saved")
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
			remove()
		} else {
			fmt.Println("Command is not valit")
		}
	}
}
