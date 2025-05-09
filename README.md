# YojoHan - prototype

## OpenAPI Docs
https://claustra01.github.io/sechack365/

## DBDocs
https://dbdocs.io/shinya430.g/sechack365

## Setup
- `.env`, `nginx/default.crt`, `nginx/default.key`を配置する
  - 開発環境の場合は`make dev-cert`で生成できる
- `make env`で必要な環境をインストール
- `make migrate`でDBマイグレーション
- `make up`で起動
- `make down`で停止
