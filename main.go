package main

import (
	"Todo/config"
	"Todo/user"
	"fmt"
)

func main() {
	var (
		err  error
		menu int
		NoHp string
	)
	config.Connect()
	db := config.GetConnection()

	for {
		if menu != 2 {
			fmt.Println("====Selamat datang====")
			fmt.Println("Silahkan pilih menu")
			fmt.Println("1. Register")
			fmt.Println("2. Login")
			fmt.Println("0. Keluar")

			fmt.Println("Masukkan pilihan :")
			fmt.Scanln(&menu)
		}

		switch menu {
		case 1:
			fmt.Println("Silahkan Register Dude!")
			if err = user.Register(db); err != nil {
				return
			}
		case 2:
			if NoHp == "" {
				fmt.Println("Silahkan Login!")
				if NoHp, err = user.Login(db); err != nil {
					menu = 0
					return
				}

			}

		default:
			fmt.Println("Yang anda inputkan salah! mohon coba lagi")
			return
		}

		var pilihan int
		fmt.Println("1. Updated Account")
		fmt.Println("2. Delete Account")
		fmt.Println("3. Kegiatan anda")

		fmt.Println("Masukkan pilihan anda :")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			if err = user.Update(db, NoHp); err != nil {
				return
			}
		}
	}

}
