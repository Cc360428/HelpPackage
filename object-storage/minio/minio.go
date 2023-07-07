/**
 * @Author: Cc
 * @Description: minio 对象存储
 * @File: minio
 * @Version: 1.0.0
 * @Date: 2023/7/7 16:45
 * @Software : GoLand
 */

// https://min.io/docs/minio/linux/developers/go/API.html#MakeBucket

package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Endpoint        string
	AccessKeyId     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
}

type Client struct {
	minioClient *minio.Client
	Bucket      string
}

func NewClient(config Config) (*Client, error) {
	client, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyId, config.SecretAccessKey, ""),
		Secure: config.UseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &Client{
		minioClient: client,
		Bucket:      config.BucketName,
	}, err
}

func (c *Client) MakeBucket() {
	err := c.minioClient.MakeBucket(context.Background(), c.Bucket, minio.MakeBucketOptions{Region: "jiangxi", ObjectLocking: true})
	if err != nil {
		fmt.Println(err)
	}
}

func (c *Client) RemoveBucket() {
	c.minioClient.RemoveBucket(context.Background(), c.Bucket)
}
