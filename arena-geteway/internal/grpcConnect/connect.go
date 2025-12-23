package grpc_connection

import (
	"context"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
)

func Connect(serviceName string) *grpc.ClientConn {
	envKey := serviceName + "_URL" 
	url := os.Getenv(envKey)
	if url == "" {
		log.Fatalf(" Env %s topilmadi", envKey)
	}

	var conn *grpc.ClientConn
	var err error

	
	for i := 1; i <= 3; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		conn, err = grpc.DialContext(ctx, url, grpc.WithInsecure(), grpc.WithBlock())
		if err == nil {
			log.Printf("GRPC connected to %s (%s)", serviceName, url)
			return conn
		}

		log.Printf(" GRPC connect attempt %d failed: %v", i, err)
		time.Sleep(1 * time.Second) 
	}

	log.Fatalf(" GRPC connection to %s failed after 3 attempts", serviceName)
	return nil
}
