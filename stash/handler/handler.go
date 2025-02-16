package handler

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/kevwan/go-stash/stash/config"
	"github.com/kevwan/go-stash/stash/es"
	"github.com/kevwan/go-stash/stash/etl"
	"github.com/kevwan/go-stash/stash/filter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type MessageHandler struct {
	writer  *es.Writer
	indexer *es.Index
	filters []filter.FilterFunc
	processor *config.Cluster
	dbConn *sqlx.SqlConn
}

func NewHandler(processor *config.Cluster,writer *es.Writer, indexer *es.Index, dbConn *sqlx.SqlConn) *MessageHandler {
	return &MessageHandler{
		writer:  writer,
		indexer: indexer,
		processor: processor,
		dbConn: dbConn,
	}
}

func (mh *MessageHandler) AddFilters(filters ...filter.FilterFunc) {
	for _, f := range filters {
		mh.filters = append(mh.filters, f)
	}
}

func (mh *MessageHandler) Consume(_, val string) error {
	var m map[string]interface{}
	if err := jsoniter.Unmarshal([]byte(val), &m); err != nil {
		return err
	}
	//index := mh.indexer.GetIndex(m)
	for _, proc := range mh.filters {
		if m = proc(m); m == nil {
			return nil
		}
	}
	/*bs, err := jsoniter.Marshal(m)
	if err != nil {
		return err
	}
	return mh.writer.Write(index, string(bs))*/


	//自定义数据解析
	ctx:= context.Background()
	lists := etl.Process(ctx,mh.processor,mh.dbConn,[]byte(val))
	for indexKey, row := range lists {
		bs, err := jsoniter.Marshal(row.Item)
		if err != nil {
			return err
		}
		if len(row.EsIndexName) == 0 {
			logx.Errorf("EsIndexName is empty:%s", row.EsIndexName)
			continue
		}
		index := row.EsIndexName
		logx.Infof("Es index:%fs;_id:%s",index,indexKey)
		err = mh.writer.Write(index,indexKey, string(bs))
		if err != nil {
			return err
		}
	}
	return nil



}
