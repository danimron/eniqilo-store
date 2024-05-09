package service

import (
	"context"
	"database/sql"
	"eniqilo_store/config"
	"eniqilo_store/helper"
	"eniqilo_store/model/domain"
	"eniqilo_store/model/web"
	"eniqilo_store/repository"
	"os"
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

func (service *StaffServiceImpl) GenerateToken(ctx context.Context, staff domain.Staff) string {
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
	helper.PanicIfError(err)
	return token
	// set cookie
}

func (service *StaffServiceImpl) Register(ctx context.Context, request web.StaffRegisterReq) (web.StaffRes, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.StaffRes{}, err
	}
	helper.PanicIfError(err)
	tx, err := service.DB.Begin() // transaction db
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	// hash password
	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	helper.PanicIfError(err)
	request.Password = string(bytes)

	staff := domain.Staff{
		PhoneNumber: request.PhoneNumber,
		Password:    request.Password,
		Name:        request.Name,
	}
	staff, err = service.StaffRepo.Save(ctx, tx, staff)
	token := service.GenerateToken(ctx, staff)
	if err != nil {
		return web.StaffRes{}, err
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
		return web.StaffRes{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(staff.Password), []byte(request.Password))
	if err != nil {
		return web.StaffRes{}, err
	}
	token := service.GenerateToken(ctx, staff)
	return helper.ToCategoryResponseStaff(staff, token), nil
}
