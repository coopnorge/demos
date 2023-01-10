//go:build tools
// +build tools

package root

//nolint:golint
import (
	_ "github.com/cortesi/modd/cmd/modd" // toolchain
	_ "github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen" // toolchain
	_ "github.com/golang/mock/mockgen/model"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc" // toolchain
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"  // toolchain
)
