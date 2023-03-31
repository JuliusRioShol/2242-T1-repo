package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"
)

// The handler function:
func h1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("My Name is Julius Shol.	My Email is 2021154969@gmail.com.\n"))
	w.Write([]byte("I am currently persuing my Associates in IT, it has been a challenge but I am almost near the finish Line.\n"))
	w.Write([]byte("I am from a village called San Pedro Columbia, which is in the Toledo District. \n"))
	w.Write([]byte("I love playing football, it is my favourite sport and I am 18 Years old.\n"))
}

func h2(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("3:04 AM") // a variable that stores the time
	currentDay := time.Now().Weekday().String() // A var , stores weekday.
	w.Write([]byte("Right now is "))
	w.Write([]byte(currentTime))
	w.Write([]byte("\nEnjoy the rest of your "))
	w.Write([]byte(currentDay))
}

// store random quotes:
var quotes = []string{
	"Talk Is Cheap. Show Me The Code:)",
	"When You Have A Dream, You've Got To Grab It And Never Let Go:)",
	"First, Solve The Problem. Then, Write The Code:)",
	"Nothing Is Impossible:)",
	"Experience Is The Name Everyone Gives To Their Mistakes:)",
	"There Is Nothing Possible To They Who Will Try:)",
	"The Bad News Is Time Flies:)",
	"Life Has Got All Those Twists And Turns:)",
}

func h3(w http.ResponseWriter, r *http.Request) {
	// Get a random quote
	quote := quotes[rand.Intn(len(quotes))]
	//write the response
	w.Write([]byte(quote))
}
func main() {
	// Create a new serveMux
	mux := http.NewServeMux()
	mux.HandleFunc("/", h1)
	mux.HandleFunc("/greeting", h2)
	mux.HandleFunc("/random", h3)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
