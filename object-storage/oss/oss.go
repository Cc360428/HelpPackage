/**
 * @Author: Cc
 * @Description: aliYun OSS 对象存储 **如果没有后缀，二进制的话需要Suffix设置为空**
 * @File: oss
 * @Version: 1.0.0
 * @Date: 2023/7/6 15:56
 * @Software : GoLand
 */

package oss

import (
	"errors"
	"io"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Client struct {
	oSSClient    *oss.Client
	bucketClient *oss.Bucket
	prefix       string
	suffix       string
}

type Config struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	BucketName      string // 使用前必须先创建好Bucket
	Prefix          string // "2020/dev/" (为了多个有多个游戏区分开来)
	Suffix          string // ".json" 文件后缀
}

func NewOSSClient(config Config) (*Client, error) {
	if config.Prefix == "" {
		return nil, errors.New("prefix path is null")
	}

	client, err := oss.New(config.Endpoint, config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	bucketClient, err := client.Bucket(config.BucketName)
	if err != nil {
		return nil, err
	}

	return &Client{
		oSSClient:    client,
		bucketClient: bucketClient,
		prefix:       config.Prefix,
		suffix:       config.Suffix,
	}, nil
}

type Object struct {
	Key   string
	Value string
}

func (c *Client) key(key string) string {
	return c.prefix + key + c.suffix
}

// Put 上传 & 更新
func (c *Client) Put(object Object) error {
	if object.Key == "" || object.Value == "" {
		return errors.New("put object key or value is null")
	}
	return c.bucketClient.PutObject(c.key(object.Key), strings.NewReader(object.Value))
}

// ClearObject 删除权限没有只能对文件晴空
func (c *Client) ClearObject(key string) error {
	if key == "" {
		return errors.New("put object key or value is null")
	}
	return c.bucketClient.PutObject(c.key(key), strings.NewReader(""))
}

func (c *Client) IsObjectExist(key string) (bool, error) {
	return c.bucketClient.IsObjectExist(c.key(key))
}

func (c *Client) Get(key string) ([]byte, error) {

	object, err := c.bucketClient.GetObject(c.key(key))
	if err != nil {
		return nil, err
	}

	value, err := io.ReadAll(object)
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c *Client) GetToStruct(key string) (io.ReadCloser, error) {
	return c.bucketClient.GetObject(c.key(key))
}

// GetObjectList 当前 bucket 前缀 下面所有的key　max 100
// marker 标记，从那个开始 bucket object 有 1，2，3，4，5，6
//
//	marker 传了 2
//
// maxKeys 最大返回数量
//
//	maxKeys 传了 2
//
// return object [3, 4]
func (c *Client) GetObjectList(marker string, maxKeys int) ([]string, error) {

	if maxKeys == 0 {
		maxKeys = 10 // oss api default maxKeys value is 100
	}

	if marker != "" {
		marker = c.key(marker)
	}

	objects, err := c.bucketClient.ListObjects(oss.Prefix(c.prefix), oss.Marker(marker), oss.MaxKeys(maxKeys))
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(objects.Objects))
	for _, i2 := range objects.Objects {
		keys = append(keys, i2.Key)
	}
	return keys, nil
}

// GetJsonObjectByLines 查询
// https://github.com/aliyun/aliyun-oss-go-sdk/blob/master/oss/select_json_object_test.go#L97
func (c *Client) GetJsonObjectByLines(key, expressionSQL string) ([]byte, error) {
	jsonMeta := oss.JsonMetaRequest{
		InputSerialization: oss.InputSerialization{
			JSON: oss.JSON{
				JSONType: "LINES",
			},
		},
	}

	_, err := c.bucketClient.CreateSelectJsonObjectMeta(c.key(key), jsonMeta)
	if err != nil {
		return nil, err
	}

	var sqlSearch oss.SelectRequest
	sqlSearch.Expression = expressionSQL
	sqlSearch.OutputSerializationSelect.JsonBodyOutput.RecordDelimiter = ","
	sqlSearch.InputSerializationSelect.JsonBodyInput.JSONType = "LINES"

	sqlSearchR, err := c.bucketClient.SelectObject(c.key(key), sqlSearch)
	if err != nil {
		return nil, err
	}

	defer sqlSearchR.Close()
	sqlSearchReadAll, err := io.ReadAll(sqlSearchR)
	if err != nil {
		return nil, err
	}
	return sqlSearchReadAll, nil
}

func (c *Client) Delete(key string) error {
	return c.bucketClient.DeleteObject(c.key(key))
}
