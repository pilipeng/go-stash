package lib

import "testing"

func TestArrayOrSlice2String(t *testing.T) {
	data := []int64{21, 23, 24}
	string := ArrayOrSlice2String(data)
	if string != "21,23,24" {
		t.Fatalf("fail")
	}
}

func TestRemoveRepeatedElement(t *testing.T) {
	data := []string{
		"index_1",
		"index_2",
		"index_3",
		"index_2",
	}
	newSlice := RemoveRepeatedElement(data)
	if len(newSlice) != 3 {
		t.Fatalf("切掉去重失败！")
	}
}
