package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"sync/atomic"

	library "./_proto"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
)

var (
	PORT = 9090

	// we use a map as an in-memory store for the application, it is protected by a RWMutes to allow
	// safe concurrent access.
	USER_LOCK = &sync.RWMutex{}

	// atomicly load and increment USER_INDEX to serve as a unique key
	USER_INDEX int64 = 3
	USERS            = map[int64]*library.User{
		1: &library.User{
			Id:        1,
			Email:     "user1@example.com",
			FirstName: "first1",
			LastName:  "last1",
		},
		2: &library.User{
			Id:        2,
			Email:     "user2@example.com",
			FirstName: "first2",
			LastName:  "last2",
		},
		3: &library.User{
			Id:        3,
			Email:     "user3@example.com",
			FirstName: "first3",
			LastName:  "last3",
		},
	}
)

func main() {
	grpcServer := grpc.NewServer()
	library.RegisterUserServiceServer(grpcServer, &userService{})
	grpclog.SetLogger(log.New(os.Stdout, "GRPC:", log.LstdFlags))

	wrappedServer := grpcweb.WrapServer(grpcServer)

	handler := func(res http.ResponseWriter, req *http.Request) {
		wrappedServer.ServeHTTP(res, req)
	}

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", PORT),
		Handler: http.HandlerFunc(handler),
	}

	grpclog.Println("Starting server...")
	log.Fatalln(httpServer.ListenAndServe())
}

type userService struct{}

func (ns *userService) AddUser(ctx context.Context, req *library.AddUserRequest) (*library.User, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.Lock()
	defer USER_LOCK.Unlock()

	atomic.AddInt64(&USER_INDEX, 1)
	id := atomic.LoadInt64(&USER_INDEX)

	user := &library.User{
		Id:        id,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	USERS[user.Id] = user
	return user, nil
}

func (ns *userService) GetUsers(ctx context.Context, req *library.GetUsersRequest) (*library.GetUsersResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.RLock()
	defer USER_LOCK.RUnlock()

	users := make([]*library.User, len(USERS))

	idx := 0
	for _, v := range USERS {
		users[idx] = v
		idx++
	}

	resp := &library.GetUsersResponse{Users: users}
	return resp, nil
}

func (ns *userService) DeleteUser(ctx context.Context, req *library.DeleteUserRequest) (*library.User, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.Lock()
	defer USER_LOCK.Unlock()

	if oldUser, exists := USERS[req.Id]; exists {
		user := &library.User{
			Id:        oldUser.Id,
			Email:     oldUser.Email,
			FirstName: oldUser.FirstName,
			LastName:  oldUser.LastName,
		}

		delete(USERS, req.Id)
		return user, nil
	}

	return nil, grpc.Errorf(codes.NotFound, "User could not be found.")
}
