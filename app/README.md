# ERVCP

ERVCP is a first, one-of-a-kind Enterprise-Ready Visitor Counting Platform. Every time someone visits your website, the counter goes up.

## Dependencies

This application uses Redis to store the data. Supported version is `5.0`.

## Configuring

This application is configured via environment variables:

* `ERVCP_PORT` - port, on which web app will run; defaults to `8080`
* `ERVCP_DB_HOST` - Redis host
* `ERVCP_DB_PORT` - Redis port
* `ERVCP_DB_PW` - Redis password

## Running

Simply type:

    go run main.go

in your terminal, then open the app in your browser. By default, it'll run on http://localhost:8080.


## Credits

Favicon by [icon8](https://icons8.com/icons/set/favicon) <3