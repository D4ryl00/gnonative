// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc.proto

package rpcconnect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	rpc "github.com/gnolang/gnomobile/service/rpc"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion0_1_0

const (
	// GnomobileServiceName is the fully-qualified name of the GnomobileService service.
	GnomobileServiceName = "land.gno.gnomobile.v1.GnomobileService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GnomobileServiceSetRemoteProcedure is the fully-qualified name of the GnomobileService's
	// SetRemote RPC.
	GnomobileServiceSetRemoteProcedure = "/land.gno.gnomobile.v1.GnomobileService/SetRemote"
	// GnomobileServiceSetChainIDProcedure is the fully-qualified name of the GnomobileService's
	// SetChainID RPC.
	GnomobileServiceSetChainIDProcedure = "/land.gno.gnomobile.v1.GnomobileService/SetChainID"
	// GnomobileServiceSetPasswordProcedure is the fully-qualified name of the GnomobileService's
	// SetPassword RPC.
	GnomobileServiceSetPasswordProcedure = "/land.gno.gnomobile.v1.GnomobileService/SetPassword"
	// GnomobileServiceGenerateRecoveryPhraseProcedure is the fully-qualified name of the
	// GnomobileService's GenerateRecoveryPhrase RPC.
	GnomobileServiceGenerateRecoveryPhraseProcedure = "/land.gno.gnomobile.v1.GnomobileService/GenerateRecoveryPhrase"
	// GnomobileServiceListKeyInfoProcedure is the fully-qualified name of the GnomobileService's
	// ListKeyInfo RPC.
	GnomobileServiceListKeyInfoProcedure = "/land.gno.gnomobile.v1.GnomobileService/ListKeyInfo"
	// GnomobileServiceCreateAccountProcedure is the fully-qualified name of the GnomobileService's
	// CreateAccount RPC.
	GnomobileServiceCreateAccountProcedure = "/land.gno.gnomobile.v1.GnomobileService/CreateAccount"
	// GnomobileServiceSelectAccountProcedure is the fully-qualified name of the GnomobileService's
	// SelectAccount RPC.
	GnomobileServiceSelectAccountProcedure = "/land.gno.gnomobile.v1.GnomobileService/SelectAccount"
	// GnomobileServiceGetActiveAccountProcedure is the fully-qualified name of the GnomobileService's
	// GetActiveAccount RPC.
	GnomobileServiceGetActiveAccountProcedure = "/land.gno.gnomobile.v1.GnomobileService/GetActiveAccount"
	// GnomobileServiceQueryProcedure is the fully-qualified name of the GnomobileService's Query RPC.
	GnomobileServiceQueryProcedure = "/land.gno.gnomobile.v1.GnomobileService/Query"
	// GnomobileServiceCallProcedure is the fully-qualified name of the GnomobileService's Call RPC.
	GnomobileServiceCallProcedure = "/land.gno.gnomobile.v1.GnomobileService/Call"
	// GnomobileServiceHelloProcedure is the fully-qualified name of the GnomobileService's Hello RPC.
	GnomobileServiceHelloProcedure = "/land.gno.gnomobile.v1.GnomobileService/Hello"
)

