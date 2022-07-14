package main

import (
	pb "apiGRPC/proto/gen"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:50555", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Deu Erro %v", err)
	}
	defer conn.Close()

	client := pb.NewUserClient(conn)

	initGetUsers(client)
	//runGetUser(client, "1")
	//runCreateUser(client, "Lucas", "Linhares", 21)
	//runUpdateUser(client, "1", "Jackosn", "Jackson", 10)
	//runDeleteUser(client, "3")
}

func initGetUsers(client pb.UserClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Empty{}
	stream, err := client.GetUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUsers(_) = _, %v", client, err)
	}
	for {
		linhas, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetUsers(_) = _, %v", client, err)
		}
		log.Printf("UsersInfo: %v", linhas)
	}

}

func runGetUser(client pb.UserClient, userid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: userid}
	res, err := client.GetUser(ctx, req)
	if err != nil {
		log.Fatalf("%v.GetUser(_) = _, %v", client, err)
	}
	log.Printf("UserInfo: %v", res)
}

func runCreateUser(client pb.UserClient, nome string, sobrenome string, age int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UserInfo{Nome: nome, Sobrenome: sobrenome, Age: age}
	res, err := client.CreateUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.CreateUser(_) = _, %v", client, err)
	}
	if res.GetValue() != "" {
		log.Printf("CreateUser Failed")
	}
}

func runUpdateUser(client pb.UserClient, userid string, nome string, sobrenome string, age int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.UserInfo{Id: userid, Nome: nome, Sobrenome: sobrenome, Age: age}
	res, err := client.UpdateUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.UpdateUser(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("UpdateUser Success")
	}
}

func runDeleteUser(client pb.UserClient, userid string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	req := &pb.Id{Value: userid}
	res, err := client.DeleteUsers(ctx, req)
	if err != nil {
		log.Fatalf("%v.DeleteUser(_) = _, %v", client, err)
	}
	if int(res.GetValue()) == 1 {
		log.Printf("DeleteUser Success")
	} else {
		log.Printf("DeleteUser Failed")
	}

}
