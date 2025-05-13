package mySQL

import (
	"database/sql"
	"fmt"

	"github.com/fatemehmirarab/gameapp/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	user := entity.User{}
	row := d.db.QueryRow(`select * from users where phone_number = ? `, phoneNumber)
	err := ScanUser(row, &user)

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
	res, err := d.db.Exec(`INSERT INTO users (name, phone_number , password) VALUES (?,? ,?)`, user.Name, user.PhoneNumber, user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can not execute command %w", err)
	}
	//error is always nil
	id, _ := res.LastInsertId()
	user.Id = uint(id)
	return user, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {

	user := entity.User{}
	row := d.db.QueryRow(`select * from users where  phone_number = ? `, phoneNumber)
	err := ScanUser(row, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, false, nil
		} else {
			return entity.User{}, false, fmt.Errorf("can not execute command %w", err)
		}
	}
	return user, true, nil
}

func (d *MySQLDB) GetUserById(userId uint) (entity.User, error) {
	user := entity.User{}
	row := d.db.QueryRow(`select * from users where  id = ? `, userId)
	err := ScanUser(row, &user)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, nil
		} else {
			return entity.User{}, fmt.Errorf("can not execute command %w", err)
		}
	}
	return user, nil
}

func ScanUser(row *sql.Row, user *entity.User) error {

	return row.Scan(&user.Id, &user.PhoneNumber, &user.Name, &user.Password) // order based on sql columns
}