// GnomobileServiceClient is a client for the land.gno.gnomobile.v1.GnomobileService service.
type GnomobileServiceClient interface {
	// Set the connection addresse for the remote node. If you don't call this,
	// the default is "127.0.0.1:26657"
	SetRemote(context.Context, *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error)
	// Set the chain ID for the remote node. If you don't call this, the default
	// is "dev"
	SetChainID(context.Context, *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error)
	// Set the password for the account in the keybase, used for later operations
	SetPassword(context.Context, *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error)
	// Generate a recovery phrase of BIP39 mnemonic words using entropy from the
	// crypto library random number generator. This can be used as the mnemonic in
	// CreateAccount.
	GenerateRecoveryPhrase(context.Context, *connect.Request[rpc.GenerateRecoveryPhraseRequest]) (*connect.Response[rpc.GenerateRecoveryPhraseResponse], error)
	// Get the keys informations in the keybase
	ListKeyInfo(context.Context, *connect.Request[rpc.ListKeyInfoRequest]) (*connect.Response[rpc.ListKeyInfoResponse], error)
	// Create a new account the keybase using the name an password specified by
	// SetAccount
	CreateAccount(context.Context, *connect.Request[rpc.CreateAccountRequest]) (*connect.Response[rpc.CreateAccountResponse], error)
	// SelectAccount selects the active account to use for later operations
	SelectAccount(context.Context, *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error)
	// GetActiveAccount gets the active account which was set by SelectAccount.
	// If there is no active account, then return ErrNoActiveAccount.
	// (To check if there is an active account, use ListKeyInfo and check the
	// length of the result.)
	GetActiveAccount(context.Context, *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error)
	// Make an ABCI query to the remote node.
	Query(context.Context, *connect.Request[rpc.QueryRequest]) (*connect.Response[rpc.QueryResponse], error)
	// Call a specific realm function.
	Call(context.Context, *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error)
	// Hello is for debug purposes
	Hello(context.Context, *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error)
}

// NewGnomobileServiceClient constructs a client for the land.gno.gnomobile.v1.GnomobileService
// service. By default, it uses the Connect protocol with the binary Protobuf Codec, asks for
// gzipped responses, and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply
// the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGnomobileServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) GnomobileServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &gnomobileServiceClient{
		setRemote: connect.NewClient[rpc.SetRemoteRequest, rpc.SetRemoteResponse](
			httpClient,
			baseURL+GnomobileServiceSetRemoteProcedure,
			opts...,
		),
		setChainID: connect.NewClient[rpc.SetChainIDRequest, rpc.SetChainIDResponse](
			httpClient,
			baseURL+GnomobileServiceSetChainIDProcedure,
			opts...,
		),
		setPassword: connect.NewClient[rpc.SetPasswordRequest, rpc.SetPasswordResponse](
			httpClient,
			baseURL+GnomobileServiceSetPasswordProcedure,
			opts...,
		),
		generateRecoveryPhrase: connect.NewClient[rpc.GenerateRecoveryPhraseRequest, rpc.GenerateRecoveryPhraseResponse](
			httpClient,
			baseURL+GnomobileServiceGenerateRecoveryPhraseProcedure,
			opts...,
		),
		listKeyInfo: connect.NewClient[rpc.ListKeyInfoRequest, rpc.ListKeyInfoResponse](
			httpClient,
			baseURL+GnomobileServiceListKeyInfoProcedure,
			opts...,
		),
		createAccount: connect.NewClient[rpc.CreateAccountRequest, rpc.CreateAccountResponse](
			httpClient,
			baseURL+GnomobileServiceCreateAccountProcedure,
			opts...,
		),
		selectAccount: connect.NewClient[rpc.SelectAccountRequest, rpc.SelectAccountResponse](
			httpClient,
			baseURL+GnomobileServiceSelectAccountProcedure,
			opts...,
		),
		getActiveAccount: connect.NewClient[rpc.GetActiveAccountRequest, rpc.GetActiveAccountResponse](
			httpClient,
			baseURL+GnomobileServiceGetActiveAccountProcedure,
			opts...,
		),
		query: connect.NewClient[rpc.QueryRequest, rpc.QueryResponse](
			httpClient,
			baseURL+GnomobileServiceQueryProcedure,
			opts...,
		),
		call: connect.NewClient[rpc.CallRequest, rpc.CallResponse](
			httpClient,
			baseURL+GnomobileServiceCallProcedure,
			opts...,
		),
		hello: connect.NewClient[rpc.HelloRequest, rpc.HelloResponse](
			httpClient,
			baseURL+GnomobileServiceHelloProcedure,
			opts...,
		),
	}
}

