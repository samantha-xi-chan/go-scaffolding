package util_etcd_v2

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var endpoints = []string{"192.168.31.8:2379"}

func ensureConnected(cli **clientv3.Client) {
	if *cli == nil || (*cli).ActiveConnection() == nil {
		var err error
		for {
			*cli, err = clientv3.New(clientv3.Config{
				Endpoints:   endpoints,
				DialTimeout: 5 * time.Second,
			})
			if err == nil && (*cli).ActiveConnection() != nil {
				break
			}
			log.Println("Connection to etcd failed, retrying...")
			time.Sleep(1 * time.Second) // 等待一秒后重试
		}
	}
}

func putKey(cli **clientv3.Client, key, value string) (e error) {
	ensureConnected(cli)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if _, err := (*cli).Put(ctx, key, value); err != nil {
		return fmt.Errorf("failed to put key: %s, value: %s, err: %w", key, value, err)
	} else {
		log.Printf("Put key: %s, value: %s successfully", key, value)
	}

	return nil
}

func getKey(cli **clientv3.Client, key string) (ee error) {
	ensureConnected(cli)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := (*cli).Get(ctx, key)
	if err != nil {
		return fmt.Errorf("failed to get key: %s, err: %w", key, err)
	}
	for _, ev := range resp.Kvs {
		log.Printf("Get key: %s, value: %s", string(ev.Key), string(ev.Value))
	}

	return nil
}

func Test() {
	var cli *clientv3.Client
	err := putKey(&cli, "testKey", "testValue")
	if err != nil {
		log.Printf("putKey err = %#v", err)
	}

	ee := getKey(&cli, "testKey")
	if ee != nil {
		log.Printf("getKey err = %#v", err)
	}

	if cli != nil {
		defer cli.Close()
	}
}
