package main

import (
	"context"
	"fmt"
	pb "github.com/pingjing0628/grpcPractice/proto"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	addr := "127.0.0.1:9999"

	// Dial 負責建立起 client and server 間的 gRPC channel 以進行溝通
	// WithInsecure 設定使用不安全的連線設定，為求簡便，否則應使用安全連線(tls/ssl)
	// WithBlock 讓 client 能在連線到 server 前先 block 住
	connection, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("can not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	// 建立 client
	client := pb.NewHelloServiceClient(connection)

	// WithTimeout 能在呼叫 sayHello 時，若超過 1 秒就是為超時
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	// response 結構有定義 GetReply()，能取得來自 gRPC server 的回應
	r, err := client.SayHello(ctx, &pb.HelloRequest{Greeting: "Moto"})

	if err != nil {
		log.Fatalf("could not get nonce: %v", err)
	}

	fmt.Println("Response:", r.GetReply())
}
