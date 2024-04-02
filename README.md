## 起動方法
1. ルートディレクトリにてdocker-compose upコマンドを実施する

### 備考
- named volumeの削除
`docker volume rm app_postgres_data`
`docker-compose down -v`

### わかったこと
- CORSについて今までふわっとしていた理解が少しだけ深まった
  - [Notionページ](https://pinto-waltz-911.notion.site/CORS-86542816a15a486ca2ae4f29f92de5a2)
- ディレクトリ構成やメソッドの作り方はイマイチになってしまったが、上位モジュール（今回はハンドラー）側で下位モジュール（サービス層）のインターフェースを定義して使うということが実際にやってみてよりイメージが明確になった。
- JWTを用いた認証機能が今までより理解できた。今度の研修タスクでJWTについて深ぼる際にもっと詳しく調べてみたい。
- .air.tomlを用いてはいないが、airの使い方と便利さを実感できた。

### 反省点
- ファイル名や機能による分割の仕方が効率的でなく、意味も分かりづらい。
  - modelsに重複しているモデルが多数存在する。
  - ディレクトリ構成などが手探りの状態で作業を開始したため、accountとuserという曖昧な2つのディレクトリの分割になってしまった。 
- respository層では汎用的なメソッドを作るべきだが、Serviceで機能を作る際の都合に合わせて作成してしまったため、メソッドが不必要に多くなってしまった。
- サービスが不要（リポジトリを呼び出すだけ）の箇所がいくつかある。
- Echo標準のJWTのmiddlewareを忘れていてカスタムmiddlewareを作成してしまった。

