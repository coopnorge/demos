# gRPC Health Checking


## Old style health checks:

example: https://github.com/coopnorge/services-interfaces/blob/7fb2bee13083e5c06614103c9716d0ec086d981b/proto/coopnorge/platform/pseudonymisation/v1beta1/pseudonymisation_api.proto#L23-L27

## gRPC Health Checking Protocol


- [GRPC Health Checking Protocol](https://github.com/grpc/grpc/blob/master/doc/health-checking.md)
- Guideline:

## Code example:

https://github.com/coopnorge/products-information-service/blob/a03c41c937cb2ad66b4790e9211a3057ea0f802e/internal/server/server.go#L163-L166

## Calling it:

```
docker-compose run --rm grpcurl -plaintext -d '{"service": "coopnorge.product.v1beta1.InformationAPI"}' "devtools:41829" grpc.health.v1.Health.Check
```

## Load Balancer

https://github.com/coopnorge/products-infrastructure/blob/9e18512cf61793ed65f654354a7fff068932a7bf/kubernetes/base/products-information-service.yaml#L177-L189

