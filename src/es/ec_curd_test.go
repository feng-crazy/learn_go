package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var host = []string{
	"http://127.0.0.1:9200/",
	//"http://10.42.0.122:9200/",
	//"http://10.42.0.123:9200/",
}

var client *elastic.Client

//初始化
func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetURL(host...))
	if err != nil {
		fmt.Printf("create client failed, err: %v", err)
	}
}

//ping 连接测试
func PingNode() {
	start := time.Now()

	info, code, err := client.Ping(host[0]).Do(context.Background())
	if err != nil {
		fmt.Printf("ping es failed, err: %v", err)
	}

	duration := time.Since(start)
	fmt.Printf("cost time: %v\n", duration)
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

//校验 index 是否存在
func IndexExists(index ...string) bool {
	exists, err := client.IndexExists(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	return exists
}

//创建 index
func CreateIndex(index, mapping string) bool {
	result, err := client.CreateIndex(index).BodyString(mapping).Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err: %v\n", err)
	}
	return result.Acknowledged
}

//删除 index
func DelIndex(index... string) bool {
	response, err := client.DeleteIndex(index...).Do(context.Background())
	if err != nil {
		fmt.Printf("delete index failed, err: %v\n", err)
	}
	return response.Acknowledged
}

//批量插入
func Batch(index string, type_ string, datas... interface{})  {

	bulkRequest := client.Bulk()
	for i, data := range datas {
		doc := elastic.NewBulkIndexRequest().Index(index).Type(type_).Id(strconv.Itoa(i)).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		panic(err)
	}
	failed := response.Failed()
	iter := len(failed)
	fmt.Printf("error: %v, %v\n", response.Errors,  iter)
}

//获取指定 Id 的文档
func GetDoc(index, id string) []byte {
	temp := client.Get().Index(index).Id(id)
	get, err := temp.Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get.Id, get.Version, get.Index, get.Type)
	}
	source, err := get.Source.MarshalJSON()
	if err != nil {
		fmt.Printf("byte convert string failed, err: %v", err)
	}
	return source
}

//term 查询
func TermQuery(index, type_, fieldName, fieldValue string) *elastic.SearchResult {
	query := elastic.NewTermQuery(fieldName, fieldValue)
	//_ = elastic.NewQueryStringQuery(fieldValue) //关键字查询

	searchResult, err := client.Search().
		Index(index).Type(type_).
		Query(query).
		From(0).Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}
	fmt.Printf("query cost %d millisecond.\n", searchResult.TookInMillis)

	return searchResult
}

func Search(index, type_ string) *elastic.SearchResult {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Must(elastic.NewMatchQuery("user", "Jame10"))
	boolQuery.Filter(elastic.NewRangeQuery("age").Gt("30"))
	searchResult, err := client.Search(index).
		Type(type_).Query(boolQuery).Pretty(true).Do(context.Background())
	if err != nil {
		panic(err)
	}

	return searchResult
}

func AggsSearch(index, type_ string) {

	minAgg := elastic.NewMinAggregation().Field("age")
	rangeAgg := elastic.NewRangeAggregation().Field("age").AddRange(0,30).AddRange(30,60).Gt(60)


	build := client.Search(index).Type(type_).Pretty(true)

	minResult, err := build.Aggregation("minAgg", minAgg).Do(context.Background())
	rangeResult, err := build.Aggregation("rangeAgg", rangeAgg).Do(context.Background())
	if err != nil {
		panic(err)
	}

	minAggRes, _ := minResult.Aggregations.Min("minAgg")
	fmt.Printf("min: %v\n", *minAggRes.Value)

	rangeAggRes, _ := rangeResult.Aggregations.Range("rangeAgg")
	for _, item := range rangeAggRes.Buckets {
		fmt.Printf("key: %s, value: %v\n", item.Key, item.DocCount)
	}

}


type Tweet struct {
	User     string                `json:"user"`
	Age      int                   `json:"age"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"image,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

var mapping = `{
	"settings":{
		"number_of_shards": 3,
		"number_of_replicas": 1
	},
	"mappings":{
		"doc":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"age":{
					"type": "integer"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"image":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func TestPingNode(t *testing.T) {
	PingNode()
}

func TestIndexExists(t *testing.T) {
	result := IndexExists("car_source", "test")
	fmt.Println("all index exists: ", result)
}

func TestDeleteIndex(t *testing.T) {
	result := DelIndex("twitter")
	fmt.Println("all index deleted: ", result)
}

func TestCreateIndex(t *testing.T) {
	result := CreateIndex("twitter", mapping)
	fmt.Println("mapping created: ", result)
}

func TestBatch(t *testing.T) {
	tweet1 := Tweet{User: "Jame1",Age: 23, Message: "Take One", Retweets: 1, Created: time.Now()}
	tweet2 := Tweet{User: "Jame2",Age: 32, Message: "Take Two", Retweets: 0, Created: time.Now()}
	tweet3 := Tweet{User: "Jame3",Age: 32, Message: "Take Three", Retweets: 0, Created: time.Now()}
	Batch("twitter", "doc", tweet1, tweet2, tweet3)
}

func TestGetDoc(t *testing.T) {
	var tweet Tweet
	data := GetDoc("twitter", "1")
	if err := json.Unmarshal(data, &tweet); err == nil {
		fmt.Printf("data: %v\n", tweet)
	}
}

func TestTermQuery(t *testing.T) {
	var tweet Tweet
	result := TermQuery("twitter", "doc", "user", "Take Two")
	//获得数据, 方法一
	for _, item := range result.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("tweet : %v\n", t)
		}
	}
	//获得数据, 方法二
	fmt.Println("num of raws: ", result.Hits.TotalHits)
	if result.Hits.TotalHits > 0 {
		for _, hit := range result.Hits.Hits {
			err := json.Unmarshal(*hit.Source, &tweet)
			if err != nil {
				fmt.Printf("source convert json failed, err: %v\n", err)
			}
			fmt.Printf("data: %v\n", tweet)
		}
	}
}

func TestSearch(t *testing.T) {
	result := Search("twitter", "doc")
	var tweet Tweet
	for _, item := range result.Each(reflect.TypeOf(tweet)) {
		if t, ok := item.(Tweet); ok {
			fmt.Printf("tweet : %v\n", t)
		}
	}
}

func TestAggsSearch(t *testing.T) {
	AggsSearch("twitter", "doc")
}
