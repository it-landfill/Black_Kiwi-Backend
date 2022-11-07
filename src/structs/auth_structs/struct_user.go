package black_kiwi_auth_structs

type User struct {
	Username string `json:"username"`
	Role string `json:"role"`
	Token string `json:"token"`
}

var MockUsers = []User{
	{Username: "admin", Role: "admin", Token: "sdfsghjkuyrtesgdfhtrg"},
	{Username: "user", Role: "user", Token: "ushtejyhreghtjyrhgdrtsegter"},
}
