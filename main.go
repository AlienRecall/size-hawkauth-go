package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"time"
)

func gen(id string, method string, path string, domain string) string {
	ts := time.Now().Unix()
	randomBytes := make([]byte, 16)
	rand.Read(randomBytes)
	nonce := b64.StdEncoding.EncodeToString([]byte(randomBytes))
	data := fmt.Sprintf(`"hawk.1.header\n%d\n%s\n%s\n%s\n%s\n80"`, ts, nonce, method, path, domain)
	h := hmac.New(sha256.New, []byte("7fbda8fefe33eec041a3cbca366fab32"))
	h.Write([]byte(data))
	mac := b64.StdEncoding.EncodeToString([]byte(h.Sum(nil)))
	header := fmt.Sprintf(`Hawk id="%s", mac="%s", ts="%d", nonce="%s"`, id, mac, ts, nonce)
	return header
}

func main() {
  sizeID := "c012543e9e"
  method := "GET"
  path := "/stores/size/carts?channel=android-app-phone"
  domain := "prod.jdgroupmesh.cloud"
  header := gen(sizeID, method, path, domain)
	fmt.Println(header)
}
