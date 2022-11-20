package main

import "fmt"

type User struct {
	name string
	age  int
}

type ChatBot interface {
	getBotVersion() string
	respondToUser(user User) string
}

type EnglishChatBot struct {
	version string
}

type MarathiChatBot struct {
	version string
}

func (bot EnglishChatBot) getBotVersion() string {
	return bot.version
}

func (bot EnglishChatBot) respondToUser(u User) string {
	return fmt.Sprintf("HELLO - ", u, "How are you")
}

func (bot MarathiChatBot) getBotVersion() string {
	return bot.version
}

func (bot MarathiChatBot) respondToUser(u User) string {
	return fmt.Sprintf("NAMASTE - ", u, "Kasa Ahes Bhau")
}

func main() {
	var eng = EnglishChatBot{}
	var mar = MarathiChatBot{}

	fmt.Println("The version of ", eng, " is ", eng.getBotVersion())
	fmt.Println("The version of ", mar, " is ", mar.getBotVersion())

	fmt.Println(eng.respondToUser(User{name: "HEMANT"}))
	fmt.Println(mar.respondToUser(User{name: "HEMANT"}))
}
