package user

import (
	"database/sql"
	"fmt"
)

func Register(db *sql.DB) (err error) {
	var (
		NamaUser     string
		EmailUser    string
		NoHpUser     string
		PasswordUser string
	)

	fmt.Println("Masukkan Nama :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&EmailUser)
	fmt.Println("Masukkan No Hp :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&PasswordUser)

	_, err = db.
		Exec("insert into user(name,email,phone_number, password) values (?, ?, ?, ?)",
			NamaUser, EmailUser, NoHpUser, PasswordUser)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Yeay! Anda sudah berhasil mendaftar Dude! :")
	return err
}
