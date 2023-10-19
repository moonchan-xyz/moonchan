package utils

func ParseWebfingerUrl(username, host string) string {
	return `https://` + host + `/.well-known/webfinger?resource=acct:` + username + `@` + host
}
