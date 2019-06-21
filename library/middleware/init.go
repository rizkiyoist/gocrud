package middleware

var myrole map[string][]string

func InitRole(roles map[string][]string) {
	myrole = roles
}
