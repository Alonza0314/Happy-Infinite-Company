package models

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// cookie儲存username+loginTime的hash結果
// session儲存username以及loginTime
type UserInfo struct {
	Username  string    `json:"userName"`
	LoginTime time.Time `json:"loginTime"`
}

type ResetInfo struct {
	Username string `json:"userName"`
	Email    string `json:"email"`
}

func GenerateHash(userInfo UserInfo) string {
	hasher := sha256.New()
	hasher.Write([]byte(userInfo.Username + userInfo.LoginTime.GoString()))
	return hex.EncodeToString(hasher.Sum(nil))
}
