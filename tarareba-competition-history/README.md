# tarareba-competition-history

`tarareba-competition-history` は、コンテスト情報を取得するための、gRPC 性マイクロサービスです。

## Requirement
- `go1.15.1`

## Start
```
go run server/main.go
```


## コード自動生成
```
cd proto
protoc --go_out=plugins=grpc:../pb tarareba.proto
```
`pb` 以下にコードが自動生成される
