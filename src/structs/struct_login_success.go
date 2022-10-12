package black_kiwi_structs

type LoginSuccess struct {

	Username string `json:"username"`

	Role string `json:"role,omitempty"`
}
