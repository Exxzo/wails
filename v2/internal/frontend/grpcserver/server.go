package grpcserver

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/wailsapp/wails/v2/internal/binding"
	"github.com/wailsapp/wails/v2/internal/frontend/dispatcher"
	"github.com/wailsapp/wails/v2/internal/logger"
	wailspb "github.com/wailsapp/wails/v2/internal/proto"
	"google.golang.org/grpc"
)

type Server struct {
	wailspb.UnimplementedWailsServiceServer
	disp      *dispatcher.Dispatcher
	bindingDB *binding.DB
	logger    *logger.Logger
}

func New(d *dispatcher.Dispatcher, db *binding.DB, l *logger.Logger) *Server {
	return &Server{disp: d, bindingDB: db, logger: l}
}

func (s *Server) Call(ctx context.Context, req *wailspb.CallRequest) (*wailspb.CallResponse, error) {
	name := req.GetMethod()
	// Lookup bound method
	method := s.bindingDB.GetMethod(name)
	if method == nil {
		return &wailspb.CallResponse{Error: "method '" + name + "' not registered"}, nil
	}

	// Parse arguments (JSON array)
	var rawArgs []json.RawMessage
	if aj := req.GetArgsJson(); aj != "" {
		if err := json.Unmarshal([]byte(aj), &rawArgs); err != nil {
			return &wailspb.CallResponse{Error: "error parsing arguments: " + err.Error()}, nil
		}
	}
	args, err := method.ParseArgs(rawArgs)
	if err != nil {
		return &wailspb.CallResponse{Error: "error parsing arguments: " + err.Error()}, nil
	}

	// Call method
	result, err := method.Call(args)
	if err != nil {
		return &wailspb.CallResponse{Error: err.Error()}, nil
	}

	// JSON encode result for compatibility
	resultJSON, err := json.Marshal(result)
	if err != nil {
		return &wailspb.CallResponse{Error: err.Error()}, nil
	}
	return &wailspb.CallResponse{ResultJson: string(resultJSON)}, nil
}

func (s *Server) StreamEvents(req *wailspb.EventSubscribe, stream wailspb.WailsService_StreamEventsServer) error {
	// TODO: wire event handler forwarding
	return nil
}

func Serve(addr string, d *dispatcher.Dispatcher, db *binding.DB, l *logger.Logger) (stop func() error, err error) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	grpcServer := grpc.NewServer()
	wailspb.RegisterWailsServiceServer(grpcServer, New(d, db, l))
	go func() { _ = grpcServer.Serve(lis) }()
	return func() error {
		grpcServer.GracefulStop()
		return lis.Close()
	}, nil
}

// Optional health endpoint (useful during orchestration)
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}
