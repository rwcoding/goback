package util

func Password(pwd string, salt string, isMd5 bool) string {
	if isMd5 {
		return Md5(pwd + salt)
	}
	return Md5(Md5(pwd) + salt)
}
