package main

import "time"

type Task struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Done  bool      `json:"done"`
	Date  time.Time `json:"date"`
}

const dataFile = "tasks.json"
