# mqtt_data_pipelines
- 大規模MQTTクライアントからのデータパイプライン実装検討
## 全体機能概要
- データ入り口
  - MQTT gatewayの様な外界からの入り口を構成
  - GCP IoT Core の様なPublic Cloud機能には依存しない構成を取る
- データハブ(中継)
  - GCP pub/sub を採用
- データパイプライン
  - cloud functions を採用
  * *GCP dataflow による実装は別途検討*
  * *GKE上にfluentおよびpluginで実装も可能*
- データ保管
  - bigqueryを採用
- data studio or looker
  - サンプルデータなのでCountだけを表示する
## GCP上のデータフロー
- 
```
|MQTT Clinent or Benchmark| --> [GKE(Data Stream pipelines)] --> [cloud pub/sub] <-- [cloud functions] --> [BigQuery] <-- [Data Studio]

```
### Client 
- [go client](https://github.com/krylovsk/mqtt-benchmark)
### GKE(Data Stream pipelines)
- IN:(fluent-bit + mqtt plugin) , OUT(fluent-bit +pub/sub plugin)
### 
- 
