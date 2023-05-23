



A brief description of what this project does and who it's for

# Golang Rest api with Mysql and authentication with Prometheus and Grafana monitoring.




## Register User



```sh
curl --location --request POST 'http://localhost:4000/api/v1/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "dad",
    "username": "bayot",
    "email": "bayot@chupa.com.",
    "password": "kaonboto1"
}'
```

## Generate Token



```sh
curl --location --request POST 'http://localhost:4000/api/v1/token' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "bayot@chupa.com.",
    "password": "kaonboto1"
}'
```


## Get all Users



```sh
curl --location --request GET 'http://localhost:4000/api/v2/allusers' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJheW90IiwiZW1haWwiOiJiYXlvdEBjaHVwYS5jb20uIiwiZXhwIjoxNjg0NDg4OTMzfQ.xUyNAGtrGafztd9llRSFWpdCNvsZKyWRkNCo0jmOQeA'
```

## Get User By Id



```sh
curl --location --request GET 'http://localhost:4000/api/v2/user/8' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJheW90IiwiZW1haWwiOiJiYXlvdEBjaHVwYS5jb20uIiwiZXhwIjoxNjg0NDg4OTMzfQ.xUyNAGtrGafztd9llRSFWpdCNvsZKyWRkNCo0jmOQeA'
```

## Update User By Id



```sh
curl --location --request PUT 'http://localhost:4000/api/v2/user/10' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJheW90IiwiZW1haWwiOiJiYXlvdEBjaHVwYS5jb20uIiwiZXhwIjoxNjg0NDg4OTMzfQ.xUyNAGtrGafztd9llRSFWpdCNvsZKyWRkNCo0jmOQeA' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "dad"
    
}'
```

## Dlete User By Id



```sh
curl --location --request DELETE 'http://localhost:4000/api/v2/user/2' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpc3Ryb3llcjMiLCJlbWFpbCI6ImFtYm90MyIsImV4cCI6MTY4NDMxMDA1Mn0.BT0LNHN_aZX8ki2xonTETcuewuygCBGQolFEWejuLvg'

```

Mysql Connection string

```sh

"root:Sethjoe12345@tcp(localhost:3306)/test_demo?parseTime=true"
