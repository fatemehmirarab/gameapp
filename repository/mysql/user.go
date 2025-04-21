package mySQL

import (
	"database/sql"
	"fmt"

	"github.com/fatemehmirarab/gameapp/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	res := d.db.QueryRow(`select * from users where phone_number = ? `, phoneNumber)
	err := res.Scan(&user.Id, &user.PhoneNumber, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		} else {
			return false, fmt.Errorf("unexpected error %v", err)
		}
	}
	return false, nil
}

func (d *MySQLDB) Register(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`INSERT INTO users (name, phone_number) VALUES (?,?)`, user.Name, user.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not execute command %w", err)
	}
	//error is always nil
	id, _ := res.LastInsertId()
	user.Id = uint(id)
	return user, nil
}
