# Grafana Embed

## Requirements

- docker
- golang

## How to use

1. generate secret key:

    ```shell
    openssl rand -base64 172 | tr -d '\n'
    ```

1. generate `jwks.json` and move to grafana config folder:

    ```shell
    go run cmd/jwk-generator/main.go --secret-key=xxxxxxx --key-id=grafana-embed-example
    mv jwks.json grafana/
    ```

1. `make setup-containers`
1. `make setup-migration`
1. open `http://localhost:3000`
1. create data source `mysql`
1. import dashboard from JSON
1. copy embed link
1. run server

    ```shell
    go run cmd/server/main.go --grafana-dashboard-url=http://localhost:3000/d/xxxxxxx/xxxxxxxx --secret-key=xxxxxxx --key-id=grafana-embed-example
    # e.g.
    # go run cmd/server/main.go --grafana-dashboard-url=http://localhost:3000/d/adoc5a1k9gzcwb/dashboard --secret-key=xxxxxxx --key-id=grafana-embed-example
    ```

1. open `http://localhost:8000` and watch effect
