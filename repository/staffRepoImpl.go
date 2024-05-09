package repository

import (
	"context"
	"database/sql"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"errors"
	"fmt"
	"time"
)

type StaffRepoImpl struct {
}

func NewStaffRepo() StaffRepo {
	return &StaffRepoImpl{}
}

func (repository *StaffRepoImpl) Save(ctx context.Context, tx *sql.Tx, staff domain.Staff) (domain.Staff, error) {
	var exists bool
	errValidation := tx.QueryRow("SELECT exists(SELECT 1 FROM staffs WHERE phone_number=$1)", staff.PhoneNumber).Scan(&exists)
	if errValidation != nil {
		fmt.Print("Phone Number not found")
	}
	if exists {
		fmt.Printf("Phone Number %s", staff.PhoneNumber)
		return staff, errors.New("Phone Number already exists")
	}
	sql := "INSERT INTO staffs(name, phone_number, password, created_at, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id"
	staff.CreatedAt = time.Now()
	staff.UpdatedAt = time.Now()
	insertedId := 0
	err := tx.QueryRowContext(ctx, sql, staff.Name, staff.PhoneNumber, staff.Password, staff.CreatedAt, staff.UpdatedAt).Scan(&insertedId)
	helper.PanicIfError(err)
	staff.Id = insertedId
	return staff, nil
}

func (repository *StaffRepoImpl) FindByPhoneNumber(ctx context.Context, tx *sql.Tx, email string) (domain.Staff, error) {
	var staff domain.Staff
	sql := "SELECT id, name, phone_number, password, created_at, updated_at FROM staffs WHERE phone_number=$1"
	err := tx.QueryRowContext(ctx, sql, email).Scan(&staff.Id, &staff.Name, &staff.PhoneNumber, &staff.Password, &staff.CreatedAt, &staff.UpdatedAt)
	if err != nil {
		return staff, errors.New("staff not found")
	}
	return staff, nil
}
