package user

import (
	"database/sql"
	"fmt"
)

func Register(db *sql.DB) (err error) {
	var IdUser int
	var NamaUser string
	var EmailUser string
	var NoHpUser string
	var PasswordUser string
	fmt.Println("Masukkan ID :")
	fmt.Scanln(&IdUser)
	fmt.Println("Masukkan Nama :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Masukkan Email:")
	fmt.Scanln(&EmailUser)
	fmt.Println("Masukkan No Hp :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password :")
	fmt.Scanln(&PasswordUser)

	_, err = db.
		Exec("insert into user(user_id,name,email,phone_number, password) values (?, ?, ?, ?, ?)",
			IdUser, NamaUser, EmailUser, NoHpUser, PasswordUser)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Yeay! Anda sudah berhasil mendaftar Dude! :")
	return err
}
