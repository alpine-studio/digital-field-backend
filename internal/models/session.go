package models

var PossibleRoles = map[string]struct{}{
	"агроном":     {},
	"виноградарь": {},
}

type Session struct {
	ID    int    `json:"-"`
	Token string `json:"token"`
	Role  string `json:"role"`
}
