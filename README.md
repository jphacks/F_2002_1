# social-creators-teamG-server
データベース・APIサーバー

## 使用方法
### サーバーの起動等
Makfile
```
up:
	docker-compose up
down:
	docker-compose down
build:
	docker-compose build
```
### APIサーバー
- ドキュメント：https://khmer495.github.io/social-creators-teamG-server/openapi_viewer/index.html
- モックエンドポイント：`https://72daaa11-a960-440a-bb1c-41e6070cc90b.mock.pstmn.io`
- ローカルエンドポイント：`localhost:8080`

## 使用技術
### 言語
- go v1.15.2
  - [echo/v4](https://github.com/labstack/echo)
  - [gorm](https://github.com/go-gorm/gorm)
  - [realize](https://github.com/oxequa/realize)
### データベース
- MySQL v8.0
### API
- ドキュメント: OpenAPI v3.0.0
- モックサーバー: Postman
### 実行環境
- Docker Compose v3

## ディレクトリ構成
./go配下はクリーンアーキテクチャを採用  
```
.
├── Makefile
├── README.md
├── api
│   └── openapi.yaml
├── database
│   ├── data
│   └── my.cnf
├── docker-compose.yml
├── docs
│   └── openapi_viewer
└── go
    ├── Dockerfile
    ├── entity
    │   └── model
    ├── go.mod
    ├── go.sum
    ├── infrastracture
    ├── interface
    │   ├── controller
    │   └── database
    ├── main.go
    └── usecase
```
参考：
- [Clean ArchitectureでAPI Serverを構築してみる](https://qiita.com/hirotakan/items/698c1f5773a3cca6193e)
- [Golang - EchoとGORMでClean Architecture APIを構築する](https://qiita.com/so-heee/items/0cca93008eae635c642a)

## データベース構成
https://github.com/Khmer495/social-creators-teamG-server/blob/main/database/README.md
