# Happy Infinite Company

---

| 內容 | 實現 | 版本 |
|-|-|-|
| OS | Linux | Ubuntu 22.04.4 LTS  |
| Frontend | HTML + CSS + JS |-|
| Backend | Golang(Gin) | go1.22.0 linux/amd64 |
| Database | MongoDB | mongod - v7.0.6 |

---

## 使用方法（測試版： non-build）

1. clone Repo
2. 進入 configs/config.conf 修改 [httpserver] addr="服務器主機ip"
3. $ go mod tidy 確保項目依賴正確
4. $ go run main.go
