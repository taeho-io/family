package todo_groups_service

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sirupsen/logrus"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/grpc/dynamodb_service"
	"github.com/taeho-io/family/services/base/grpc/interceptors"
	"github.com/taeho-io/family/services/discovery/grpc/discovery_service"
	"github.com/taeho-io/family/services/todo_groups/config"
	"github.com/taeho-io/family/services/todo_groups/grpc/todo_groups_service/handlers"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_group_permits_repo"
	"github.com/taeho-io/family/services/todo_groups/repos/todo_groups_repo"
)

type IFace interface {
	dynamodb_service.IFace

	TodoGroupsTable() *todo_groups_repo.Table
	TodoGroupPermitsTable() *todo_group_permits_repo.Table
}

type Service struct {
	dynamodb_service.IFace

	todoGroupsTable       *todo_groups_repo.Table
	todoGroupPermitsTable *todo_group_permits_repo.Table
}

func New(cfg config.IFace) (IFace, error) {
	dynamodbSvc, err := dynamodb_service.New(cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		IFace:                 dynamodbSvc,
		todoGroupsTable:       todo_groups_repo.New(dynamodbSvc.Dynamodb(), cfg),
		todoGroupPermitsTable: todo_group_permits_repo.New(dynamodbSvc.Dynamodb(), cfg),
	}, nil
}

func NewMock() (IFace, error) {
	return New(config.NewMock())
}

func (s *Service) RegisterService(srv *grpc.Server) {
	todo_groups.RegisterTodoGroupsServiceServer(srv, s)
}

func (s *Service) TodoGroupsTable() *todo_groups_repo.Table {
	return s.todoGroupsTable
}

func (s *Service) TodoGroupPermitsTable() *todo_group_permits_repo.Table {
	return s.todoGroupPermitsTable
}

func (s *Service) CreateTodoGroup(
	ctx context.Context,
	req *todo_groups.CreateTodoGroupRequest,
) (*todo_groups.CreateTodoGroupResponse, error) {
	logger := s.Logger().WithFields(logrus.Fields{
		"methodName": "CreateTodoGroup",
		"req":        req,
	})

	return handlers.CreateTodoGroup(logger, s.TodoGroupsTable(), s.TodoGroupPermitsTable())(ctx, req)
}

func (s *Service) GetTodoGroup(
	ctx context.Context,
	req *todo_groups.GetTodoGroupRequest,
) (*todo_groups.GetTodoGroupResponse, error) {
	logger := s.Logger().WithFields(logrus.Fields{
		"methodName": "GetTodoGroup",
		"req":        req,
	})

	return handlers.GetTodoGroup(logger, s.TodoGroupsTable(), s.TodoGroupPermitsTable())(ctx, req)
}

func (s *Service) ListTodoGroups(
	ctx context.Context,
	req *todo_groups.ListTodoGroupsRequest,
) (*todo_groups.ListTodoGroupsResponse, error) {
	return handlers.ListTodoGroups(s.TodoGroupsTable(), s.TodoGroupPermitsTable())(ctx, req)
}

func (s *Service) UpdateTodoGroup(
	ctx context.Context,
	req *todo_groups.UpdateTodoGroupRequest,
) (*todo_groups.UpdateTodoGroupResponse, error) {
	return handlers.UpdateTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func (s *Service) DeleteTodoGroup(
	ctx context.Context,
	req *todo_groups.DeleteTodoGroupRequest,
) (*todo_groups.DeleteTodoGroupResponse, error) {
	return handlers.DeleteTodoGroup(s.TodoGroupsTable())(ctx, req)
}

func Serve() error {
	addr := discovery_service.ServiceAddrMap[discovery.Service_TODOGROUPS]
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	var whiteListPrefixes []string

	svr := grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			interceptors.NewRequestIdIfNotExistsUnaryServerInterceptor,
			interceptors.ForwardRequestIdLogFieldUnaryServerInterceptor,
			interceptors.LogrusUnaryServerInterceptor,
			interceptors.AuthWhileListUnaryServerInterceptor(whiteListPrefixes),
			interceptors.AuthUnaryServerInterceptor,
		),
	)

	svc, err := New(config.New(config.NewSettings()))
	if err != nil {
		return err
	}

	svc.RegisterService(svr)
	reflection.Register(svr)
	return svr.Serve(lis)
}
