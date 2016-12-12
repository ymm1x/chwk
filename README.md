# chwk
チャットワークにメッセージを投げるコマンド

## Usage

環境変数CHATWORK_TOKENにWebUIより取得したトークンを代入。

```sh
$ export CHATWORK_TOKEN=xxxxxxxxxx
$ chwk MESSAGE
$ chwk --room ROOM_ID --to ACCOUNT_ID MESSAGE
```

実行時間の長いバッチの完了の通知などに。

```sh
$ script && chwk "complite"
```
