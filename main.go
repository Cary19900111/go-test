package main

import (
	"fmt"
	"reflect"
)

type entOpenAcctResultReq struct {
	version       string
	pay_tenant_id uint64
	mch_no        string
	order_id      string
	send_date     string
}

func (req entOpenAcctResultReq) GetSortString() string {
	rv := reflect.ValueOf(req)
	for k, v := range rv {
		fmt.Printf("%+v", k)
	}
	// for k, v := range req {
	// 	fmt.Printf("%+v", k)
	// }
	return "#########"
}

func main() {
	req := entOpenAcctResultReq{
		version:       "v0.0.1",
		pay_tenant_id: 123,
		mch_no:        "mch_no12332123",
		order_id:      "order_id13493042",
		send_date:     "201902010000",
	}
	var ver = req.GetSortString()
	fmt.Printf("%+v", ver)
	fmt.Printf("\n")

	// resp, err := http.Get("http://www.baidu.com")
	// if err != nil {
	// 	// handle error
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("%+v", resp)
}
