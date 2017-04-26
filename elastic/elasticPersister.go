package elastic

import (
	"context"
	"gopkg.in/olivere/elastic.v5"
	"message-persistence-service/common"
)

const msgIndex string = "messages"
const msgType string = "message"

type elasticAdapter struct {
	client *elastic.Client
	ctx    context.Context
}

func NewElasticAdapter(ctx context.Context) (common.PersistenceAdapter, error) {
	// Create a client
	client, err := elastic.NewClient(elastic.SetURL("http://elastic:9200"))
	if err != nil {
		return nil, err
	}

	indexExists, err := client.IndexExists().Index([]string{msgIndex}).Do(ctx)
	if err != nil {
		return nil, err
	}
	if !indexExists {
		_, err = client.CreateIndex(msgIndex).Do(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &elasticAdapter{client: client, ctx: ctx}, err

}

func (ea *elasticAdapter) Create(msg common.Message) (id string, err error) {
	resp, err := ea.client.Index().
		Index(msgIndex).
		Type(msgType).
		BodyJson(msg).
		Do(ea.ctx)

	return resp.Id, err
}

func (ea *elasticAdapter) Update(id string, msg common.Message) error {
	_, err := ea.client.Index().
		Index(msgIndex).
		Type(msgType).
		Id(id).
		BodyJson(msg).
		Do(ea.ctx)

	return err
}

func (ea *elasticAdapter) Delete(id string) error {
	_, err := ea.client.Delete().
		Index(msgIndex).
		Type(msgType).
		Id(id).
		Do(ea.ctx)

	return err
}
