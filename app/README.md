# ERVCP

ERVCP is a first in it's class, one of a kind Enterprise-Ready Visitor Counting Platform. Every time someone visits your website, the counter goes up.

## Dependencies

This application is written in [Go](https://golang.org/). Supported version is `1.12`.

This application uses [Redis](https://redis.io/) to store the data. Supported version is `5.0`.

## Configuring

This application is configured via environment variables:

* `ERVCP_PORT` - port, on which web app will run; defaults to `8080`
* `ERVCP_DB_HOST` - Redis host
* `ERVCP_DB_PORT` - Redis port
* `ERVCP_DB_PW` - Redis password

## Building and running

To build the application, run:

    go build

This will create binary `ervcp` that you can run:

    ervcp

By default, ERVCP run on http://localhost:8080.


## Credits

Favicon by [icon8](https://icons8.com/icons/set/favicon) <3