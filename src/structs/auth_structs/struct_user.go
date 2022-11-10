package black_kiwi_auth_structs

type User struct {
	Username string `json:"username"`
	Role int8 `json:"role"`
}

var MockUsers = []User{
	{Username: "admin", Role: 2},
	{Username: "user", Role: 1},
}
