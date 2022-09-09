package presenter

type StudentResponse struct {
	IdStudent   string
	FullName    string
	Email       string
	PhoneNumber string
	Token       string
}

type StudentRequest struct {
	FullName    string `json:"fullname" valid:"required~fullname is required,type(string)"`
	Email       string `json:"email" valid:"required~email is required,email,type(string)"`
	PhoneNumber string `json:"phoneNumber" valid:"required~PhoneNumber is required,,type(string)"`
	Password    string `json:"password" valid:"required,length(4|20)"`
}

type StudentUpdateRequest struct {
	FullName    string `json:"fullname" valid:"required~fullname is required,type(string)"`
	Email       string `json:"email" valid:"required~email is required,email,type(string)"`
	PhoneNumber string `json:"phoneNumber" valid:"required~PhoneNumber is required,,type(string)"`
}

type StudentLoginRequest struct {
	PhoneNumber string `json:"phoneNumber" valid:"required~PhoneNumber is required,,type(string)"`
	Password    string `json:"password" valid:"required"`
}

type StudentFindRequest struct {
	SearchText string `json:"searchText"`
	Limit      int    `json:"limit"`
	Offset     int    `json:"offset"`
}

type StudentFindResponse struct {
	Students *[]StudentResponse `json:"students"`
	Count    int                `json:"count"`
}
