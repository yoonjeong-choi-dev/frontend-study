package main

import (
	"google.golang.org/grpc"
	healthz "google.golang.org/grpc/health"
	healthsvc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/test/bufconn"
	"log"
	svc "robust-app/service"
	"robust-app/utils"
	"testing"
)

func startTestServer() (*bufconn.Listener, *healthz.Server) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	h := healthz.NewServer()

	registerService(s, h)
	updateHealthCheckService(
		h,
		svc.Users_ServiceDesc.ServiceName,
		healthsvc.HealthCheckResponse_SERVING,
	)

	go func() {
		defer s.GracefulStop()
		log.Fatal(s.Serve(l))
	}()
	return l, h
}

func checkError(errorType string, err error) {
	if err != nil {
		log.Fatalf("Error for %s: %s\n", errorType, utils.GetJsonStringUnsafe(err))
	}
}

func checkIntTypeResponse(t *testing.T,
	keyName string, expected, got int) {
	if expected != got {
		t.Errorf("Expetecd %s to be %d, Got: %d\n", keyName, expected, got)
	}
}

func checkStringTypeResponse(t *testing.T,
	keyName, expected, got string) {
	if expected != got {
		t.Errorf("Expetecd %s to be %s, Got: %s\n", keyName, expected, got)
	}
}

func printResponse(t *testing.T, res interface{}) {
	t.Logf("Response: %s\n", utils.GetJsonStringUnsafe(res))
}
