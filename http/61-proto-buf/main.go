package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-demo/database"
	"grpc-demo/models"
	userv1 "grpc-demo/proto"
	"log/slog"
	"net"
	"time"

	"os"

	"google.golang.org/grpc"
)

var (
	PORT, DB_URL string
)

func init() {
	PORT = os.Getenv("PORT")
	DB_URL = os.Getenv("DB_URL")
}

type UserServer struct {
	userv1.UnimplementedUserServiceServer
	IUserDB database.IUserDB
}

func NewUserServer(db database.IUserDB) *UserServer {
	return &UserServer{IUserDB: db}
}

func (us *UserServer) InsertUser(ctx context.Context, req *userv1.CreateUserRequest) (rep *userv1.CreateUserResponse, err error) {
	var user *models.User = new(models.User)
	user.Name = req.Name
	user.Email = req.Email
	user.Mobile = req.Mobile
	user.Status = "active"
	user.LastModified = time.Now().Unix()

	user, err = us.IUserDB.Create(user)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	rep = new(userv1.CreateUserResponse)

	repUser := new(userv1.User)

	repUser.Id = uint64(user.Id)
	repUser.Name = user.Name
	repUser.Email = user.Email
	rep.User = repUser
	return rep, nil
}

func main() {

	//args := os.Args // --port=8089 -port=8089 --port 8089 -port 8089

	if PORT == "" {
		flag.StringVar(&PORT, "port", "28090", "--port=8089 -port=8089 --port 8089 -port 8089")

	}
	if DB_URL == "" {
		flag.StringVar(&DB_URL, "db", `host=localhost user=postgres password=postgres dbname=demodb port=5432 sslmode=disable TimeZone=Asia/Shanghai`, "db=db string")
	}

	flag.Parse()

	db, err := database.GetConnection(DB_URL)
	if err != nil {
		panic(err.Error())
	}

	lis, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	IuserDB := database.NewUserDB(db)
	userServer := NewUserServer(IuserDB)

	s := grpc.NewServer()

	userv1.RegisterUserServiceServer(s, userServer)
	userv1.RegisterCustomerServer(s, customerServer)

	if err := s.Serve(lis); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
