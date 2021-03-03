package model

import (
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	//"github.com/jmoiron/sqlx"
	//"encoding/json"
	//"context"
	//"fmt"
	//"net/http"
	//"strings"
	//"log"
	//"github.com/farmerx/elasticsql"
	//"github.com/jmoiron/sqlx"
	"gopkg.in/olivere/elastic.v6"
)

//实例化esclient
var EsClient *elastic.Client

func EsConnect() {
	host := "http://192.168.1.143:9200" //测试机器
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
		fmt.Println("es client error！")
	} else {

		fmt.Println("es client success！")
	}
	EsClient = client
	//defer client.Stop()
}
