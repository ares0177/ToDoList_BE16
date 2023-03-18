package user

import (
	"database/sql"
	"fmt"
)

func Update(db *sql.DB, noHp string) (err error) {
	var (
		EmailUser    string
		NamaUser     string
		PasswordUser string
		NoHpUser     string
	)

	fmt.Println("Nama yang ingin di ubah :")
	fmt.Scanln(&NamaUser)
	fmt.Println("Password yang ingin di ubah :")
	fmt.Scanln(&PasswordUser)
	fmt.Println("Email yang ingin di ubah :")
	fmt.Scanln(&EmailUser)

	_, err = db.
		Exec("update user set name=?, password=?, email=? where phone_number = ?",
			NamaUser, PasswordUser, EmailUser, NoHpUser)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Update berhasil !!")
	fmt.Printf("Nama : %s\nPassword : %s\nEmail : %s\nNomor Handphone : %s\n",
		NamaUser, PasswordUser, EmailUser, NoHpUser)

	return err
}
