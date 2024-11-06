package view

import (
	"database/sql"
	"fmt"

	"github.com/project-aplikasi-restorant/repository"
	"github.com/project-aplikasi-restorant/service"
	"github.com/project-aplikasi-restorant/utils"
)

func Login(db *sql.DB) {
	userService := service.UserService{
		UserRepo: &repository.UserRepository{DB: db},
	}

	var username, password string
	fmt.Println("Masukkan username:")
	fmt.Scanln(&username)
	fmt.Println("Masukkan password:")
	fmt.Scanln(&password)

	user, err := userService.Login(username, password)
	if err != nil {
		utils.DisplayError("Login gagal:", err)
		return
	}

	fmt.Printf("Selamat datang, %s! Anda login sebagai %s.\n", user.Username, user.Role)
	DisplayMenu(user.Role)
}

func DisplayMenu(s string) {
	panic("unimplemented")
}

// Menambahkan tampilan berdasarkan role, misalnya DisplayMenu untuk admin, koki, atau pelanggan
