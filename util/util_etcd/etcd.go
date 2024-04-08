package util_etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func Test() {
	for true {
		key := "mykey"

		cli, err := clientv3.New(clientv3.Config{
			Endpoints:   []string{"192.168.31.29:2379"},
			DialTimeout: 5 * time.Second,
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		defer cli.Close()

		go func() {
			for true {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second)
				_, err = cli.Put(ctx, key, fmt.Sprintf("second: %d", time.Now().Second()))
				cancel()
				if err != nil {
					fmt.Println(err)
					return
				}

				time.Sleep(time.Second)
			}
		}()

		ctx, _ := context.WithTimeout(context.Background(), 99999999*time.Second)
		log.Println("Watching ... key = ", key)
		watcher := cli.Watch(ctx, key, clientv3.WithPrefix())
		for {
			select {
			case resp := <-watcher:
				for _, event := range resp.Events {
					log.Printf("Event Type: %v, Key: %s, Value: %s\n", event.Type, event.Kv.Key, event.Kv.Value)
				}
			}
		}
		log.Println("end Watching ... key = ", key)

		time.Sleep(time.Second * 5)
	}
}
