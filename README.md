# Yojo+han(仮)
SecHack365'24 開発駆動コース仲山ゼミ
個人運用に特化した理想のSNS

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
