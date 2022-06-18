# U.M.A. Industrial Engineering Master thesis

start up db and envoy for rpc proxy

    make start 

in the raspberry

    ngrok tpc 8082
    go run client.go

in localhost

    go run manger.go
    yarn start