# Happy Infinite Company

---

## clientAccount.go

底層處理客戶帳號資料與資料庫交互的主要function
會調用cookieSession.go以及mongodb.go的函數

---

## cookiesession.go

用於記錄網站登入資訊（暫定5分鐘保持登入）
以及更改密碼時的臨時session - "resetid"（暫定5分鐘）

### Cookie

+ userid: username + login time的sha256密文

### Session

+ {key: userid, value: UserInfo}

```golang
    type UserInfo struct {
        UserName  string    `json:"userName"`
        LoginTime time.Time `json:"loginTime"`
    }
```

+{key: "resetid", value: ResetInfo}

```golang
      type ResetInfo struct {
      Username string `json:"userName"`
      Email    string `json:"email"`
    }
```

---

## mongodb.go

Database：HIC
Collection：{
    clientAccount：記錄客戶帳號資料
}
