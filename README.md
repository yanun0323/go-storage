# GO-STORAGE

go-storage is a mini file storage service written in Go.

It stores files in SQLite or MySQL database.

## Requirements

- Go 1.20+
- Docker

## Usage

1. #### Clone this repo to local

2. #### Copy config and edit

```bash
$ cp config.yaml.example config.yaml
```

3. #### Run

```bash
# local run
$ go run main.go

# docker build and run
$ docker build . -t go-storage --no-cache
$ docker run -d -p 8001:8001 --name go-storage go-storage
```

## API

- ### GET

  Get file by filename.

  #### PATH

  `/file/:filename`

  #### Response

  - `Success` `HTTP 200`

    return file blob

  - `Error`

    ```json
    {
      "message": "{any error message here}"
    }
    ```

- ### POST

  Upload file.

  #### PATH

  `/file`

  #### Request

  `multipart/form-data`

  | Key  | Type | Description |
  | ---- | ---- | ----------- |
  | file | file | File        |

  ##### Header

  | Key           | Type   | Description     |
  | ------------- | ------ | --------------- |
  | Authorization | string | Token to verify |

  #### Response

  - `Success` `HTTP 200`

    ```json
    {
      "filename": "{filename}"
    }
    ```

  - `Error`

    ```json
    {
      "message": "{any error message here}"
    }
    ```
