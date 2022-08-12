# mqtt_data_pipelines
- 大規模MQTTクライアントからのデータパイプライン実装検討/実証
## 全体機能概要
- データ入り口
  - "MQTT gateway" 外界からの入り口を構成
  - GCP IoT Core の様なPublic Cloud機能には依存しない構成とする
- データハブ(中継)
  - GCP pub/sub を採用
- データパイプライン
  - cloud functions
  - cloud dataflow(Apache Airflow)
  
- データ保管
  - bigquery を採用
- Data studio or looker
  - サンプルデータなのでCountだけを表示する
### データフロー概要
- Simple workload 
```
|MQTT Clinent or Benchmark| --> [GKE(Data Injest and route)] --> [cloud pub/sub] <-- [cloud functions] --> [BigQuery] <-- [Data Studio]
```
- Heave workload
|MQTT Clinents| --> [GKE(Data Injest and route)] --> [cloud pub/sub] <-- [dataflows(airflow)] --> [BigQuery] <-- [Data Studio]
### MQTT Client 
- python3 + paho(mqtt module) 
- [go benchmark](https://github.com/krylovsk/mqtt-benchmark)
### Data Injest and route on GKE
- IN : mqtt core plugin
- OUT: gcp pub/sub plugin
