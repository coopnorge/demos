# Custom errors in Go

Demonstration of how to do implement custom errors in Go.

The PR comment that prompted this demo:
<https://github.com/coopnorge/idp-user-service/pull/2564#issuecomment-2724283941>

Implemented for this demo session:
<https://github.com/coopnorge/engineering-issues/issues/442>

Go released native support for `errors.Is`, `errors.As`, and the `Unwrap` method
(and the `%w` formatting verb) in
[Go 1.13](https://go.dev/doc/go1.13#error_wrapping).

Popular multi-error libraries before Go released native support:

- [`github.com/pkg/errors`](https://github.com/pkg/errors)
- [`github.com/uber-go/multierr`](https://github.com/uber-go/multierr)
