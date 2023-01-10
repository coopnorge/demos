# example service

```bash
# list targets
make help
```

```bash
# see commands
go run ./cmd/iaeggosvc/ --help
LOGRUS_CONSOLE=1 make golang-run CLI_ARGS='service dump-config'
LOGRUS_CONSOLE=1 make golang-run CLI_ARGS='service run'
```

```bash
socat -v STDIO TCP4:127.0.0.1:$(go run ./cmd/iaeggosvc service dump-config | jq '.Port')
```

```bash
grpc_address="$(ip route get 8.8.8.8 | sed -n 's/^.*src \([^ ]*\).*$/\1/gp'):$(go run ./cmd/iaeggosvc/ service dump-config | jq .Port)"

docker-compose run --rm grpcurl -plaintext "${grpc_address}" list
docker-compose run --rm grpcurl -plaintext "${grpc_address}" describe
docker-compose run --rm grpcurl -plaintext -d '{"service": "example.greeter.v1.GreeterAPI"}' "${grpc_address}" grpc.health.v1.Health.Check

docker-compose run --rm grpcurl -plaintext "${grpc_address}" example.greeter.v1.GreeterAPI.Greet
docker-compose run --rm grpcurl -plaintext -d '{"name": "John Smith"}' "${grpc_address}" example.greeter.v1.GreeterAPI.Greet


```

## Using devtools

```bash
# Build the latest devtools environment
docker-compose build

# See available targets
docker-compose run --rm devtools make help

# Get a shell
docker-compose run --rm devtools bash
docker-compose run --rm devtools make validate

# watch for changes, validate and run
docker-compose run --rm --service-ports devtools
```

## running docker

```
make build-oci
docker run --rm --network host  -it ocreg.invalid/iagotmpl:latest
```
