package users_service

import (
	"sync"
	"sync/atomic"

	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var (
	// we use a map as an in-memory store for the application, it is protected by a RWMutes to allow
	// safe concurrent access.
	USER_LOCK = &sync.RWMutex{}

	// atomicly load and increment USER_INDEX to serve as a unique key
	USER_INDEX int64 = 3

	// initialize test data for User map
	USERS = map[int64]*User{
		1: &User{
			Id:        1,
			Email:     "user1@example.com",
			FirstName: "first1",
			LastName:  "last1",
		},
		2: &User{
			Id:        2,
			Email:     "user2@example.com",
			FirstName: "first2",
			LastName:  "last2",
		},
		3: &User{
			Id:        3,
			Email:     "user3@example.com",
			FirstName: "first3",
			LastName:  "last3",
		},
	}
)

type UserService struct{}

func (ns *UserService) AddUser(ctx context.Context, req *AddUserRequest) (*User, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.Lock()
	defer USER_LOCK.Unlock()

	atomic.AddInt64(&USER_INDEX, 1)
	id := atomic.LoadInt64(&USER_INDEX)

	user := &User{
		Id:        id,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	USERS[user.Id] = user
	return user, nil
}

func (ns *UserService) GetUsers(ctx context.Context, req *GetUsersRequest) (*GetUsersResponse, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.RLock()
	defer USER_LOCK.RUnlock()

	users := make([]*User, len(USERS))

	idx := 0
	for _, v := range USERS {
		users[idx] = v
		idx++
	}

	resp := &GetUsersResponse{Users: users}
	return resp, nil
}

func (ns *UserService) DeleteUser(ctx context.Context, req *DeleteUserRequest) (*User, error) {
	grpc.SendHeader(ctx, metadata.Pairs("Pre-Response-Metadata", "Is-sent-as-headers-unary"))

	USER_LOCK.Lock()
	defer USER_LOCK.Unlock()

	if oldUser, exists := USERS[req.Id]; exists {
		user := &User{
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
