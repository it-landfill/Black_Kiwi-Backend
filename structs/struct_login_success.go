package black_kiwi

type LoginSuccess struct {

	Username string `json:"username"`

	Role string `json:"role,omitempty"`
}
