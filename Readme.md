# MyMandap GoLang Assignment

A go based command line messaging app that allows users to register, send messages between two users, broadcast messages to all users, view message log of a user, display logs and create a log file for each user.

## Features

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit

## How to run

```bash
go run .
```

## Sample Output

```bash
$ go run .
Welcome to CLI Messaging App

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 1
Enter your name: uday
Registered successfully. Your ID is: 1

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 1
Enter your name: notuday
Registered successfully. Your ID is: 2

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 2
Enter sender ID: 1
Enter receiver ID: 2
Enter message content(or press Enter to send a random fact): 
User 1(uday) sent message to User 2(notuday)

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 2
Enter sender ID: 2
Enter receiver ID: 1
Enter message content(or press Enter to send a random fact): some test
 message
User 2(notuday) sent message to User 1(uday)

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 1
Enter your name: yetanotheruday
Registered successfully. Your ID is: 3

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 3
Enter sender ID: 3
Enter broadcast message content(or press Enter to send a random fact):
 broadcasting a message - test
User 3(yetanotheruday) sent message to User 1(uday)
User 3(yetanotheruday) sent message to User 2(notuday)

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 4
Enter user ID: 1
Message Log of User 1(uday):
You sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
User 2(notuday) sent message to User 1(uday): some test message       
User 3(yetanotheruday) sent message to User 1(uday): broadcasting a message - test

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 4
Enter user ID: 2
Message Log of User 2(notuday):
User 1(uday) sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
You sent message to User 1(uday): some test message
User 3(yetanotheruday) sent message to User 2(notuday): broadcasting a message - test

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 4
Enter user ID: 3
Message Log of User 3(yetanotheruday):
You sent message to User 1(uday): broadcasting a message - test       
You sent message to User 2(notuday): broadcasting a message - test    

1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 5
User 1(uday) log:
You sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
User 2(notuday) sent message to User 1(uday): some test message       
User 3(yetanotheruday) sent message to User 1(uday): broadcasting a message - test

User 2(notuday) log:
User 1(uday) sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
You sent message to User 1(uday): some test message
User 3(yetanotheruday) sent message to User 2(notuday): broadcasting a message - test

User 3(yetanotheruday) log:
You sent message to User 1(uday): broadcasting a message - test       
You sent message to User 2(notuday): broadcasting a message - test    


1. Register
2. Send Message between two users
3. Broadcast Message to all users
4. View Message Log of a User
5. Display Logs and create a log file for each user
6. Exit
Enter your choice: 6
User 2(notuday) log:
User 1(uday) sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
You sent message to User 1(uday): some test message
User 3(yetanotheruday) sent message to User 2(notuday): broadcasting a message - test

User 3(yetanotheruday) log:
You sent message to User 1(uday): broadcasting a message - test       
You sent message to User 2(notuday): broadcasting a message - test    

User 1(uday) log:
You sent message to User 2(notuday): A cat has two vocal chords, and can make over 100 sounds.
User 2(notuday) sent message to User 1(uday): some test message       
User 3(yetanotheruday) sent message to User 1(uday): broadcasting a message - test
```