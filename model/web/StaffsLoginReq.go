package web

type StaffsLoginReq struct {
	PhoneNumber    	string `validate:"required" json:"phoneNumber"`
	Password 		string `validate:"required" json:"password"`
}
