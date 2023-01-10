package grpcsvc

import (
	"context"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	greeterpb "gitlab.com/aucampia/eg/service-golang/internal/generated/proto/example/greeter/v1"
	typesv1 "gitlab.com/aucampia/eg/service-golang/internal/generated/proto/example/types/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"
)

type testConfine struct {
	t        *testing.T
	svc      *Service
	srv      *grpc.Server
	listener *bufconn.Listener
}

func newTestConfine(t *testing.T, svc *Service) (serverTestHelper *testConfine) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)

	serverOptions := []grpc.ServerOption{}
	srv := grpc.NewServer(serverOptions...)
	greeterpb.RegisterGreeterAPIServer(srv, svc)
	go func() {
		if err := srv.Serve(listener); err != nil {
			panic(fmt.Errorf("failed to start grpc server: %w", err))
		}
	}()
	return &testConfine{t: t, svc: svc, srv: srv, listener: listener}
}

func (c *testConfine) connectClient() (ctx context.Context, conn *grpc.ClientConn, client greeterpb.GreeterAPIClient) {
	ctx = context.Background()
	dialer := func(context.Context, string) (net.Conn, error) {
		return c.listener.Dial()
	}
	conn, err := grpc.DialContext(ctx, "bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		c.t.Fatalf("Failed to dial context: %v", err)
	}
	client = greeterpb.NewGreeterAPIClient(conn)
	return ctx, conn, client
}

func (c *testConfine) finish() {
	c.srv.Stop()
}

func TestSelect(t *testing.T) {
	svc, err := NewService()
	if err != nil {
		t.Fatal(err)
	}

	svcc := newTestConfine(t, svc)
	defer svcc.finish()
	ctx, conn, client := svcc.connectClient()
	defer conn.Close()

	testCases := []struct {
		name string
		note string
		req  *greeterpb.GreetRequest
		resp *greeterpb.GreetResponse
	}{
		{
			name: "empty",
			req:  &greeterpb.GreetRequest{Name: "Ola Nordmann", SomeValue: &typesv1.SomeType{Value: "some value"}},
			resp: &greeterpb.GreetResponse{Message: "Hello Ola Nordmann"},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Logf("testCase = %#v", testCase)
			response, err := client.Greet(ctx, testCase.req)
			if err != nil {
				t.Fatal(err)
			}
			assert.True(t, proto.Equal(testCase.resp, response))
		})
	}
}