// gnomobileServiceClient implements GnomobileServiceClient.
type gnomobileServiceClient struct {
	setRemote              *connect.Client[rpc.SetRemoteRequest, rpc.SetRemoteResponse]
	setChainID             *connect.Client[rpc.SetChainIDRequest, rpc.SetChainIDResponse]
	setPassword            *connect.Client[rpc.SetPasswordRequest, rpc.SetPasswordResponse]
	generateRecoveryPhrase *connect.Client[rpc.GenerateRecoveryPhraseRequest, rpc.GenerateRecoveryPhraseResponse]
	listKeyInfo            *connect.Client[rpc.ListKeyInfoRequest, rpc.ListKeyInfoResponse]
	createAccount          *connect.Client[rpc.CreateAccountRequest, rpc.CreateAccountResponse]
	selectAccount          *connect.Client[rpc.SelectAccountRequest, rpc.SelectAccountResponse]
	getActiveAccount       *connect.Client[rpc.GetActiveAccountRequest, rpc.GetActiveAccountResponse]
	query                  *connect.Client[rpc.QueryRequest, rpc.QueryResponse]
	call                   *connect.Client[rpc.CallRequest, rpc.CallResponse]
	hello                  *connect.Client[rpc.HelloRequest, rpc.HelloResponse]
}

// SetRemote calls land.gno.gnomobile.v1.GnomobileService.SetRemote.
func (c *gnomobileServiceClient) SetRemote(ctx context.Context, req *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error) {
	return c.setRemote.CallUnary(ctx, req)
}

// SetChainID calls land.gno.gnomobile.v1.GnomobileService.SetChainID.
func (c *gnomobileServiceClient) SetChainID(ctx context.Context, req *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error) {
	return c.setChainID.CallUnary(ctx, req)
}

// SetPassword calls land.gno.gnomobile.v1.GnomobileService.SetPassword.
func (c *gnomobileServiceClient) SetPassword(ctx context.Context, req *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error) {
	return c.setPassword.CallUnary(ctx, req)
}

// GenerateRecoveryPhrase calls land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase.
func (c *gnomobileServiceClient) GenerateRecoveryPhrase(ctx context.Context, req *connect.Request[rpc.GenerateRecoveryPhraseRequest]) (*connect.Response[rpc.GenerateRecoveryPhraseResponse], error) {
	return c.generateRecoveryPhrase.CallUnary(ctx, req)
}

// ListKeyInfo calls land.gno.gnomobile.v1.GnomobileService.ListKeyInfo.
func (c *gnomobileServiceClient) ListKeyInfo(ctx context.Context, req *connect.Request[rpc.ListKeyInfoRequest]) (*connect.Response[rpc.ListKeyInfoResponse], error) {
	return c.listKeyInfo.CallUnary(ctx, req)
}

// CreateAccount calls land.gno.gnomobile.v1.GnomobileService.CreateAccount.
func (c *gnomobileServiceClient) CreateAccount(ctx context.Context, req *connect.Request[rpc.CreateAccountRequest]) (*connect.Response[rpc.CreateAccountResponse], error) {
	return c.createAccount.CallUnary(ctx, req)
}

// SelectAccount calls land.gno.gnomobile.v1.GnomobileService.SelectAccount.
func (c *gnomobileServiceClient) SelectAccount(ctx context.Context, req *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error) {
	return c.selectAccount.CallUnary(ctx, req)
}

// GetActiveAccount calls land.gno.gnomobile.v1.GnomobileService.GetActiveAccount.
func (c *gnomobileServiceClient) GetActiveAccount(ctx context.Context, req *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error) {
	return c.getActiveAccount.CallUnary(ctx, req)
}

// Query calls land.gno.gnomobile.v1.GnomobileService.Query.
func (c *gnomobileServiceClient) Query(ctx context.Context, req *connect.Request[rpc.QueryRequest]) (*connect.Response[rpc.QueryResponse], error) {
	return c.query.CallUnary(ctx, req)
}

// Call calls land.gno.gnomobile.v1.GnomobileService.Call.
func (c *gnomobileServiceClient) Call(ctx context.Context, req *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error) {
	return c.call.CallUnary(ctx, req)
}

