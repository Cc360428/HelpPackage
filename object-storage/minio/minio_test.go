/**
 * @Author: Cc
 * @Description: 描述
 * @File: minio_test.go
 * @Version: 1.0.0
 * @Date: 2023/7/7 16:53
 * @Software : GoLand
 */

package minio

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"testing"
)

func TestNewClient(t *testing.T) {
	var conf Config

	conf.UseSSL = false
	conf.Endpoint = "172.12.12.189:10030"
	conf.AccessKeyId = "mo9fzgBX1AYvsymRwypR"
	conf.SecretAccessKey = "niRQcMSX3tCCalfCJ2tBpvjM4cTJqyJGOXW7vWK9"
	conf.BucketName = "cctest"
	client, err := NewClient(conf)
	if err != nil {
		t.Error(err)
		return

	}
	client.MakeBucket()

	object, err := client.minioClient.PutObject(context.Background(), client.Bucket, "user1.json", bytes.NewReader([]byte(`{"Id":1}`)), 1, minio.PutObjectOptions{
		ContentType: "application/json",
	})
	if err != nil {
		return
	}
	t.Log(client, object.Key)
}
