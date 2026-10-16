package prisma

import (
	"fmt"
	"log"
	"time"

	prismadb "Momentum/prisma/db"
)

const maxConnectAttempts = 5

func Open() (*prismadb.PrismaClient, error) {
	client := prismadb.NewClient()

	var err error
	for attempt := 1; attempt <= maxConnectAttempts; attempt++ {
		err = client.Prisma.Connect()
		if err == nil {
			log.Println("Connected to Prisma")
			return client, nil
		}
		log.Printf("Prisma connect attempt %d/%d failed: %v", attempt, maxConnectAttempts, err)
		time.Sleep(time.Duration(attempt) * time.Second)
	}

	_ = client.Prisma.Disconnect()
	return nil, fmt.Errorf("connect to Prisma after retries: %w", err)
}
