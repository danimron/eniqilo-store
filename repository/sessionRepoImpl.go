package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/model/domain"
	"errors"
	"fmt"
	"time"
)

type SessionRepoImpl struct {
}

func NewSessionRepo() SessionRepo {
	return &SessionRepoImpl{}
}

func (repository *SessionRepoImpl) Save(ctx context.Context, tx *sql.Tx, staff domain.Staff, token string) error {
	var exists bool
	var sql string
	var session domain.Session
	fmt.Println("staffidatas:", staff.Id)
	errValidation := tx.QueryRow("SELECT exists(SELECT 1 FROM sessions WHERE user_id=$1)", staff.Id).Scan(&exists)
	if errValidation != nil {
		fmt.Println("Staff not found")
	}
	session.UpdatedAt = time.Now()
	insertedId := 0
	if exists {
		// make query to update the existing session with new token
		sql = "UPDATE sessions SET token=$1, updated_at=$2 WHERE staff_id=$3 RETURNING id"
		fmt.Println("sql:", sql)
		err := tx.QueryRowContext(ctx, sql, token, session.UpdatedAt, staff.Id).Scan(&insertedId)
		if err != nil {
			return errors.New("error update session")
		}
	} else {
		// make query to insert new session
		sql = "INSERT INTO sessions(staff_id, created_at, updated_at, token) VALUES($1, $2, $3, $4) RETURNING id"
		fmt.Println("sql:", sql)
		session.CreatedAt = time.Now()
		fmt.Println("id:", session.Id)
		fmt.Println("staffid:", session.StaffId)
		err := tx.QueryRowContext(ctx, sql, staff.Id, session.CreatedAt, session.UpdatedAt, token).Scan(&insertedId)
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(err)
			fmt.Println("id:", session.Id)
			fmt.Println("staffid:", session.StaffId)
			return errors.New("error insert session")
		}
	}
	session.Id = insertedId
	return nil
}

func (repository *SessionRepoImpl) FindByStaffId(ctx context.Context, tx *sql.Tx, staffId int) error {
	var session domain.Session
	sql := "SELECT user_id, token FROM sessions WHERE user_id=$1"
	err := tx.QueryRowContext(ctx, sql, staffId).Scan(&session.Id, &session.StaffId, &session.Token)
	if err != nil {
		return errors.New("staff not found")
	}
	return nil

}
