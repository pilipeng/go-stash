// package lib

// import (
// 	"encoding/json"
// 	"testing"
// 	"time"
// )

// type OrderDay struct {
// 	UpdateTime Date `json:"update_time"`
// }

// type OrderDatetime struct {
// 	UpdateDatetime Datetime `json:"update_datetime"`
// }

// func TestDate_MarshalJSON(t *testing.T) {
// 	date := time.Date(2021, 9, 26, 14, 9, 11, 11, time.Local)
// 	demo := &OrderDay{UpdateTime: Date{date}}
// 	v, _ := json.Marshal(demo)
// 	//fmt.Println(string(v)) // {"data":"2021-03-30"}
// 	if string(v) != `{"update_time":"2021-09-26"}` {
// 		t.Fatalf("Date_MarshalJSON fail")
// 	}

// }

// func TestDate_UnmarshalJSON(t *testing.T) {
// 	demo := &OrderDay{}
// 	err := json.Unmarshal([]byte(`{"update_time":"2021-03-30"}`), demo)
// 	if err != nil {
// 		t.Fatalf("Date json unmarsha error:%s", err)
// 	}
// 	//fmt.Printf("Date :2021-03-30 =>%s \n", demo.UpdateTime )
// 	if demo.UpdateTime.String() != "2021-03-30 00:00:00 +0800 CST" {
// 		t.Fatalf("Date_UnmarshalJSON \n")
// 	}
// }

// func TestDatetime_MarshalJSON(t *testing.T) {
// 	date := time.Date(2021, 9, 26, 14, 9, 11, 11, time.Local)
// 	demo := &OrderDatetime{UpdateDatetime: Datetime{date}}
// 	v, err := json.Marshal(demo)
// 	if err != nil {
// 		t.Fatalf("fail:%s\n", err)
// 	}
// 	//fmt.Println(string(v)) // {"update_datetime":"2021-09-26 14:09:11"}
// 	if string(v) != `{"update_datetime":"2021-09-26 14:09:11"}` {
// 		t.Fatalf("Date_MarshalJSON fail")
// 	}
// }

// func TestDatetime_UnmarshalJSON(t *testing.T) {
// 	demo := &OrderDatetime{}
// 	err := json.Unmarshal([]byte(`{"update_datetime":"2021-03-30 11:33:34"}`), demo)
// 	if err != nil {
// 		t.Fatalf("Date json unmarsha error:%s", err)
// 	}
// 	//fmt.Printf("Date :%s \n", demo.UpdateDatetime )
// 	if demo.UpdateDatetime.String() != "2021-03-30 11:33:34 +0000 UTC" {
// 		t.Fatalf("UnmarshalJSON error \n")
// 	}

// }
