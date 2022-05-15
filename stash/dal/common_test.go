package dal

import (
	"testing"
)

func TestTable2index(t *testing.T) {
	tableName := "pn_order"
	var createTime int64
	createTime = 1632272916
	indexName := Table2index("go_by_", tableName, createTime)
	if indexName != "go_by_order_202109" {
		t.Fatalf("indexName:%s", indexName)
	}
}
