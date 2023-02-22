# Test Project

## Prerequisites

This project requires [Docker](https://docs.docker.com/get-docker/) and [Task](https://taskfile.dev/installation/) to be installed.

## Starting Project Locally

```sh
task run
```

This will start up the project at http://localhost:8080

Executing `task run:docker` will start the project in the detached docker container.

Run `task` without arguments to see all available commands.

## Project Structure

```
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── model
│   └── unconfirmed.go
├── README.md
├── route
│   └── route.go
└── Taskfile.yaml
```

### Data Models
The `model` directory contains collections models, which is basically the structure of the document persisted
in the particular collection.

For example:

```golang
type SpendingOutpoint struct {
	N int64 `json:"n"`
	TxIndex int64 `json:"tx_index"`
}

type PrevOut struct {
	Addr string `json:"addr"`
	N int64 `json:"n"`
	Script string `json:"script"`
	SpendingOutpoints []SpendingOutpoint `json:"spending_outpoints"`
	Spent bool `json:"spent"`
	TxIndex int64 `json:"tx_index"`
	Type int64 `json:"type"`
	Value int64 `json:"value"`
}

type Input struct {
	Index int64 `json:"index"`
	PrevOut PrevOut `json:"prev_out"`
	Script string `json:"script"`
	Sequence int64 `json:"sequence"`
	Witness string `json:"witness"`
}

type Tx struct {
	DoubleSpend bool `json:"double_spend"`
	Fee int64 `json:"fee"`
	Hash []byte `json:"hash"`
	Inputs []Input `json:"inputs"`
	LockTime int64 `json:"lock_time"`
	Outs []PrevOut `json:"out"`
	Rbf bool `json:"rbf"`
	RelayedBy string `json:"relayed_by"`
	Size int64 `json:"size"`
	Time int64 `json:"time"`
	TxIndex int64 `json:"tx_index"`
	Ver int64 `json:"ver"`
	VinSz int64 `json:"vin_sz"`
	VoutSz int64 `json:"vout_sz"`
	Weight int64 `json:"weight"`
}

type Unconfirmed struct {
	DoubleSpend bool `json:"double_spend"`
	Fee int64 `json:"fee"`
	Hash []byte `json:"hash"`
	Id uuid.UUID `json:"id" tigris:"primaryKey:1,autoGenerate"`
	Inputs []Input `json:"inputs"`
	LockTime int64 `json:"lock_time"`
	Outs []PrevOut `json:"out"`
	Rbf bool `json:"rbf"`
	RelayedBy string `json:"relayed_by"`
	Size int64 `json:"size"`
	Time int64 `json:"time"`
	TxIndex int64 `json:"tx_index"`
	Txs []Tx `json:"txs"`
	Ver int64 `json:"ver"`
	VinSz int64 `json:"vin_sz"`
	VoutSz int64 `json:"vout_sz"`
	Weight int64 `json:"weight"`
}
```

This model types can be modified to add new fields to the document.

### Routes

The `route/routes.go` defines SetupCRUD function which is used in the `main.go` to set up [Gin](https://github.com/gin-gonic/gin)
Web framework CRUD routes for every collection model.
Once project is started, they can be tested using curl commands.

For example:

#### Create document in the `unconfirmed` collection:
```
curl -X POST "localhost:8080/unconfirmed" -H 'Content-Type: application/json' 
    -d "{ JSON document body corresponding to the model.Unconfirmed }"
```

#### Read document from the `unconfirmed` collection:
```
curl -X GET "localhost:8080/unconfirmed/{document id}"
```

#### Delete document from the `unconfirmed` collection:
```
curl -X DELETE "localhost:8080/unconfirmed/{document id}"
```

Full Tigris documentation [here](https://docs.tigrisdata.com).

Be brave. Have fun!
