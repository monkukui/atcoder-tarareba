## 開発の進め方
ここでは、ローカルにサーバーを立ち上げて、開発 & デバッグを行う方法を記載します。

### step 1. マイクロサービスを立ち上げる。
`atcoder-tarareba` には、二つのマイクロサービスが存在します。

- `tarareba-competition-history`
- `tarareba-algorithms`

の二つです。

これらのサーバーを立ち上げてください。
ターミナルを複数用意する必要があります。

#### tarareba-competition-history
```
cd tarareba-competition-history
go run server/main.go
```
```
run...
```

#### tarareba-competition-algorithms
```
cd tarareba-algorithms
go run server/main.go
```
```
run...
```

### step 2. GraphQL サーバーを立ち上げる
api エンドポイントを提供する、GraphQL サーバーを立ち上げます。

```
cd tarareba-bff
go run server.go
```
```
2020/12/05 15:59:53 connect to http://localhost:8080/ for GraphQL playground
```

[http://localhost:8080](http://localhost:8080) にアクセスすると、GraphQL プレイグラウンドでクエリを試せます。

#### 例 1. ユーザー id から、コンテスト情報を取ってくるやつ

AtCoder ID を引数で渡す必要があります。

```
query {
  contestsByUserID(userID: "monkukui") {
    contestName
    performance
    actualOldRating
    actualNewRating
    optimalOldRating
    optimalNewRating
  }
}
```

### 例 2. パフォーマンス列から、レート推移を取ってくるやつ
正確には、

- performances
- innerPerformances
- isParticipated

の、計 3 つの列を引数で渡す必要があります。

```
query {
    ratingTransitionByPerformance(performances: [1, 2, 10], innerPerformances: [100, 200, 300], isParticipated: [false, false, false]) {
    oldRating
    newRating
  }  
}
```


## step 3. フロントエンドサーバーを立ち上げる
`yarn` が必要

```
cd tarareba-frontend
yarn
yarn start
```


`yarn` で、必要なモジュールをインストールします。
その後、`yarn start` で、ローカルホストを立ち上げます。
