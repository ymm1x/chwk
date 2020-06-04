# chwk

A command line tool for sending messages or creating tasks to a chat room via Chatwork API v2.

## Installation

```sh
git clone https://github.com/ymm1x/chwk
cd chwk
go install .
```

## Environment variables

### `CHATWORK_TOKEN` (required)

```sh
export CHATWORK_TOKEN=xxxxxxxxxx
```

## Usage

### Send a message

```sh
# Send to a my chat of chatwork token's user
chwk MESSAGE

# With pipeline
echo MESSAGE | chwk

# Send to a speficied room
chwk --room ROOM_ID --to ACCOUNT_ID MESSAGE
```

### Create a new task

```sh
chwk -task "new task!"
```
