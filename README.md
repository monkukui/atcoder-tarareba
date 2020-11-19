# AtCoder tarareba
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

## アーキテクチャ
全体のアーキテクチャや使用技術をまとめています。
[overview.md](https://github.com/monkukui/atcoder-tarareba/blob/master/docs/overview.md)
をみてください。

    
## イメージ

### ユーザー名入力前
![tarareba_image_2](https://user-images.githubusercontent.com/47474057/99639464-fb020880-2a8a-11eb-8483-3a0232a0f6bb.jpg)

### ユーザー名入力後
![tarareba_image_1](https://user-images.githubusercontent.com/47474057/99639474-fccbcc00-2a8a-11eb-9d32-c9ca55f1bd30.jpg)
