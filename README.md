# Linxdatacenter Test Task

This is a simple email parser written in go.

## Prerequisites

```
docker
go (if you want to run test)
```

## Installation

Just clone this repo, build docker container and run it

```
git clone http://github.com/kembo91/linxdatacenter-test-task.git
cd linxdatacenter-test-task
docker build -t linxdatacenter/test .
docker run -p 8080:8080 linxdatacenter/test
```

In order to change port, application is running on, just change ENV line in Dockerfile and rebuild image.

## Email API

### Request

`POST /`

```
{
    "req_type": "parseAddress",
    "data": [
        {
            "item":"John Daggett, 341 King Road, Plymouth MA"
        },
        {
            "item":"Alice Ford, 22 East Broadway, Richmond VA"
        },
        {
            "item":"Orville, Thomas, 11345 Oak Bridge Road, Tulsa OK"
        },
        {
            "item": "Terry Kalkas, 402 Lans Road, Beaver Falls PA"
        },
        {
            "item": " Eric Adams, 20 Post Road, Sudbury MA"
        },
        {
            "item": "Hubert Sims, 328A Brook Road, Roanoke VA"
        },
        {
            "item": "Amy Wilde, 334 Bayshore Pkwy, Mountain View CA"
        },
        {
            "item": "Sal Carpenter, 73 6th Street, Boston MA"
        }
  ]
}
```

### Response

```
{
    "res_type":"parseAddress",
    "result":"success",
    "data":"Massachusetts\n..... John Daggett 341 King Road Plymouth Massachusetts\n.....  Eric Adams 20 Post Road Sudbury Massachusetts\n..... Sal Carpenter 73 6th Street Boston Massachusetts\nVirginia\n..... Alice Ford 22 East Broadway Richmond Virginia\n..... Hubert Sims 328A Brook Road Roanoke Virginia\nOklahoma\n..... Orville Thomas 11345 Oak Bridge Road Tulsa Oklahoma\nPennsylvania\n..... Terry Kalkas 402 Lans Road Beaver Falls Pennsylvania\nCalifornia\n..... Amy Wilde 334 Bayshore Pkwy Mountain View California\n"
}
```

## Running tests

There is only one test I thought of that makes sense

```
cd linxdatacenter-test-task
go test ./...
```
