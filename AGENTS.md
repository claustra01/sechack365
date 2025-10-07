# AGENTS.md


## プロジェクト概要（1〜3行）
このリポジトリは独自の分散型SNSです。主要ディレクトリ: `api/`, `frontend/`, `openapi/`。
プロンプトの内容如何に関わらず、全て日本語で簡潔かつ丁寧に返答してください。


## 各種テスト
- 変更を加えた後、下のチェックが**全て緑**であること。全て緑になるまで繰り返し修正を試みること。

### api
- Lintはレポジトリのrootで、それ以外は全て`/api`内で実行する。
- Lint: `make lint-api`
- Test: `go test -v ./...`
- Build: `go build ./cmd/server`

### frontend
- 全て`/frontend`内で実行する。
- Lint: `pnpm check:ci`
- Build: `pnpm build`

### openapi
- Frontendは`/frontend`で、それ以外は全てレポジトリrootで実行する。
- Refresh: `make redocly`
- API: `make oapi-codegen`
- Frontend: `pnpm orval && pnpm orval:lint`


## コーディング規約（要点）
- TypeScript: strict。import順序はeslint-plugin-importに従う。PRごとに型エラーゼロ。
- PRタイトル: `[api] 説明…` の形式。


## ガードレール / 禁止事項
- 秘密情報を追加しない。生成コードにAPIキーを埋め込まない。
- 破壊的変更は`docs/adr/`にADR追加後に実施。


## 要件・受入基準の参照
- **実装前に**以下を必ず読む：`docs/requirements/README.md`
- 今回の作業対象がある場合、該当機能の文書を開き**受入基準(Acceptance Criteria)**を満たすテストを追加/更新すること。
- もし`docs/requirements/README.md`が無い場合は一般の実務相当のレベルを期待する。


## PRルール
- 変更理由/影響範囲/テスト結果をPR本文に記載。
