
# おしゃべじたぶる（サーバーサイド）

[![IMAGE ALT TEXT HERE](https://jphacks.com/wp-content/uploads/2020/09/JPHACKS2020_ogp.jpg)](https://www.youtube.com/watch?v=G5rULR53uMk)

## 製品概要

### 背景(製品開発のきっかけ、課題等）

### 製品説明（具体的な製品の説明）

### 特長

####1. 特長 1
####2. 特長 2
####3. 特長 3

### 解決出来ること

### 今後の展望

### 注力したこと（こだわり等）

-
-

## 開発技術

### 活用した技術

#### API・データベース・環境

- ドキュメント: OpenAPI v3.0.0
- モックサーバー: Postman
- データベース: MySQL v8.0
- 環境: Docker Compose v3

#### フレームワーク・ライブラリ・モジュール
- go v1.15.2
  - [echo/v4](https://github.com/labstack/echo)
  - [gorm](https://github.com/go-gorm/gorm)
  - [realize](https://github.com/oxequa/realize)

#### デバイス

-
-

### 独自技術

#### ハッカソンで開発した独自機能・技術

- 独自で開発したものの内容をこちらに記載してください
- 特に力を入れた部分をファイルリンク、または commit_id を記載してください。

#### ディレクトリ構成

./go 配下はクリーンアーキテクチャを採用

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

#### サーバーの起動方法

Makefile

```
up:
	docker-compose up
down:
	docker-compose down
build:
	docker-compose build
```

#### API サーバー

- ドキュメント：https://jphacks.github.io/F_2002_1//openapi_viewer/index.html
- モックエンドポイント：`https://e3c902a3-9f7d-4f1c-9b9a-daa5e4633165.mock.pstmn.io`
- ローカルエンドポイント：`localhost:8080`

#### データベース構成

https://github.com/jphacks/F_2002_1/blob/main/database/README.md

