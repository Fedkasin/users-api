# users-api

**Requirements**

- Go version >= 1.12 
- Dependency manager f.e. dep `go get -u github.com/golang/dep/cmd/dep`
- Docker
- Docker-compose 

**How to run**
1. Install go dependencies `dep ensure`
2. Create config file in `config` folder with current environment configuration f.e. for local `config.local.json`
(see `/config/config.go` file SampleConfig var)
3. Run:
- Via Docker `docker-compose up -d`
- `go run main.go`
- Or compile and run `go install && go run users-api`