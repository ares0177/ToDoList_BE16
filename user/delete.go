package user

import (
	"database/sql"
	"fmt"
)

func Delete(db *sql.DB, NoHp string) (err error) {

	var konfirmasi string
	fmt.Println("Anda serius mau hapus akun ? (Y/n)")
	fmt.Scanln(&konfirmasi)

	if konfirmasi != "Y" {
		fmt.Println("Tidak jadi menghapus akun")
		return err
	}

	_, err = db.
		Exec("delete from user where phone_number = ?", NoHp)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("delete success!")

	return err
}
