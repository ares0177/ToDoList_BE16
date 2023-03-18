package user

import (
	"database/sql"
	"fmt"
)

func Login(db *sql.DB) (noHp string, err error) {
	var PasswordUser string
	var NoHpUser string
	fmt.Println("Masukkan nomor Telepon anda :")
	fmt.Scanln(&NoHpUser)
	fmt.Println("Masukkan Password anda :")
	fmt.Scanln(&PasswordUser)

	var user User
	err = db.
		QueryRow("SELECT user_id, name, phone_number FROM user WHERE phone_number = ? and password = ?",
			NoHpUser, PasswordUser).
		Scan(&user.User_ID, &user.Nama, &user.Phone)
	if err != nil {
		fmt.Println("Yang anda masukkan salah. Mohon coba lagi")
		return "", err
	}

	fmt.Println("Yeay keren. Anda berhasil masuk!")
	fmt.Printf("No Handphone : %s\nNama : %s\n", user.Phone, user.Nama)

	return user.Phone, nil
}
