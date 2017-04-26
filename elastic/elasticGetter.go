package elastic

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"message-persistence-service/common"
	"encoding/json"
)

type elasticGetter struct {
	client *elastic.Client
	ctx    context.Context
}

/*
If no message is found nil, nil is returned as there was no error
 */
func (eg elasticGetter) GetById(id string) (common.Message, error) {
	res, err := eg.client.Get().Index(msgIndex).Type(msgType).Id(id).Do(eg.ctx)
	if err != nil {
		return nil, err
	}

	if res.Found {
		msg := &common.Message{}
		err = json.Unmarshal([]byte(res.Source),msg)
		if err != nil {
			return nil, err
		}
	}

	return nil, err

}

func (eg elasticGetter) Get(beginning int, amount int) ([]common.Message, error) {
	searchTerm := elastic.NewTermQuery(
	res, err := eg.client.Search()
}

func NewElasticGetter(ctx context.Context, client *elastic.Client) (common.Getter) {

	return &elasticGetter{client: client, ctx: ctx}
}