package util_etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

/* 服务端部署
docker network create etcd-net
docker rm -f "etcd"
docker run -d --name etcd \
    --network etcd-net \
    -p 2379:2379 \
    --restart always \
    quay.io/coreos/etcd:v3.3.8 \
    etcd \
    --advertise-client-urls http://0.0.0.0:2379 \
    --listen-client-urls http://0.0.0.0:2379
*/

const (
	key    = "mykey"
	SERVER = "192.168.31.8:2379"
)

func WriteTest(ctx context.Context) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{SERVER},
		DialTimeout: 3 * time.Second,
	})
	if err != nil {
		return err
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 设置适当的超时时间
	defer cancel()

	for {
		if _, err := cli.Put(ctx, "ready_key", "ready"); err != nil {
			log.Println("etcd not ready, retrying...")
			time.Sleep(500 * time.Millisecond)
			continue
		}
		break
	}

	_, err = cli.Put(ctx, key, "my-value")
	if err != nil {
		return err
	}

	fmt.Printf("etcd write ok，key: %s, value: %s\n", key, "value")
	return nil
}
func StartReadTest(ctx context.Context) {
	var cli *clientv3.Client
	var watcher clientv3.WatchChan
	var err error

	for {
		// Establish the initial connection
		cli, err = clientv3.New(clientv3.Config{
			Endpoints:   []string{SERVER},
			DialTimeout: 3 * time.Second,
		})
		if err != nil {
			log.Println("Failed to connect to the server:", err)
			time.Sleep(5 * time.Second) // Wait before attempting to reconnect
			continue
		}

		// Create a new watcher
		watcher = cli.Watch(ctx, key, clientv3.WithPrefix())

		// Check for errors on watcher creation
		if watcher == nil {
			log.Println("Failed to create watcher.")
			cli.Close()
			time.Sleep(5 * time.Second) // Wait before attempting to reconnect
			continue
		}

		log.Println("Watching ... key =", key)

		// Loop to handle events
		for resp := range watcher {
			for _, event := range resp.Events {
				log.Printf("Event Type: %v, Key: %s, Value: %s\n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}

		// Watcher channel closed, attempt reconnection
		log.Println("Watcher channel closed. Attempting to reconnect...")
		cli.Close()
		time.Sleep(5 * time.Second) // Wait before attempting to reconnect
	}
}
