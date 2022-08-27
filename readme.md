:point_right: Con Docker-Compose (ElasticSearch & Kibana)

```
docker-compose up
```

:point_right: Docker (only Elastic)

```
docker pull docker.elastic.co/elasticsearch/elasticsearch:7.17.6
```

```
docker run -p 127.0.0.1:9200:9200-p 127.0.0.1:9300:9300-e "discovery.type=single-node" docker.elastic.co/elasticsearch/elasticsearch:7.17.6
```

:point*right: \_Estructura del json*

```
DiscordDoc struct {
	Url   string   `json:"url"`
	Time  int64    `json:"time"`
	Tags  []string `json:"tags"`
	Title string   `json:"title"`
}
```

:point_right: Go

```
go run main.go
```

Once run go run main.go on browser http://localhost:9200/discord/\_doc/jona-v1?pretty or Postman
