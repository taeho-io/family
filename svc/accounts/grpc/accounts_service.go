package grpc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/svc/accounts/config"
	"github.com/taeho-io/family/svc/accounts/crypt"
	"github.com/taeho-io/family/svc/accounts/grpc/handlers"
	"github.com/taeho-io/family/svc/accounts/repos/account_email_repo"
	"github.com/taeho-io/family/svc/accounts/repos/account_repo"
	"github.com/taeho-io/family/svc/srv/aws"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
	srvGRPC "github.com/taeho-io/family/svc/srv/grpc"
)

type IFace interface {
	Config() config.IFace
	Crypt() crypt.IFace
	Dynamodb() dynamodb.IFace
	AccountTable() *account_repo.Table
	AccountEmailTable() *account_email_repo.Table
}

type Service struct {
	srvGRPC.IFace

	cfg               config.IFace
	crypt             crypt.IFace
	ddb               dynamodb.IFace
	accountTable      *account_repo.Table
	accountEmailTable *account_email_repo.Table
}

func New(cfg config.IFace) (*Service, error) {
	bcrypt := crypt.New()

	a, err := aws.New()
	if err != nil {
		return nil, err
	}
	ddb := dynamodb.New(a)

	return &Service{
		cfg:               cfg,
		crypt:             bcrypt,
		ddb:               ddb,
		accountTable:      account_repo.New(ddb, cfg),
		accountEmailTable: account_email_repo.New(ddb, cfg),
	}, nil
}

func NewMock() (*Service, error) {
	return New(config.NewMock())
}

func (s *Service) Config() config.IFace {
	return s.cfg
}

func (s *Service) Crypt() crypt.IFace {
	return s.crypt
}

func (s *Service) Dynamodb() dynamodb.IFace {
	return s.ddb
}

func (s *Service) AccountTable() *account_repo.Table {
	return s.accountTable
}

func (s *Service) AccountEmailTable() *account_email_repo.Table {
	return s.accountEmailTable
}

func (s *Service) RegisterService(srv *grpc.Server) {
	accounts.RegisterAccountsServiceServer(srv, s)
}

func (s *Service) Register(ctx context.Context, req *accounts.RegisterRequest) (*accounts.RegisterResponse, error) {
	return handlers.Register(s.AccountTable(), s.AccountEmailTable(), s.Crypt())(ctx, req)
}

func (s *Service) LogIn(ctx context.Context, req *accounts.LogInRequest) (*accounts.LogInResponse, error) {
	return handlers.LogIn(s.AccountTable(), s.AccountEmailTable(), s.Crypt())(ctx, req)
}