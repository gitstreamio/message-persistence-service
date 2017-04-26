package elastic

import (
	"gopkg.in/olivere/elastic.v5"
	"context"
	"message-persistence-service/common"
	"encoding/json"
	"strings"
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

func (eg elasticGetter) Get(timeline string, from int, size int) ([]common.Message, error) {
	//we need a term query for the user/orga AND optional project so that we can do a boolquery on elastic
	terms := strings.Split(timeline, "/")
	termQueries := make([]*elastic.TermQuery, 1)
	for term := range terms {
		termQueries = append(termQueries, elastic.NewTermQuery("timeline", term))
	}
	searchQuery := elastic.NewBoolQuery().Must(termQueries...)

	results, err := doSearch(eg, searchQuery, from, size)
	if err != nil {
		return nil, err
	}

	messages := make([]common.Message, results.Hits.TotalHits)
	for i, hit := range results.Hits.Hits {
		msg := &common.Message{}
		err := json.Unmarshal([]byte(hit.Source), msg)
		if err != nil {
			return nil, err
		}
		messages[i] = *msg
	}
	return messages, err

}
func doSearch(eg elasticGetter, searchQuery *elastic.BoolQuery, from int, size int) (*elastic.SearchResult, error) {
	return eg.client.Search(msgIndex).Type(msgType).Query(searchQuery).From(from).Size(size).Do(eg.ctx)
}

func NewElasticGetter(ctx context.Context, client *elastic.Client) (common.Getter) {

	return &elasticGetter{client: client, ctx: ctx}
}