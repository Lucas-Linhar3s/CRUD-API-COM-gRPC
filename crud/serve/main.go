package main

import (
	pb "apiGRPC/proto/gen"
	"context"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"strconv"
)

var Users []*pb.UserInfo

type userServer struct {
	pb.UnimplementedUserServer
}

func main() {
	initUsers()
	lis, err := net.Listen("tcp", ":50555")
	if err != nil {
		log.Fatalf("Deu Erro %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServer(s, &userServer{})
	log.Println("SERVE RODANDO NA PORTA :50555")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Deu Erro %v", err)
	}

}

func initUsers() {
	user1 := &pb.UserInfo{Id: "1", Nome: "Lucas", Sobrenome: "Linhares", Age: 21}
	user2 := &pb.UserInfo{Id: "2", Nome: "Lucas", Sobrenome: "Linhares", Age: 22}
	user3 := &pb.UserInfo{Id: "3", Nome: "Lucas", Sobrenome: "Linhares", Age: 23}

	Users = append(Users, user1)
	Users = append(Users, user2)
	Users = append(Users, user3)
}

func (s *userServer) GetUsers(in *pb.Empty, steam pb.User_GetUsersServer) error {
	log.Printf("MENSAGEM RECEBIDA %v", in)
	for _, user := range Users {
		if err := steam.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func (s *userServer) GetUser(ctx context.Context, in *pb.Id) (*pb.UserInfo, error) {
	log.Printf("MENSAGEM RECEBIDA %v", in)
	res := &pb.UserInfo{}
	for _, user := range Users {
		if user.GetId() == in.GetValue() {
			res = user
			break
		}
	}
	return res, nil
}

func (s *userServer) CreateUsers(ctx context.Context, in *pb.UserInfo) (*pb.Id, error) {
	log.Printf("MENSAGEM RECEBIDA %v", in)
	res := pb.Id{}
	res.Value = strconv.Itoa(rand.Intn(100000000))
	in.Id = res.GetValue()
	Users = append(Users, in)
	return &res, nil
}

func (s *userServer) UpdateUsers(ctx context.Context, in *pb.UserInfo) (*pb.Status, error) {
	log.Printf("MENSAGEM RECEBIDA %v", in)
	res := pb.Status{}
	for index, User := range Users {
		if User.GetId() == in.GetId() {
			Users = append(Users[:index], Users[index+1:]...)
			in.Id = User.GetId()
			Users = append(Users, in)
			res.Value = 1
			break
		}
	}
	return &res, nil
}

func (s *userServer) DeleteUsers(ctx context.Context, in *pb.Id) (*pb.Status, error) {
	log.Printf("MENSAGEM RECEBIDA %v", in)
	res := pb.Status{}
	for index, User := range Users {
		if User.GetId() == in.GetValue() {
			Users = append(Users[:index], Users[index+1:]...)
			res.Value = 1
			break
		}
	}
	return &res, nil
}
