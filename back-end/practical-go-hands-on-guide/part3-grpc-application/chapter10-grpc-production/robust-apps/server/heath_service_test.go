package main

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	healthsvc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"net"
	"robust-app/utils"
	"testing"
)

func createHealthCheckClient(l *bufconn.Listener) (healthsvc.HealthClient, error) {
	dialer := func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(),
		"",
		grpc.WithInsecure(),
		grpc.WithContextDialer(dialer),
	)
	if err != nil {
		return nil, err
	}

	return healthsvc.NewHealthClient(client), nil
}

func TestHealthService_EmptyCheck(t *testing.T) {
	l, _ := startTestServer()
	client, err := createHealthCheckClient(l)
	checkError("Creating Dial", err)

	res, err := client.Check(
		context.Background(),
		&healthsvc.HealthCheckRequest{},
	)
	checkError("Response", err)

	serviceStatus := res.Status.String()
	checkStringTypeResponse(t, "Status", "SERVING", serviceStatus)
	printResponse(t, res)
}

func TestHealthService_UserCheck(t *testing.T) {
	l, _ := startTestServer()
	client, err := createHealthCheckClient(l)
	checkError("Creating Dial", err)

	res, err := client.Check(
		context.Background(),
		&healthsvc.HealthCheckRequest{
			Service: "Users",
		},
	)
	checkError("Response", err)

	serviceStatus := res.Status.String()
	checkStringTypeResponse(t, "Status", "SERVING", serviceStatus)
	printResponse(t, res)
}

func TestHealthService_UnknownCheck(t *testing.T) {
	l, _ := startTestServer()
	client, err := createHealthCheckClient(l)
	checkError("Creating Dial", err)

	_, err = client.Check(
		context.Background(),
		&healthsvc.HealthCheckRequest{
			Service: "UnknownService",
		},
	)
	if err == nil {
		t.Fatalf("Expected non-nil error but got nil error")
	}

	expectedErr := status.Errorf(codes.NotFound, "unknown service")

	if !errors.Is(err, expectedErr) {
		t.Errorf("Expected error: %v, Got: %v\n", expectedErr, err)
	}

	printResponse(t, err)
}

func TestHealthService_UserWatch(t *testing.T) {
	l, h := startTestServer()
	client, err := createHealthCheckClient(l)
	checkError("Creating Dial", err)

	stream, err := client.Watch(
		context.Background(),
		&healthsvc.HealthCheckRequest{
			Service: "Users",
		},
	)
	checkError("Creating Health Check Watch Stream", err)

	// Serving status
	res, err := stream.Recv()
	checkError("Streaming", err)

	if res.Status != healthsvc.HealthCheckResponse_SERVING {
		t.Errorf("Expected SERVING but got: %s\n", utils.GetJsonStringUnsafe(res.Status))
	}
	printResponse(t, res)

	// Not-Serving status
	updateHealthCheckService(h,
		"Users",
		healthsvc.HealthCheckResponse_NOT_SERVING,
	)
	res, err = stream.Recv()
	checkError("Streaming", err)

	if res.Status != healthsvc.HealthCheckResponse_NOT_SERVING {
		t.Errorf("Expected SERVING but got: %s\n", utils.GetJsonStringUnsafe(res.Status))
	}
	printResponse(t, res)
}
