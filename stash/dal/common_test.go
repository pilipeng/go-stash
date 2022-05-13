package dal

import (
	"testing"
)

func TestTable2index(t *testing.T) {
	tableName := "pn_order"
	createTime := 1632272916
	indexName := Table2index("100006_", tableName, createTime)
	if indexName != "100006_fx_order_202109" {
		t.Fatalf("indexName:%s", indexName)
	}
}
