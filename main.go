package main

import "github.com/nelsongp/goSimpleWs/handlers"

// Main function
func main() {
	handlers.Manejadores()
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
