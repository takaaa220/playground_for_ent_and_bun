# シンプルなブログアプリ

## 概要

ユーザーが投稿を作成・編集・削除できるブログアプリを構築します。基本的な CRUD 操作を試すのに適しています。

## 要件

### エンティティ

#### User（ユーザー）

- ID
- 名前
- メールアドレス

#### Post（投稿）

- ID
- タイトル
- コンテンツ
- 作成日時
- 更新日時
- 公開状態
- 公開日時
- User への外部キー

### 機能

- ユーザーの作成、取得、更新、削除。
- 投稿の作成、取得、更新、削除。
- 投稿ごとのユーザー情報の取得（リレーションの確認）。
- ユーザーごとの投稿リストを取得。
