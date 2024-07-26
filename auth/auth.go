package auth

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
        b64 "encoding/base64"
)

func Hash(s string) string {
        hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
        if err != nil {
                fmt.Println(err)
        }
        return string(hash)
}

func Encode(s string) string {
	enc := b64.URLEncoding.EncodeToString([]byte(s))
	return enc
}

func Decode(s string) (string, string) {
	dec, _ := b64.URLEncoding.DecodeString(s)
	n := strings.Split(string(dec), " ")
	if len(n) > 1 {
		return fmt.Sprint(n[0]), n[1]
	}
	return "",""
}