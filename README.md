# chwk
チャットワークにメッセージを投げるコマンド

## Usage

ひとまず拾ってきてbin下に入れて実行権限つける。

```sh
$ wget https://github.com/blue1st/chwk/releases/download/v_0.0.1/chwk_linux_x86_64.zip -O chwk.zip
$ unzip chwk
$ mv chwk ~/bin/
$ chmod +x ~/bin/chwk
$ chwk --help
Usage of chwk:
  -r int
    	Room ID (short)
  -room int
    	Room ID
  -t int
    	Account ID (short)
  -to int
    	Account ID`
```

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
