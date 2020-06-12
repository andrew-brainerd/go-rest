package main

// Person Type
type Person struct {
	Name string `json:"name" binding:"required"`
}
