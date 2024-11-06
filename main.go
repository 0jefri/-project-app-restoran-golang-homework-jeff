package main

import (
	"fmt"

	"github.com/project-aplikasi-restorant/utils"
)

func main() {
	db := utils.InitDB()
	defer db.Close()

	fmt.Println("Selamat datang di Aplikasi Restoran")
	view.Login(db)
}
