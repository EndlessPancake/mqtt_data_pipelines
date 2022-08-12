# BigQuery
## Prepare
- set IAM account for suitable role for BigQuery
- and also gcloud CLI 
### create datasets and tables
```
$ export DATASET=2022summer
$ export NAMES=sensors
$ bq mk ${DATASET}
${Dataset} 'YOUR_PROJECT:DATASET' successfully created.
$ bq mk --table ${DATASET}.${NAMES} config.json
Table 'YOUR_PROJECT:${DATASET}.${NAMES}' successfully created.
```
