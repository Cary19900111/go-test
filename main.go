package main

import (
	"fmt"
	"reflect"
	"sort"
	"time"
)

type entOpenAcctResultReq struct {
	version       string
	pay_tenant_id uint64
	mch_no        string
	order_id      string
	send_date     string
}

// func (req entOpenAcctResultReq) GetSortString() string {
// 	key1 := reflect.TypeOf(req)
// 	value1 := reflect.ValueOf(req)

// 	var data = make(map[string]interface{})
// 	for i := 0; i < key1.NumField(); i++ {
// 		data[key1.Field(i).Name] = value1.Field(i).Interface()
// 	}
// 	// fmt.Printf("%+v", data)
// 	// var newMp = make([]string, 0)
// 	// for k, _ := range data {
// 	// 	newMp = append(newMp, k)
// 	// }
// 	// sort.Strings(newMp)
// 	// for _, v := range newMp {
// 	// 	fmt.Printf(v, data[v])
// 	// }
// 	// return data[v]
// 	// fmt.Printf("%+v", data)
// 	// for k, v := range rv {
// 	// 	fmt.Printf("%+v", k)
// 	// }
// 	// for k, v := range req {
// 	// 	fmt.Printf("%+v", k)
// 	// }
// 	return req.version
// }

type User struct {
	Id        int64
	Username  string
	Password  string
	Logintime time.Time
}

func Struct2Map(obj interface{}) string {
	var s string = ""
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
	return s
}

func main() {
	user := User{5, "zhangsan", "pwd", time.Now()}
	data := Struct2Map(user)
	fmt.Println(data)
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
