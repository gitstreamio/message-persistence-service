package elastic

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
)
const msgIndex string = "messages"
const msgType string = "message"

func NewElasticClient(ctx context.Context) (*elastic.Client, error) {
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

	return client, err
}