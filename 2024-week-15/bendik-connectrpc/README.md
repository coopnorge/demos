# ConnectRPC demo

This demo demonstrates how ConnectRPC can be a drop-in replacement of gRPC, but
with additional benefits.

Pros:

1. Drop-in replacement of gRPC. Can replace gRPC server/client/both.
2. Compatible with web-browsers.
3. Can add interceptors/middleware before/after parsing the body.
    (gRPC only support interceptors after parsing the body)
4. Can use standard http.Server middleware for AuthN/Z and CORS.
5. Does automatic transcoding between JSON and Protobuf.
    As a result, it can be called with normal JSON over REST.
6. Uses stdlib's http2-implementation, instead of gRPC's own implementation.
7. Made by the company that makes Buf.

Cons:

1. Slightly different method signature for the handler/client, but still uses
    the same Protobuf-structs (non RPC-related structs).
2. Different interceptor-signatures, so new middlewares must be created.
3. Less widespread, so fewer readily available middlewares exist.

References:

- [Official webpage](https://connectrpc.com/)
