package main

import (
	"fmt"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esutil"
)

type DiscordDoc struct {
	Url   string   `json:"url"`
	Time  int64    `json:"time"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}

func indexDoc(c *elasticsearch.Client, doc DiscordDoc) error {
	res, err := c.Index("discord", esutil.NewJSONReader(&doc), c.Index.WithPretty(), c.Index.WithDocumentID("jona-v1"))
	if err != nil {
		fmt.Println(err.Error())
	}

	// Debería ser 201
	fmt.Println(res.StatusCode)

	return nil
}

func createLocalClient() *elasticsearch.Client {
	fmt.Println("hola")
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses:  []string{"http://localhost:9200"},
		MaxRetries: 5,
	})

	if err != nil {
		panic(err)
	}

	res, err := client.Ping()
	if err != nil {
		panic(err)
	}

	println(res.StatusCode)
	println(res.String())

	return client
}

func main() {
	c := createLocalClient()

	err := indexDoc(c, DiscordDoc{
		Url:   "https://medium.com/@nate501/structuring-go-grpc-microservices-dd176fd28d0",
		Time:  time.Now().Unix(),
		Tags:  []string{"golang", "grpc", "microservices", "medium"},
		Title: "Structuring Go GRPC microservices",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("El documento se guardó correctamente =) ")
}
