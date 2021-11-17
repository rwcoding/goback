package config

type Config struct {
	IsDev  bool
	Log    string
	Lang   string
	Header bool
	OnlyGP bool
}

var CC = &Config{
	IsDev: false,
	Lang:  "zh",
}

func IsDev() bool {
	return CC.IsDev
}

func Log() string {
	return CC.Log
}

func GetLang() string {
	return CC.Lang
}

func NeedWriteHeader() bool {
	return CC.Header
}

func OnlyGetPost() bool {
	return CC.OnlyGP
}
