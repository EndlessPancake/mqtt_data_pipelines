// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"cloud.google.com/go/bigquery"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

//BigQueryにデータを追加するための構造体、タグで変数とキーを紐づける
type Data struct {
        Id     string    `bigquery:"ID"`
        Datetime time.Time `bigquery:"DATETIME"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	log.Println(string(m.Data))
	return nil
}
// load Puib/Sub message to bigquery
func WriteToBigQuery(name string) {

        // データ格納
        data := Data{}
        data.Id = id
        data.Datetime = time.Now()

        //コンテキスト取得と使用するProject指定
        ctx := context.Background()
        projectID := os.Getenv("GCP_PROJECT")
	datasets  := os.Getenv("GCP_DATASET")
	tablenames := os.Getenv("GCP_TABLENAME")
	
        //BigQuery操作用クライアント
        client, err := bigquery.NewClient(ctx, projectID)
        if err != nil {
                log.Printf("BigQuery接続エラー　Error:%T message: %v", err, err)
                return
        }

        defer client.Close()

        //テーブル操作用アップローダー
        // u := client.Dataset("DATASET").Table("NAMES").Uploader()
        u := client.Dataset(datasets).Table(tablenames).Uploader()
	
        items := []Data{data}

        // data put
        err = u.Put(ctx, items)
        if err != nil {
                log.Printf("データ書き込みエラー　Error:%T message: %v", err, err)
                return
        }
}
