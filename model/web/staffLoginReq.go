package web

type StaffLoginReq struct {
	PhoneNumber string `validate:"required" json:"phoneNumber"`
	Password    string `validate:"required" json:"password"`
}
