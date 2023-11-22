package clientmodels

// Generated by https://quicktype.io

type CreateUserRequest struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	ID       string      `json:"id"`
	Username string      `json:"username"`
	Name     string      `json:"name"`
	Email    string      `json:"email"`
	Roles    []ClaimRole `json:"roles"`
	Claims   []ClaimRole `json:"claims"`
}

type ClaimRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserClaimRoleCreate struct {
	Name string `json:"name"`
}
