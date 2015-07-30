# loghttpserver

A noop webserver that logs its requests to stdout. You know, for tests.

## Installation

```
go get github.com/slank/loghttpserver
```

## Usage

```
loghttpserver [options]

  -bind="0.0.0.0:9999": IP address and port to bind to
  -response=200: HTTP response code
```
