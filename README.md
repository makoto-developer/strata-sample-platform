# strata-sample-platform

[Strata](https://github.com/makoto-developer/strata)（依存関係の可視化・解析ツール）の**動作確認用サンプルシステム**です。

> 📖 Strata の使い方: <https://makoto-developer.github.io/strata/>
実運用の EC プラットフォームを模した、**多言語・複数プロトコル**のマイクロサービス群になっています。

実行して動くサービスではなく、**解析対象として現実的な形**を持たせることが目的です
（依存ライブラリの取得や protoc 生成は不要）。

## 構成

| レイヤ | サービス | 言語 | 公開 API |
| --- | --- | --- | --- |
| BFF | `gateway` | Go (gin) | REST 9 本 → gRPC へ中継 |
| GraphQL | `graphql/gateway` | TS (Apollo Gateway) | スーパーグラフ |
| GraphQL | `graphql/subgraphs/users` | TS (Apollo) | Query/Mutation 5 本・`User` エンティティ所有 |
| GraphQL | `graphql/subgraphs/orders` | TS (Apollo) | Query/Mutation 4 本・`extend type User` |
| GraphQL | `graphql/subgraphs/catalog` | Go (gqlgen) | Query/Mutation 3 本・`Product` 所有 |
| ドメイン | `services/user` | Go | gRPC 5 RPC + REST 4 本 + webhook 受信 |
| ドメイン | `services/catalog` | Go (echo) | gRPC 4 RPC + REST 3 本 |
| ドメイン | `services/inventory` | Go | gRPC 4 RPC |
| ドメイン | `services/payment` | Go (chi) | gRPC 4 RPC + webhook 受信 + 外部決済 API |
| ドメイン | `services/order` | Go | gRPC 4 RPC（注文 Saga） |
| ドメイン | `services/notification` | Go | gRPC 3 RPC + 外部メール API |
| ドメイン | `services/shipping` | Python (FastAPI) | REST 3 本 + 運送会社 webhook |
| ドメイン | `services/search` | TS (Fastify) | REST 3 本 + 外部検索エンジン |
| フロント | `web` | TS | GraphQL 操作 + BFF の REST |

### 意図的に含めてある「解析の題材」

- **gRPC**: 共有 `proto/` を 6 ドメイン分。クライアント呼び出し → RPC → サーバハンドラの接続
- **REST**: gin / echo / chi / net-http / Fastify / FastAPI と、フレームワークごとに違う書き方
- **サービス間 HTTP**: `catalog → search`、`order → shipping`、`web → gateway`、`graphql-* → gateway/user`
- **webhook**: 受信 3 本（IdP・決済プロバイダ・運送会社）と、送信 4 系統（決済・メール・チャット・集荷）
- **GraphQL federation**: `orders` サブグラフが `users` の `User` を `@key` で拡張
- **注文 Saga**: `order` が inventory → payment → shipping → notification を順に呼び、失敗時に補償する
- **アーキテクチャルール**: `strata.config.json` の `forbidden` に「web は各サービスを直接叩かない」を定義

## Strata で見る

```sh
# 構造・API・アーキ図をブラウザで
strata serve /path/to/strata-sample-platform

# CI 用の検査（循環・禁止依存・しきい値）
strata check /path/to/strata-sample-platform

# Markdown レポート
strata report /path/to/strata-sample-platform -o report.md
```

## ライセンス

サンプルコードなので自由に流用して構いません（MIT）。
