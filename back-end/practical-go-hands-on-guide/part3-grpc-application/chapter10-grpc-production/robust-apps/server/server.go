package main

import (
	"google.golang.org/grpc"
	healthz "google.golang.org/grpc/health"
	healthsvc "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"robust-app/server/interceptor"
	svc "robust-app/service"
	"robust-app/utils"
)

func createServerWithInterceptor() *grpc.Server {
	return grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.LoggingUnaryInterceptor,
			interceptor.DisconnectUnaryInterceptor,
			interceptor.TimeoutUnaryInterceptor,
			interceptor.PanicUnaryInterceptor,
		),
		grpc.ChainStreamInterceptor(
			interceptor.LoggingStreamInterceptor,
			interceptor.DisconnectStreamInterceptor,
			interceptor.TimeoutStreamInterceptor,
			interceptor.PanicStreamInterceptor,
		),
	)
}

func registerService(s *grpc.Server, h *healthz.Server) {
	svc.RegisterUsersServer(s, &userService{})
	healthsvc.RegisterHealthServer(s, h)
}

func startServer(s *grpc.Server, l net.Listener) error {
	return s.Serve(l)
}

func updateHealthCheckService(
	h *healthz.Server,
	service string,
	status healthsvc.HealthCheckResponse_ServingStatus) {
	h.SetServingStatus(service, status)
}

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error for TCP Listener: %s\n",
			utils.GetJsonStringUnsafe(err),
		)
	}

	s := createServerWithInterceptor()
	h := healthz.NewServer()
	registerService(s, h)
	updateHealthCheckService(
		h,
		svc.Users_ServiceDesc.ServiceName,
		healthsvc.HealthCheckResponse_SERVING,
	)

	log.Fatal(startServer(s, l))
}
