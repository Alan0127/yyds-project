package elastic

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
	"yyds-pro/model"
)

var DefaultElasticClient *elastic.Client

//初始化一个elastic客户端
func InitElastic(config model.AppConfig) (err error) {
	client, err := elastic.NewClient(
		elastic.SetURL(config.App.ElasticSearch.Url+":"+config.App.ElasticSearch.Port),
		//elastic.SetURL("http://159.75.35.70:9200"),
		//elastic.SetBasicAuth(config.App.ElasticSearch.User,config.App.ElasticSearch.Password),
		//elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetSniff(false),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	DefaultElasticClient = client
	return
}
