package black_kiwi_auth_structs

type User struct {
	Username string `json:"username"`
	Role int8 `json:"role"`
	Token string `json:"token"`
}

var MockUsers = []User{
	{Username: "admin", Role: 1, Token: "sdfsghjkuyrtesgdfhtrg"},
	{Username: "user", Role: 2, Token: "ushtejyhreghtjyrhgdrtsegter"},
}
