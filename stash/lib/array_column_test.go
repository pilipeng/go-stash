package lib

import (
	"testing"
)

type User struct {
	ID int
	NAME string
}

func TestBindFromArrayColumn(t *testing.T) {
	user1 := User{
		ID:   1,
		NAME: "zwk",
	}
	user2 := User{
		ID:   2,
		NAME: "zzz",
	}
	var list3 []User
	list3 = append(list3, user1)
	list3 = append(list3, user2)

	var userMap map[int]string
	//map[int]string{1:"zwk", 2:"zzz"}
	StructColumn(&userMap, list3, "NAME", "ID")
	//fmt.Printf("%#v\n", userMap)

	if userMap[1] != "zwk" || userMap[2] != "zzz"{
		t.Fatalf("fail\n")
	}

	var userMap1 map[int]User
	//map[int]lib.User{1:lib.User{ID:1, NAME:"zwk"}, 2:lib.User{ID:2, NAME:"zzz"}}
	StructColumn(&userMap1, list3, "", "ID")
	//fmt.Printf("%#v\n", userMap1)


	var userSlice []int
	//[]int{1, 2}
	StructColumn(&userSlice, list3, "ID", "")
	//fmt.Printf("%#v\n", userSlice)
}
