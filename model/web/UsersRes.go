package web

type UsersRes struct {
	UserId    		string `json:"id"`
	PhoneNumber    	string `json:"phoneNumber"`
	Name     		string `json:"name"`
	Token 			string `json:"accessToken"`
}
