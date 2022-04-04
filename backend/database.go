// make a struct and append new users information to a JSON file
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type User struct {
	Name     string
	Messages string
	Location string
	Time     string
}

func ToIds(name string, messages string, location string, time string) {
	user := User{
		Name:     name,
		Messages: messages,
		Location: location,
		Time:     time,
	}
	users := []User{}
	file, err := os.Open("database.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &users)
	users = append(users, user)
	json, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("database.json", json, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func GetUsers() []User {
	users := []User{}
	file, err := os.Open("database.json")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(data, &users)
	return users
}

func main() {
	http.ListenAndServe(":8080", nil)
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
	})
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			messages := r.FormValue("messages")
			location := r.FormValue("location")
			time := r.FormValue("time")
			ToIds(name, messages, location, time)
		}
	})
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Users: %s\n", GetUsers())
	})
}
