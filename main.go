package main

import (
	"Todo/config"
	"Todo/user"
	"fmt"
)

func main() {
	var err error
	var menu int

	config.Connect()
	db := config.GetConnection()

	for {
		fmt.Println("====Selamat datang====")
		fmt.Println("Silahkan pilih menu")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Out")
		fmt.Scanln(&menu)

		switch menu {
		case 1:
			fmt.Println("Silahkan Register Dude!")
			if err = user.Register(db); err != nil {
				return
			}
		case 2:
			if noHp == "" {
				fmt.Println("Silahkan Login Dude!")
				if noHp, err = user.Login(db); err != nil {
					pilihan = 0
					return
				}

				// // case 0:
				// // 	fmt.Println("Terima kasih telah menggunakan aplikasi ini")
				// // 	return
				// default:
				// 	fmt.Println("Menu yang Anda pilih tidak tersedia, silahkan coba lagi")
			}
		}
	}

}
