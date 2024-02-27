package paser

func WebfingerUrl(username, host string) string {
	return `https://` + host + `/.well-known/webfinger?resource=acct:` + username + `@` + host
}
