//go:build tools
// +build tools

package root

//nolint:golint
import (
	_ "github.com/GoogleCloudPlatform/protoc-gen-bq-schema" // toolchain
	_ "github.com/bufbuild/buf/cmd/buf"                     // toolchain
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking" // toolchain
	_ "github.com/bufbuild/buf/cmd/protoc-gen-buf-lint"     // toolchain
	_ "github.com/go-task/task/v3/cmd/task"                 // toolchain
)
