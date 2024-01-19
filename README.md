
# fib_API

## フィボナッチ数を返すAPIサービスの開発

## Golang Air AWS ECS(Fargate)

実行方法

```bash
git clone https://github.com/higayu624/fib_api.git
cd fib_api
go mod init fib_api
docker compose up --no-deps --build go
```

## 技術選定の理由

Golang：インタプリタ言語ではなく、コンパイル言語であるため、実行速度が速く、サービスが大きくなって行った時に恩恵を受けやすい。

Air：ホットリロードをするために導入（変更するたびサーバーを再起動しなくていいため開発速度が早くなる）

AWS ECS(Fargate)：クラウドマネージドなサービスで、運用、保守、監視などをよしなにやってくれるため、アプリケーション開発に集中できるため採用。サービスが大きくなって行った場合にも使い続けられる。

## Author

* 東谷有真
* 広島市立大学
* higayu624@gmail.com