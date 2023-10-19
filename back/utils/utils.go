package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"net/url"
	"strings"
	"sync"
	"time"
)

// var MetaData map[string]string = make(map[string]string)

// const SALT = "[salt]"
func Hash(s string) string {
	// Compute the SHA-256 hash of the message
	hash := sha256.Sum256([]byte(s + SALT))

	// Convert the hash to a hexadecimal string
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

var lastTimestamp int64 = 0
var muxTimestamp sync.Mutex

// ms = time.Now().UnixMilli()
func Timestamp() (ms int64) {
	muxTimestamp.Lock()
	defer muxTimestamp.Unlock()
	ms = time.Now().UnixMilli()
	if ms == lastTimestamp {
		ms++
	}
	lastTimestamp = ms
	return
}

// acct = username + "@" + host
func ParseAcct(username, host string) (acct string) {
	acct = username + "@" + host
	return
}

// username, host = "[username]@[host]"
func ParseUserAndHost(acct string) (username, host string, err error) {
	acct, err = url.QueryUnescape(acct)
	if err != nil {
		return
	}

	acct = strings.TrimPrefix(acct, "@")
	arr := strings.Split(acct, "@")
	if len(arr) != 2 {
		err = errors.New("format")
		return
	}
	username, host = arr[0], arr[1]
	return
}
