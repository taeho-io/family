package main

import (
	"flag"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/xissy/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
	"github.com/taeho-io/family/idl/generated/go/pb/family/auth"
	"github.com/taeho-io/family/idl/generated/go/pb/family/discovery"
	"github.com/taeho-io/family/idl/generated/go/pb/family/notes"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
	accountsService "github.com/taeho-io/family/services/accounts"
	authService "github.com/taeho-io/family/services/auth"
	discoveryService "github.com/taeho-io/family/services/discovery"
	notesService "github.com/taeho-io/family/services/notes"
	todogroupsService "github.com/taeho-io/family/services/todogroups"
	todosService "github.com/taeho-io/family/services/todos"
)

var (
	gatewayAddr        = ":" + os.Getenv("PORT")
	requestIdHeaderKey = "x-request-id"

	authServerEndpoint = flag.String(
		"auth_server_endpoint",
		discoveryService.ServiceAddrMap[discovery.Service_AUTH],
		"endpoint of AuthServer",
	)
	accountsServerEndpoint = flag.String(
		"accounts_server_endpoint",
		discoveryService.ServiceAddrMap[discovery.Service_ACCOUNTS],
		"endpoint of AccountsServer",
	)
	todoGroupsServerEndpoint = flag.String(
		"todogroups_server_endpoint",
		discoveryService.ServiceAddrMap[discovery.Service_TODOGROUPS],
		"endpoint of TodoGroupsServer",
	)
	todosServerEndpoint = flag.String(
		"todos_server_endpoint",
		discoveryService.ServiceAddrMap[discovery.Service_TODOS],
		"endpoint of TodosServer",
	)
	notesServerEndpoint = flag.String(
		"notes_server_endpoint",
		discoveryService.ServiceAddrMap[discovery.Service_NOTES],
		"endpoint of NotesServer",
	)
)

func requestIDMatcher(headerName string) (mdName string, ok bool) {
	lowerCasedHeaderName := strings.ToLower(headerName)
	if lowerCasedHeaderName == requestIdHeaderKey {
		return lowerCasedHeaderName, true
	}
	return "", false
}

func serveGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(requestIDMatcher))
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := auth.RegisterAuthServiceHandlerFromEndpoint(
		ctx,
		mux,
		*authServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := accounts.RegisterAccountsServiceHandlerFromEndpoint(
		ctx,
		mux,
		*accountsServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := todogroups.RegisterTodoGroupsServiceHandlerFromEndpoint(
		ctx,
		mux,
		*todoGroupsServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := todos.RegisterTodosServiceHandlerFromEndpoint(
		ctx,
		mux,
		*todosServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	if err := notes.RegisterNotesServiceHandlerFromEndpoint(
		ctx,
		mux,
		*notesServerEndpoint,
		opts,
	); err != nil {
		return err
	}

	return http.ListenAndServe(gatewayAddr, mux)
}

func startGrpcServices() error {
	errs := make(chan error, 1)

	type serveFunc func() error

	serveFuncs := []serveFunc{
		authService.Serve,
		accountsService.Serve,
		todogroupsService.Serve,
		todosService.Serve,
		notesService.Serve,
	}

	for _, serve := range serveFuncs {
		go func(serve serveFunc) {
			if err := serve(); err != nil {
				errs <- err
			}
		}(serve)
	}

	if err, open := <-errs; open {
		return err
	}

	return nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&logrus.ApexUpJSONFormatter{
		DisableTimestamp: true,
	})

	go func() {
		log.WithField("server_type", "grpc").Info("initializing")

		err := startGrpcServices()
		if err != nil {
			log.WithField("server_type", "grpc").WithError(err).Fatal("failed to listen")
			return
		}
	}()

	// sleep a second to make sure all servers are ready.
	// TODO: find a better way to optimize the waiting time.
	time.Sleep(time.Millisecond * 1500)
	log.WithField("server_type", "grpc_gw").Info("initializing")
	if err := serveGateway(); err != nil {
		log.WithField("server_type", "grpc_tw").WithError(err).Fatal("failed to listen")
		log.Fatalf("gateway: failed to listen: %x", err.Error())
	}
}
