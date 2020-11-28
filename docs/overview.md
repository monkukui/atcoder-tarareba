# AtCoder tarareba
「このコンテストがなかっ**たら**、あのコンテストに参加していなけ**れば**」

AtCoder tarareba は、過去を改竄してレートを最大化するサービスです。


## 使用技術
**（注意）**
興味ドリブンで技術選定しています。そのため、アーキテクチャが冗長に感じる可能性があります。

- Go（サーバーサイド全般）（主要アルゴリズム）
- gqlgen（Go で GraphQL を書くためのフレームワーク。サーバーサイドと GraphQL をつなぐ）[参考](https://gqlgen.com/getting-started/)
- grpc-go [参考](https://github.com/grpc/grpc-go)
- React （今一番いけてるフロントエンドのフレームワーク。AtCoder Problems もこれ） [参考](https://ja.reactjs.org/)
- Appolo Client（GraphQL を React につなぐためのフレームワーク）[参考](https://www.apollographql.com/docs/react/)

## アーキテクチャ
![S__1612452](https://user-images.githubusercontent.com/75159676/100497467-1cf03f00-319f-11eb-973e-3d306865b4f9.jpg)



`atcoder-tarareba` は、マイクロサービスアーキテクチャを採用しています。

現在、二つのマイクロサービスが存在します。

- `tarareba-algorithms`：レート最大化や、レート計算を請け負う
- `tarareba-competition-history`：コンテスト情報やパフォーマンスを取得する

`tarareba-bff` は、これら二つのマイクロサービスから情報を取得し、フロントエンドに提供する役割を果たす GraphQL サーバーです。（BFF とは、Backend for Frontend の略です）

`tarareba-frontend` は、TypeScript（React）で記述されたフロントエンドサーバーです。Apollo Client というフレームワークを用いて、上記の GraphQL サーバーと接続しています。

## 処理の流れ
処理の全体の流れは以下の通りです。

競技プログラミングに例えると、インタラクティブ問題です。インタラクティブ問題を解いているつもりで考えてください。

上の図と合わせてみると、理解しやすいと思います。

- `tarareba-frontend` は、ユーザーから `user_id（AtCoder ID）` を入力として受け取り、`tarareba-bff` にリクエストを投げる。
- `tarareba-bff` は、`tarareba-frontedn` からリクエストを受け取り、`tarareba-competition-history` にリクエストを投げる。コンテスト情報を頂戴！！
- `tarareba-competition-history` は、ユーザーのコンテスト情報を渡す。（https://atcoder.jp/users/monkukui/history っぽいやつ）
- `tarareba-bff` は、`tarareba-competition-history` からコンテスト情報を受け取り、パフォーマンス列を `tarareba-algorithms` に渡す。レートを最大化してくれ！
- `tarareba-algorithms` は、レートを最大化するアルゴリズムを走らせる。DP かもしれないし、貪欲法かもしれない。
- `tarareba-bff` は、`tarareba-algorithms` から復元つきのレート最大化情報を受け取り、よしなに形式をまとめて、`tarareba-frontend` に渡す。
- `tarareba-frontend` は、`tarareba-bff` からレスポンスを受け取り、WEB ページに表示させる。

### マイクロサービスアーキテクチャの特徴

マイクロサービスアーキテクチャの特徴として、

- 責務がそれぞれ独立しているので、開発・テストが閉じた状態でできる
- 変化に対応しやすい。アルゴリズムを変更したければ、`tarareba-algorithms` だけを修正すればよい

などが挙げられます。
もし、`tarareba-algorithms` と `tarareba-competition-history` を独立に処理できるなら、
並行処理（concurrent）ができます。今回は、

`tarareba-competition-history` からコンテスト情報を取得し、それを `tarareba-algorithms` に渡す必要があるので、並行処理ができません。
（そもそも、独立じゃないのにマイクロサービスにして分離しているのは、破綻している気がしてきました。）
