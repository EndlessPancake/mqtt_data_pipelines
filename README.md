# mqtt_data_pipelines
- MQTTクライアントからのデータパイプライン実装検討/実証
- InjestからStoreまでを集中的に検討し、DataStudioやLooker等による可視化部分は未実装
## 実装概要
- 模擬データに2014年Uber提供のPublicデータを利用
```
|MQTT Clinent模擬データ| --> [GKE(Data Injest and route)] --> [cloud pub/sub] <-- [cloud dataflow] --> [Google BigQuery]
```
### データ入力(Endpoint)
  - "MQTT gateway" 外界からの入り口を構成
  - GCP IoT Core の様なPublic Cloud機能には依存しない構成とする
### データハブ(中継)
  - GCP pub/sub を採用
  
### データパイプライン
  - cloud dataflow(Apache Airflow)
  - *cloud functions (Streamデータを想定したため、不採用)*
### メタデータストア
  - bigquery を採用
