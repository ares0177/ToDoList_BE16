package user

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) (noHp string, err error) {
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Masukkan nomor Telepon anda Dude! :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password anda Dude! :")
	fmt.Scanln(&PasswordUser)

	var user User
	err = db.
		QueryRow("SELECT id, name, no_hp FROM user WHERE no_hp = ? and password = ?",
			NoHpUser, PasswordUser).
		Scan(&user.User_ID, &user.Nama, &user.Phone)
	if err != nil {
		fmt.Println("Waduh. Kata sandi anda salah nih Dude! coba periksa kembali ya Dude!")
		return "", err
	}

	fmt.Println("Yeay keren. Anda berhasil masuk Dude!")
	fmt.Printf("No Handphone : %s\nNama : %s\n", user.Phone, user.Nama)

	return user.Phone, nil
}
