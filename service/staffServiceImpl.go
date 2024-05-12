package service

import (
	"context"
	"database/sql"
	"eniqilo_store/config"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"eniqilo_store/model/web"
	"eniqilo_store/pkg/errorwrapper"
	"eniqilo_store/repository"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type StaffServiceImpl struct {
	StaffRepo repository.StaffRepo
	DB        *sql.DB
	Validate  *validator.Validate
}

func NewStaffService(staffRepository repository.StaffRepo, DB *sql.DB, validate *validator.Validate) StaffService {
	return &StaffServiceImpl{
		StaffRepo: staffRepository,
		DB:        DB,
		Validate:  validate,
	}
}

func (service *StaffServiceImpl) GenerateToken(ctx context.Context, tx *sql.Tx, staff domain.Staff) (string, error) {
	expTime := time.Now().Add(time.Hour * 8)
	claims := &config.JWTClaim{
		Name:    staff.Name,
		StaffId: staff.Id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}
	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key := []byte(os.Getenv("JWT_SECRET"))
	token, err := generateToken.SignedString(key)
	if err != nil {
		return "", errors.New("error generate token")
	}
	// err = service.SessionRepo.Save(ctx, tx, staff, token)
	// if err != nil {
	// 	return "", errors.New("error save session")
	// }

	return token, nil
}

func (service *StaffServiceImpl) Register(ctx context.Context, request web.StaffRegisterReq) (web.StaffRes, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.StaffRes{}, err
	}
	if !isPhoneNumberValid(request.PhoneNumber) {
		return web.StaffRes{}, errorwrapper.New(http.StatusBadRequest, errors.New("invalid phone number"), "")
	}
	tx, err := service.DB.Begin() // transaction db
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusInternalServerError, err, "error database transaction")
	}
	defer helper.CommitOrRollback(tx)
	// hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error generate password"), "")
	}
	request.Password = string(bytes)

	staff := domain.Staff{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		Name:        request.Name,
	}
	staff, err = service.StaffRepo.Save(ctx, tx, staff)
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error save data to database"), "")
	}
	fmt.Println("staff:", staff.Id)
	token, err := service.GenerateToken(ctx, tx, staff)
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusInternalServerError, errors.New("error generate token"), "")
	}
	return helper.ToCategoryResponseStaff(staff, token), nil
}

func (service *StaffServiceImpl) Login(ctx context.Context, request web.StaffLoginReq) (web.StaffRes, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.StaffRes{}, err
	}
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	staff, err := service.StaffRepo.FindByPhoneNumber(ctx, tx, request.PhoneNumber)
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusNotFound, err, "")
	}
	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(request.Password))
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusBadRequest, errors.New("invalid password"), "")
	}
	token, err := service.GenerateToken(ctx, tx, staff)
	if err != nil {
		return web.StaffRes{}, errorwrapper.New(http.StatusInternalServerError, err, "")
	}
	return helper.ToCategoryResponseStaff(staff, token), nil
}

func isPhoneNumberValid(phoneNumber string) bool {
	// This is a simple regex for validating an international phone number, which allows for country codes starting with '+'
	// followed by up to 15 digits. This may not cover all possible international phone number formats.
	// You may need to adjust this regex to suit your specific needs.
	regex := `^\+[1-9]{1}[0-9]{9,15}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match
}
