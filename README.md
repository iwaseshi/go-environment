# Environment Golang


## 概要


本リポジトリはGolang(v1.7)による開発で使用する開発環境をローカル上に構築する。
本資源を利用することで、**10分程度**(*1)で以下のことが可能となる。

- PCのローカル環境への変更なく、Golangの開発環境を利用できる。
- PCのローカル環境への変更なく、mysql8.0を利用できる。
- 構築した環境が不要になった場合、直ちに環境を破棄できる。
- 破棄した後も、再び同じ手順で環境を構築し直すことができる。

※OSは不問だが、OS X等のUnixベースのOSでの利用を推奨する。

(*1) : ネットワーク環境によって多少の増減あり

## 事前準備


### 以下のアプリケーションのインストール及び動作が確認できていること。

- Git
  - `git clone`が可能であること。
- Docker
  - コンテナの実行が可能であること。
    - インストール時のPractice用コンテナの起動/終了が行えていればよい。
- Visual Studio Code (以降：VSCode)
  - 拡張機能である[Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)のインストールが済んでいること。


Windowsの場合の追加インストール：

- [Cmder](https://cmder.net/)
  - Unixライクのターミナルソフト。かっこいい
  - 必須ではないが、OS標準のコマンドプロンプトの100倍は推奨。
  - Downloadする場合は「Download Full」を選択する。

Goのインストールはこの時点で必要ない。

※インストールされていても支障はないため、インストール済みのものをアンインストールする必要もない。



### 環境変数の設定

以下2種の環境変数を設定する。
値は任意だが、環境構築後にローカルに立ち上げたmysqlのroot用の認証情報になるため留意しておくこと。

```
${MYSQL_ROOT_USER}
${MYSQL_ROOT_PASSWORD}
```



## 開発環境の構築


事前準備がすべて完了している場合、本ステップに進む。
構築の手順は大きく分けて以下の4ステップのみである。

- 本リポジトリのClone
- docker composeを使用したコンテナの起動
- VSCodeを起動したコンテナとつなげる
- アプリケーションの起動確認



### 本リポジトリのClone

自身のPCの任意のディレクトリにフォルダを作成する。

フォルダ名に特に指定はないが、半角英字で名付けると不要なトラブルを避けることができる。
以下は作成例である。

```shell
OS X   :  ~/development/golang
windows: C:/development/golang
```

次にターミナル上で作成したフォルダに入る。
上記例の場合は、golangフォルダを指す。

```shell
cd  ~/development/golang
```

最後に本資源をcloneする。
GitLab上の「クローン」を押下し、「HTTPSでクローン」の下部にあるURLをクリップボードにコピーする。

![image-20210916153511602](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916153511602.png)

コピー後、ターミナル上で以下のコマンドを入力する。

```shell
git clone {ここにコピーしたURLを貼り付ける}
```

※ユーザ名とパスワードを要求された場合、GitLabにサインした際に使用したユーザ名とパスワードを入力する。
GoogleやGitHub等の連携認証を使用してサインインしていた場合は、その連携元のアプリのユーザ名とパスワードを使用する。

![image-20210916154400721](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916154400721.png)

↑のような表示がされ、フォルダ内に本資源が存在していれば完了。



### docker composeを使用したコンテナの起動

ターミナル上でCloneした資源の内部へ入る。

```shell
cd environment
```

続けてターミナル上でdocker compose upし、コンテナを起動する。

```shell
docker compose up -d --build
```

※初回はイメージファイルをダウンロードする都合、時間がかってしまう。

![image-20210916155101581](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916155101581.png)

↑のような表示がされ、Dockerデスクトップ上でも「go_environment」と「mysql」がRUNNINGされていれば完了。

![image-20210916155303159](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916155303159.png)



### VSCodeを起動したコンテナとつなげる

VSCodeを起動し、左下にある緑色の箇所を押下する。

![image-20210916155514497](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916155514497.png)

**ここ↑**

次にAttach to Running Containerの選択

![image-20210916155821028](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916155821028.png)

go_enviromentを選択

![image-20210916155902167](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916155902167.png)

VSCodeの新しいWindowが表示され、appディレクトリがカレントディレクトリとなっていることを確認する。

このappディレクトリの配下にGolangのコードを置くことで、コーディングした任意のGolangのプログラムを動かすことができる。

以下のようになっていれば完了。

![image-20210916160119994](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916160119994.png)

↑コンテナ名が表示されいることを確認。

またこの時点で、拡張機能をインストールしておくとよい。
拡張機能の検索欄に `@recommended`と入力すれば、筆者の推奨する最低限の拡張機能が表示されるため、
インストールすることを推奨する。

![image-20210916160449829](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916160449829.png)



### アプリケーションの起動確認

最後に本資源にあらかじめ同梱されているGolangのプログラムを実行し、起動確認を行う。

appディレクトリの直下で以下を入力する。（remote containerで開いたVSCode上のターミナルで実行する）

```
cd /go/app/src
```

```go
go run main.go
```

ターミナル上に、`Environment building completed!`と書かれたアスキーアートが表示されれば、環境構築完了。

また、コンテナが起動していれば、[VSCodeを起動したコンテナとつなげる](https://github.com/iwaseatsuya/go-environment#vscodeを起動したコンテナとつなげる)の手順のみで再び同環境を利用できます。

お疲れ様でした。


## Appendix


### コンテナの起動/終了

開発を終えた際に、コンテナも終了させたい場合は以下の手順で行う。
（STOPを選択し、押下）
![image-20210916163956684](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916163956684.png)

開発を再開する場合は再度コンテナを起動させる。
（STARTを選択し、押下）

![image-20210916162407898](https://github.com/iwaseatsuya/go-environment/blob/assets/typora-user-images/image-20210916162407898.png)



### Detabaseの利用

環境構築が完了している場合、Webアプリの開発などに利用するmysqlのコンテナも利用できる状態となっている。

認証情報はdocker-compose.ymlファイルの以下を参考に設定する。

```
MYSQL_DATABASE:
MYSQL_USER:
MYSQL_PASSWORD:
ports:
```

※ymlファイルの上記箇所を任意の値に変更することで、認証情報を任意の情報に変更できる。利用の際は変更を**強く推奨する。**

また、本資源内でGolangのアプリケーションを実行する際に、mysqlのIPを指定する際にlocalhostを指定しても接続がうまく行かないため、

docker-compose.yml内でdb.container_name箇所で宣言しているコンテナ名を使用する（本資源での初期値はmysql_container）

### Flywayの利用

flyway_enviromentコンテナに入ることで、各種flywayコマンドが利用できる。
