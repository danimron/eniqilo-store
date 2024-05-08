package web

type StaffsRegisterReq struct {
	PhoneNumber    	string `validate:"required" json:"phoneNumber"`
	Password 		string `validate:"required" json:"password"`
	Name     		string `validate:"required" json:"name"`
}
