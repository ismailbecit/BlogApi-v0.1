package config

var fields = map[string]string{
	"Email":    "E-posta",
	"Password": "Şifre",
}

func GetFields() *map[string]string {
	return &fields
}
