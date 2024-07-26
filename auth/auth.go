package auth

import (
	"log"
	"golang.org/x/crypto/bcrypt"
	"time"
        b64 "encoding/base64"
)

func Hash(s string) string {
        hash, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
        if err != nil {
                log.Println(err)
        }
        return string(hash)
}

func Refresh(s string) string {
        now := time.Now()
        sec := now.Unix()
	enc := b64.StdEncoding.EncodeToString([]byte(s+string(sec)))
	return enc
}