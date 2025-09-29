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
	defer db.CloseDB()
	cmd.Execute()
}
