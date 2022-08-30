package presenter

type StudentResponse struct {
	IdStudent   string
	Email       string
	PhoneNumber string
	Password    string
}

type StudentRequest struct {
	FullName    string `json:"fullname" valid:"required~fullname is required,type(string)"`
	Email       string `json:"email" valid:"required~email is required,email,type(string)"`
	PhoneNumber string `json:"phoneNumber" valid:"required~PhoneNumber is required,,type(string)"`
	Password    string `json:"password" valid:"required,length(4|20)"`
}
