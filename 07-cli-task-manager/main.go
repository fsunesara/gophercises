/*
Copyright © 2025 fsunesara
*/
package main

import (
	"task/cmd"
	"task/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
	db.CloseDB()
}
