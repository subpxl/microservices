package main

type Todo struct {
	ID     int    `json:"id"`
	Todo   string `json:"todo"  binding:"required"`
	Status bool   `json:"status"`
}
