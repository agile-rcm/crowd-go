package crowd

// General structs

type crowdErrorMessage struct {
	Message string `json:"message"`
}

type Attributes struct {
	Attributes []*Attribute `json:"attributes,omitempty"`
}

type Attribute struct {
	Name   string   `json:"name"`
	Values []string `json:"values"`
}

type PasswordValue struct {
	Value string `json:"value,omitempty"`
}

// Group Structs

type GroupName struct {
	Name string `json:"name"`
}

type Group struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Type 		string	`json:"type"`
	Active		bool	`json:"active,omitempty"`
}

type GroupAttributes struct {
	Name		string			`json:"attributes"`
	Attributes	[]*Attribute	`json:"attribute"`
}

// User Structs

type User struct {
	Name        string			`json:"name"`
	FirstName   string			`json:"first-name"`
	LastName    string			`json:"last-name"`
	DisplayName string			`json:"display-name"`
	Email       string			`json:"email"`
	Key         string			`json:"key,omitempty"`
	IsActive    bool			`json:"active"`
	Password    PasswordValue	`json:"password,omitempty"`
	Attributes  Attributes		`json:"attributes,omitempty"`
}

type UserRename struct {
	NewName string `json:"new-name"`
}