tCoder tarareba
「このコンテストがなかっ**たら**、あのコンテストに参加していなけ**れば**」

AtCoder tarareba は、過去を改竄してレートを最大化するサービスです。

## どんなことができるサービスか？
最初はミニマムな機能を実装して、後から機能追加できたらいいと思っています。
他にもどんな機能があればいいか、何かあれば教えてください。

### レートを最大化させる
[参考：olphe くんのツイート](https://twitter.com/_olphe/status/1221687722324049920?s=20)

- 入力：AtCoder id
- 出力：最適な、出場するコンテストのサブセットと、架空のレート推移
- 何が嬉しいか：「失敗したコンテストたちに出ていなかったら、俺今頃暖色コーダーだったのに...」みたいな感じがわかる

### 出場したコンテストを自由に選べる（仮）

- 入力：AtCoder id, どのコンテストに出場したか
- 出力：架空のレート推移
- 何が嬉しいか：「AGC〇〇 にさえ出てなかったら、俺今頃暖色コーダーだったのに...」みたいな感じがわかる


## 使用技術
**（注意）**
個人的に試してみたい技術や、言語、フレームワークがあります。そのため、このプロダクトを開発することのみを目的とした場合、アーキテクチャが冗長に感じる場合があります。ごめん。

- Go（サーバーサイド全般）（主要アルゴリズム）
- gqlgen（Go で GraphQL を書くためのフレームワーク。サーバーサイドと GraphQL をつなぐ）[参考](https://gqlgen.com/getting-started/)
- grpc-go [参考](https://github.com/grpc/grpc-go)
- React （今一番いけてるフロントエンドのフレームワーク。AtCoder Problems もこれ） [参考](https://ja.reactjs.org/)
- Appolo Client（GraphQL を React につなぐためのフレームワーク）[参考](https://www.apollographql.com/docs/react/)

### 概要
大きく分けて、やることは 3 つだと思っています。

- AtCoder id を受け取り、パフォーマンス列（や、その他コンテストに関する情報）を構築する処理
    - kenkooo さんあたりが api を提供しているなら、それを使う
        - api がないなら、自前で AtCoder からスクレイピングしてくる
            
            - パフォーマンス列を受け取り、レートを最大化させるアルゴリズムの実装
                - サービスの根幹ロジック
                    - DP + 復元をすれば良い？
                        - AtCoder のレーティングシステムを知る必要がある（[参考: qiita](https://qiita.com/anqooqie/items/92005e337a0d2569bdbd)）
                         
                         - フロントエンドの実装
                             - React を使う
                                 

                                 これら 3 つの間のデータやりとりに、`GraphQL` という技術を使います。（ここがかなりマニアックです。一般的には、`REST` と呼ばれるプロトコルに従って、WEB 開発をします。なんか、`REST` に変わるナウいやつが、`GraphQL` だと思ってください。今回、この `GraphQL` を試してみたいというモチベーションがあります。）（monkukui がやる予定です。）

                                 WEB への公開は GCP + Kubernetes を考えています。が、難しそうであれば、heroku で公開します。（monkukui がやります。）


#### アーキテクチャをまとめた図
![tarareba_architecture](https://user-images.githubusercontent.com/47474057/99642940-71a10500-2a8f-11eb-91d5-626f8fdf5567.jpg)


#### [参考: メルカリのブログ](https://engineering.mercari.com/blog/entry/2019-12-14-110000/)

    
## イメージ

### ユーザー名入力前
![tarareba_image_2](https://user-images.githubusercontent.com/47474057/99639464-fb020880-2a8a-11eb-8483-3a0232a0f6bb.jpg)

### ユーザー名入力後
![tarareba_image_1](https://user-images.githubusercontent.com/47474057/99639474-fccbcc00-2a8a-11eb-9d32-c9ca55f1bd30.jpg)

