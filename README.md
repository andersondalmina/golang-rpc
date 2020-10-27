# Golang RPC

### How to compile

With Golang installed just run `make build`

### How to Execute

Copy the `.env.example` file to `.env` and ajust for your needs.

Start containers: `docker-compose up -d`

Server: `./golang_rpc server [host] [port]`

Client: `./golang_rpc client [host] [port]`