// Hello calls land.gno.gnomobile.v1.GnomobileService.Hello.
func (c *gnomobileServiceClient) Hello(ctx context.Context, req *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error) {
	return c.hello.CallUnary(ctx, req)
}

// GnomobileServiceHandler is an implementation of the land.gno.gnomobile.v1.GnomobileService
// service.
type GnomobileServiceHandler interface {
	// Set the connection addresse for the remote node. If you don't call this,
	// the default is "127.0.0.1:26657"
	SetRemote(context.Context, *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error)
	// Set the chain ID for the remote node. If you don't call this, the default
	// is "dev"
	SetChainID(context.Context, *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error)
	// Set the password for the account in the keybase, used for later operations
	SetPassword(context.Context, *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error)
	// Generate a recovery phrase of BIP39 mnemonic words using entropy from the
	// crypto library random number generator. This can be used as the mnemonic in
	// CreateAccount.
	GenerateRecoveryPhrase(context.Context, *connect.Request[rpc.GenerateRecoveryPhraseRequest]) (*connect.Response[rpc.GenerateRecoveryPhraseResponse], error)
	// Get the keys informations in the keybase
	ListKeyInfo(context.Context, *connect.Request[rpc.ListKeyInfoRequest]) (*connect.Response[rpc.ListKeyInfoResponse], error)
	// Create a new account the keybase using the name an password specified by
	// SetAccount
	CreateAccount(context.Context, *connect.Request[rpc.CreateAccountRequest]) (*connect.Response[rpc.CreateAccountResponse], error)
	// SelectAccount selects the active account to use for later operations
	SelectAccount(context.Context, *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error)
	// GetActiveAccount gets the active account which was set by SelectAccount.
	// If there is no active account, then return ErrNoActiveAccount.
	// (To check if there is an active account, use ListKeyInfo and check the
	// length of the result.)
	GetActiveAccount(context.Context, *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error)
	// Make an ABCI query to the remote node.
	Query(context.Context, *connect.Request[rpc.QueryRequest]) (*connect.Response[rpc.QueryResponse], error)
	// Call a specific realm function.
	Call(context.Context, *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error)
	// Hello is for debug purposes
	Hello(context.Context, *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error)
}

// NewGnomobileServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGnomobileServiceHandler(svc GnomobileServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	gnomobileServiceSetRemoteHandler := connect.NewUnaryHandler(
		GnomobileServiceSetRemoteProcedure,
		svc.SetRemote,
		opts...,
	)
	gnomobileServiceSetChainIDHandler := connect.NewUnaryHandler(
		GnomobileServiceSetChainIDProcedure,
		svc.SetChainID,
		opts...,
	)
	gnomobileServiceSetPasswordHandler := connect.NewUnaryHandler(
		GnomobileServiceSetPasswordProcedure,
		svc.SetPassword,
		opts...,
	)
	gnomobileServiceGenerateRecoveryPhraseHandler := connect.NewUnaryHandler(
		GnomobileServiceGenerateRecoveryPhraseProcedure,
		svc.GenerateRecoveryPhrase,
		opts...,
	)
	gnomobileServiceListKeyInfoHandler := connect.NewUnaryHandler(
		GnomobileServiceListKeyInfoProcedure,
		svc.ListKeyInfo,
		opts...,
	)
	gnomobileServiceCreateAccountHandler := connect.NewUnaryHandler(
		GnomobileServiceCreateAccountProcedure,
		svc.CreateAccount,
		opts...,
	)
	gnomobileServiceSelectAccountHandler := connect.NewUnaryHandler(
		GnomobileServiceSelectAccountProcedure,
		svc.SelectAccount,
		opts...,
	)
	gnomobileServiceGetActiveAccountHandler := connect.NewUnaryHandler(
		GnomobileServiceGetActiveAccountProcedure,
		svc.GetActiveAccount,
		opts...,
	)
	gnomobileServiceQueryHandler := connect.NewUnaryHandler(
		GnomobileServiceQueryProcedure,
		svc.Query,
		opts...,
	)
	gnomobileServiceCallHandler := connect.NewUnaryHandler(
		GnomobileServiceCallProcedure,
		svc.Call,
		opts...,
	)
	gnomobileServiceHelloHandler := connect.NewUnaryHandler(
		GnomobileServiceHelloProcedure,
		svc.Hello,
		opts...,
	)
	return "/land.gno.gnomobile.v1.GnomobileService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GnomobileServiceSetRemoteProcedure:
			gnomobileServiceSetRemoteHandler.ServeHTTP(w, r)
		case GnomobileServiceSetChainIDProcedure:
			gnomobileServiceSetChainIDHandler.ServeHTTP(w, r)
		case GnomobileServiceSetPasswordProcedure:
			gnomobileServiceSetPasswordHandler.ServeHTTP(w, r)
		case GnomobileServiceGenerateRecoveryPhraseProcedure:
			gnomobileServiceGenerateRecoveryPhraseHandler.ServeHTTP(w, r)
		case GnomobileServiceListKeyInfoProcedure:
			gnomobileServiceListKeyInfoHandler.ServeHTTP(w, r)
		case GnomobileServiceCreateAccountProcedure:
			gnomobileServiceCreateAccountHandler.ServeHTTP(w, r)
		case GnomobileServiceSelectAccountProcedure:
			gnomobileServiceSelectAccountHandler.ServeHTTP(w, r)
		case GnomobileServiceGetActiveAccountProcedure:
			gnomobileServiceGetActiveAccountHandler.ServeHTTP(w, r)
		case GnomobileServiceQueryProcedure:
			gnomobileServiceQueryHandler.ServeHTTP(w, r)
		case GnomobileServiceCallProcedure:
			gnomobileServiceCallHandler.ServeHTTP(w, r)
		case GnomobileServiceHelloProcedure:
			gnomobileServiceHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGnomobileServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGnomobileServiceHandler struct{}

func (UnimplementedGnomobileServiceHandler) SetRemote(context.Context, *connect.Request[rpc.SetRemoteRequest]) (*connect.Response[rpc.SetRemoteResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.SetRemote is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) SetChainID(context.Context, *connect.Request[rpc.SetChainIDRequest]) (*connect.Response[rpc.SetChainIDResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.SetChainID is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) SetPassword(context.Context, *connect.Request[rpc.SetPasswordRequest]) (*connect.Response[rpc.SetPasswordResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.SetPassword is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) GenerateRecoveryPhrase(context.Context, *connect.Request[rpc.GenerateRecoveryPhraseRequest]) (*connect.Response[rpc.GenerateRecoveryPhraseResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.GenerateRecoveryPhrase is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) ListKeyInfo(context.Context, *connect.Request[rpc.ListKeyInfoRequest]) (*connect.Response[rpc.ListKeyInfoResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.ListKeyInfo is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) CreateAccount(context.Context, *connect.Request[rpc.CreateAccountRequest]) (*connect.Response[rpc.CreateAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.CreateAccount is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) SelectAccount(context.Context, *connect.Request[rpc.SelectAccountRequest]) (*connect.Response[rpc.SelectAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.SelectAccount is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) GetActiveAccount(context.Context, *connect.Request[rpc.GetActiveAccountRequest]) (*connect.Response[rpc.GetActiveAccountResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.GetActiveAccount is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) Query(context.Context, *connect.Request[rpc.QueryRequest]) (*connect.Response[rpc.QueryResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.Query is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) Call(context.Context, *connect.Request[rpc.CallRequest]) (*connect.Response[rpc.CallResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.Call is not implemented"))
}

func (UnimplementedGnomobileServiceHandler) Hello(context.Context, *connect.Request[rpc.HelloRequest]) (*connect.Response[rpc.HelloResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("land.gno.gnomobile.v1.GnomobileService.Hello is not implemented"))
}
