package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"
)

// func main() {
// 	// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing\n")
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//
// 	tmpl := template.Must(
// 		template.New("example").Parse("Welcome, {{.name}}! How are you doing\n"),
// 	) // This will automaticall panic
//
// 	// Define data for the message
// 	data := map[string]any{ // Recommend by lsp
// 		"name": "Adam",
// 	}
//
// 	err := tmpl.Execute(os.Stdout, data)
// 	if err != nil {
// 		panic(err)
// 	}
// }

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")
	nameString, _ := reader.ReadString('\n')
	nameString = strings.TrimSpace(nameString)
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! to the text template\n",
		"notification": "{{.name}}, you have notification {{.notification}}\n",
		"error":        "Error occured: {{.error}}\n",
	}

	parseTemplate := make(map[string]*template.Template)
	for k, v := range templates {
		parseTemplate[k] = template.Must(template.New(k).Parse(v))
	}

	for {
		fmt.Println("Menu")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]any
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parseTemplate["welcome"]
			data = map[string]any{"name": nameString}
		case "2":
			fmt.Println("Enter your notification:")
			notifString, _ := reader.ReadString('\n')
			notifString = strings.TrimSpace(notifString)
			tmpl = parseTemplate["notification"]
			data = map[string]any{"name": nameString, "notification": notifString}
		case "3":
			fmt.Println("Enter your error:")
			errorString, _ := reader.ReadString('\n')
			errorString = strings.TrimSpace(errorString)
			tmpl = parseTemplate["error"]
			data = map[string]any{"error": errorString}
		case "4":
			fmt.Println("Bye")
			return
		default:
			fmt.Println("Invalid input")
			continue
		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println(err)
		}

	}
}
