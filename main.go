package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"sort"
	"strings"
)

type entOpenAcctResultReq struct {
	Version     string
	PayTenantId uint64
	MchNo       string
	OrderId     string
	SendDate    string
}

type User struct {
	Id       int64
	Username string
	Password string
}

func GetUsernameAndPwd(obj interface{}) (string, string) {
	var s string
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		KeyValue := fmt.Sprintf("%v=%v", k, data[k])
		s += "&" + KeyValue
		// fmt.Println("Key:", k, "Value:", data[k])
	}
	return "", ""
	// timeNow := time.Now().Format("20060102150405")
	// keyID := "KY0123456789012345678900"
	// s = keyID + "&" + timeNow + "&" + "POST" + "&" + "/v1/open" + s
	// privateKey, _ := GetPrivateKey()
	// p, privateKey := pem.Decode(privateKey)
	// // fmt.Println(privateKey)
	// //var priKey *rsa.PrivateKey
	// priKey, err := x509.ParsePKCS1PrivateKey(p.Bytes)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", ""
	// }
	// pwd, err := RSAWithSHA1(s, priKey)
	// if err != nil {
	// 	fmt.Printf("%+v", err)
	// 	return "", ""
	// }
	// username := keyID + "_" + timeNow
	// return username, pwd
	// return "test"
}

func GetPrivateKey() ([]byte, error) {
	filePth := "./rsa-private_key.pem"
	f, err := os.Open(filePth)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func RSAWithSHA1(s string, privateKey *rsa.PrivateKey) (string, error) {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write([]byte(s))
	hashed := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey,
		crypto.SHA1, hashed)
	if err != nil {
		fmt.Println("Error from signing: %s\n", err)
		return "", err
	}
	// fmt.Printf("Signature: %x\n", signature)
	signRet := fmt.Sprintf("%x", signature)
	// fmt.Printf("sigRet: %s\n", signRet)
	return signRet, nil
}

func httpGet(username string, pwd string) {
	authString := fmt.Sprintf("%s:%s", username, pwd)
	encodeString := base64.StdEncoding.EncodeToString([]byte(authString))
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://open.test.cibfintech.com/api/cloudwallet/mchAccsQuery", strings.NewReader("name=cjb"))
	if err != nil {
		fmt.Printf("%+v", err)
	}
	req.Header.Set("Authorization", encodeString)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%+v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%+v", err)
	}

	fmt.Println(string(body))
}

func main() {
	// user := User{5, "zhangsan", "pwd"}
	req := entOpenAcctResultReq{
		Version:     "v0.0.1",
		PayTenantId: 123,
		MchNo:       "mch_no12332123",
		OrderId:     "order_id13493042",
		SendDate:    "201902010000",
	}
	// fmt.Println(user)
	// fmt.Println(req)
	username, pwd := GetUsernameAndPwd(req)
	httpGet(username, pwd)
	fmt.Println("test")
	// username,password:=
	// req := entOpenAcctResultReq{
	// 	version:       "v0.0.1",
	// 	pay_tenant_id: 123,
	// 	mch_no:        "mch_no12332123",
	// 	order_id:      "order_id13493042",
	// 	send_date:     "201902010000",
	// }
	// var ver = req.GetSortString()
	// fmt.Printf("%+v", ver)
	// fmt.Printf("\n")

	// resp, err := http.Get("http://www.baidu.com")
	// if err != nil {
	// 	// handle error
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v", resp)
}
