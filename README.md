## 概要
- ユーザー登録と他のユーザーの検索、自身の画像、ニックネームが編集可能。
- 今回はアプリの目的等はなく、GoでJWTと画像を扱っったみたかったため作成。

## 起動方法
1. ルートディレクトリにて`docker-compose up`コマンドを実施する
1. [http://localhost:3000/login](http://localhost:3000/login) へアクセスする(Next.jsのビルドに多少時間を要する）

### ボリューム削除
- named volumeの削除<br>
`docker volume rm go-next-postgres_postgres_data`<br>
`docker-compose down -v`

## わかったこと
- CORSについて今までふわっとしていた理解が深まった。
  - [Notionページ](https://pinto-waltz-911.notion.site/CORS-86542816a15a486ca2ae4f29f92de5a2)
- ディレクトリ構成やメソッドの作り方はイマイチになってしまったが、上位モジュール（今回はハンドラー）側で下位モジュール（サービス層）のインターフェースを定義して使うということが実際にやってみてよりイメージが明確になった。
- JWTを用いた認証機能が今までより理解できた。
- airの使い方と便利さを実感できた。

## 反省点
### Go
- ファイル名や機能による分割の微妙で、意味が分かりづらい。
- modelsに重複が多数存在する。
- Echo標準のJWTのmiddlewareを使わずカスタムmiddlewareを作成してしまった。
