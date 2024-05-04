package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"encoding/json"

	"sync"
)

type User struct {
	ID      int
	Name    string
	LogFile string
	Log     []string
}

var users map[int]*User
var mu sync.Mutex

func main() {
	users = make(map[int]*User)
	defer displayLogs()

	fmt.Println("Welcome to CLI Messaging App")

	for {
		fmt.Println("\n1. Register")
		fmt.Println("2. Send Message between two users")
		fmt.Println("3. Broadcast Message to all users")
		fmt.Println("4. View Message Log of a User")
		fmt.Println("5. Display Logs and create a log file for each user");
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			registerUser()
		case 2:
			sendMessage()
		case 3:
			broadcastMessage()
		case 4:
			viewMessageLog()
		case 5:
			displayLogs()
		case 6:
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func registerUser() {
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	name := scanner.Text()

	mu.Lock()
	defer mu.Unlock()

	userID := len(users) + 1
	logFile := fmt.Sprintf("user_%d.log", userID)
	user := &User{ID: userID, Name: name, LogFile: logFile}
	users[userID] = user

	fmt.Printf("Registered successfully. Your ID is: %d\n", userID)
}

func sendMessage() {
	fmt.Print("Enter sender ID: ")
	var senderID int
	fmt.Scanln(&senderID)

	sender, ok := users[senderID]
	if !ok {
		fmt.Println("Invalid sender ID.")
		return
	}

	fmt.Print("Enter receiver ID: ")
	var receiverID int
	fmt.Scanln(&receiverID)

	receiver, ok := users[receiverID]
	if !ok {
		fmt.Println("Invalid receiver ID.")
		return
	}

	fmt.Print("Enter message content(or press Enter to send a random fact): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	if message == "" {
		message = getRandomFact()
		if message == "" {
			fmt.Println("Error fetching random fact. Please try again.")
			return
		}
	}

	sendMessageToUser(sender, receiver, message)
}

func sendMessageToUser(sender, receiver *User, message string) {
	mu.Lock()
	defer mu.Unlock()

	log := fmt.Sprintf("User %d(%s) sent message to User %d(%s): %s", sender.ID, sender.Name, receiver.ID, receiver.Name, message)
	receiver.Log = append(receiver.Log, log)
	fmt.Printf("User %d(%s) sent message to User %d(%s)\n", sender.ID, sender.Name, receiver.ID, receiver.Name)

	sender.Log = append(sender.Log, fmt.Sprintf("You sent message to User %d(%s): %s", receiver.ID, receiver.Name, message))
}

func broadcastMessage() {
	fmt.Print("Enter sender ID: ")
	var senderID int
	fmt.Scanln(&senderID)

	sender, ok := users[senderID]
	if !ok {
		fmt.Println("Invalid sender ID.")
		return
	}

	fmt.Print("Enter broadcast message content(or press Enter to send a random fact): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	message := scanner.Text()

	if message == "" {
		message = getRandomFact()
		if message == "" {
			fmt.Println("Error fetching random fact. Please try again.")
			return
		}
	}

	for _, receiver := range users {
		if receiver != sender {
			sendMessageToUser(sender, receiver, message)
		}
	}
}

func viewMessageLog() {
	fmt.Print("Enter user ID: ")
	var userID int
	fmt.Scanln(&userID)

	user, ok := users[userID]
	if !ok {
		fmt.Println("Invalid user ID.")
		return
	}

	fmt.Printf("Message Log of User %d(%s):\n", user.ID, user.Name)
	for _, log := range user.Log {
		fmt.Println(log)
	}
}

func displayLogs() {
	mu.Lock()
	defer mu.Unlock()

	for _, user := range users {
		fmt.Printf("User %d(%s) log:\n", user.ID, user.Name)
		for _, log := range user.Log {
			fmt.Println(log)
		}
		fmt.Println()
		saveLog(user.LogFile, user.Log)
	}
}

func saveLog(filename string, log []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for _, entry := range log {
		_, err := file.WriteString(entry + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}

func getRandomFact() string {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		log.Println("Error fetching random fact:", err)
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return ""
	}

	var result map[string]interface{}
	err = json.Unmarshal([]byte(body), &result)

	if err != nil {
		log.Println("Error unmarshalling response body:", err)
		return ""
	}

	fact := result["fact"].(string)
	return fact
}
