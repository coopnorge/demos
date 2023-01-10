# Protobuf validation

- [protoc-gen-validate (PGV)](https://github.com/bufbuild/protoc-gen-validate)
  - [Constraint Rules](https://github.com/bufbuild/protoc-gen-validate#constraint-rules)
  - [`validate.proto`](https://github.com/bufbuild/protoc-gen-validate/blob/main/validate/validate.proto)

## ...

- Message with implicit presence requirement: [`recommendations_api.proto`](https://github.com/coopnorge/services-interfaces/blob/cd6050f1f32c3d48a4052b894c76416983032576/proto/coopnorge/offers/recommendations/v1beta1/recommendations_api.proto#L25-L30)
- Extensive use of validation annotations: [`partner.proto`](https://github.com/coopnorge/sapcrm-profile-data/blob/e7397d84621aab1d8bca13e9e3956fd0a0018808/spec/proto/coopnorge/customer/privy/sapcrm/v1/partner.proto)
- Coops Protobuf guidelines regarding presence discipline: [`Guideline: Use the "no presence" discipline`](https://inventory.internal.coop/docs/default/component/coop-playbook/guidelines/api/protocol-buffers-specifications/#use-the-no-presence-discipline)

## Demo Commands

```bash
# Build the devtools
docker-compose build
# Start the service
docker-compose run --rm --service-ports devtools make generate golang-run CLI_ARGS='service run'

# Send a request that will fail validation
docker-compose run --rm grpcurl -plaintext \
    -d '{}' \
    "127.0.0.1:32604" example.greeter.v1.GreeterAPI.Greet

# Send a request that will pass validation
docker-compose run --rm grpcurl -plaintext \
    -d '{"name": "John Smith", "some_value": {}}' \
    "127.0.0.1:32604" example.greeter.v1.GreeterAPI.Greet
```


## Other Commands

```bash
LOCALSTATEDIR=$(pwd)/var docker-compose run --rm devtools make generate
## check the generated docs
google-chrome var/generated/doc/index.html
```
