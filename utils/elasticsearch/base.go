package main

import (
	"context"
	"github.com/olivere/elastic/v7"
)

var (
	subject   Subject
	indexName = "subject"
	servers   = []string{"http://172.12.15.134:9200/"}
)

type Subject struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Genres []string `json:"genres"`
}

func main() {
	const mapping = `
{
    "mappings": {
        "properties": {
            "id": {
                "type": "long"
            },
            "title": {
                "type": "text"
            },
            "genres": {
                "type": "keyword"
            }
        }
    }
}`

	ctx := context.Background()
	//client, err := elastic.NewClient(elastic.SetURL(servers...))
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(servers...))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(indexName).Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		_, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
	}
}
