
# fib_API

## フィボナッチ数を返すAPIサービスの開発

local環境構築方法

```bash
git clone https://github.com/higayu624/fib_api.git
cd fib_api
go mod init fib_api
docker compose up --no-deps --build go
# import Errorなどが出る時よしなにモジュールの依存関係を整理してくれる（go.modファイルがあるディレクトリで実行）
go mod tidy
```

単体テストの実行方法
```bash
cd src
go test -r (テストしたい関数名の一部 例：Fibonacci) ./...
```

リリース済みのAPIを叩く方法

```bash
# 今回はフィボナッチ数列の小さい方から99番目の数を取得している。
curl -X GET -H "Content-Type: application/json" "http://fibapi.com/fib?n=99
```

## プロダクトの目的の定義
このプロダクトの目的は、数学の研究を行っている人が研究の過程でフィボナッチ数を知りたいという需要に応えるプロダクトであると定義する

## プロダクトの価値
1 人が計算するのではなく、機械が計算するため速い
2 正確な値
この2つはmustで満たせるプロダクトを目指す

## 採用した技術一覧
Golang Gin Air AWS ECS(Fargate) MVC

## 技術選定の理由

Golang：インタプリタ言語ではなく、コンパイル言語であるため、実行速度が速く、サービスが大きくなって行った時に恩恵を受けやすい。

Gin：Goのrouting系のフレームワークはたくさんある。その中でGinを採用したのは速さ。数あるフレームワークの中でも1番, 2番を争う速さ。さらに今後このプロダクトが意外にも需要があった場合、ユーザの情報などを扱う変更があると想定。ユーザ情報を守るためセキュリティ的に強い特徴があるGinを採用。

Air：ホットリロードをするために導入（変更するたびサーバーを再起動しなくていいため開発速度が早くなる。今回の開発期間はとても短いため、開発の効率化も重要）

AWS ECS(Fargate)：クラウドマネージドなサービスで、運用、保守、監視などをよしなにやってくれるため、アプリケーション開発に集中できるため採用。スピード感を持った開発ができる。サービスが大きくなって行った場合にも使い続けられる。

MVC：今回main.goで全ての処理を書くこともできたが、チーム開発、今後のプロダクトのグロースを意識した開発をしたかったため、読みやすさ、編集のしやすさを意識して採用した。

## ソースコードの構成
コンプロダクトで最初に呼ばれるのはmain.goとなっている。
そしてmain.goでルーターを呼び出し、portを指定してリクエストを待つ状態を作っている。

route.goでGinを利用し、ルーティングを行なっている。

今回は具体的な処理を書くためにController層とModel層を作成した。
Controllerは、データの受け取り、バリデーション、データの加工などを担当させた。
Modelは、今回のプロダクトの価値となるところのロジック（今回はfibonacci数の計算アルゴリズム）を担当させた。

環境系については、docker-copose.ymlとDockerfileでコンテナの構築、.code-workspaceでGoのコード整形ツールの使用、.air.tomlはホットリロードする時の設定、README.mdではこのプロダクトの説明を書いた。

コンパイル後のバイナリはfib_api/tmp配下に置かれている。

## Author

* 東谷有真
* 広島市立大学　大学院　情報科学研究科　知能工学選考　データ科学講座
* higayu624@gmail.com