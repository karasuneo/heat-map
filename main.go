package main

import "fmt"

type User struct {
	name         string
	age          int
	introduction string
}

type Introduction struct {
	introduction *string
}

func main() {
	users := []User{
		{"Alice", 25, "Hello, I'm Alice!"},
		{"Bob", 30, "Hi there, I'm Bob."},
		{"Charlie", 22, "Greetings! I'm Charlie."},
	}
	introductions := makeIntroductions(users)

	for _, introduction := range introductions {
		fmt.Println(introduction.introduction)
	}
}

func makeIntroductions(us []User) []Introduction {
	var introductions []Introduction

	for _, user := range us {
		introduction := user.introduction
		introductions = append(introductions, Introduction{
			introduction: &introduction,
		})
	}

	return introductions
}
