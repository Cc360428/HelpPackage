package result

import (
	"bytes"
	"fmt"
	"github.com/Cc360428/HelpPackage/beego_utils"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

//
var (
	// websocket Hub 注册
	GlobalOssEndpoint        string
	GlobalOssAccessKeyId     string
	GlobalOssAccessKeySecret string
	GlobalOssBucket          string
)

/*
*
**/
func GlobalOssInit() {
	var err error

	//
	GlobalOssEndpoint, err = beego_utils.HelperConfigSessStringcfg("oss", "endpoint")
	if err != nil {
		GlobalOssEndpoint = "oss-cn-shanghai.aliyuncs.com"
	}

	//
	GlobalOssAccessKeyId, err = beego_utils.HelperConfigSessStringcfg("oss", "accesskeyid")
	if err != nil {
		GlobalOssAccessKeyId = "LTAIwOPiTFbYNtMf"
	}

	//
	GlobalOssAccessKeySecret, err = beego_utils.HelperConfigSessStringcfg("oss", "accesskeysecret")
	if err != nil {
		GlobalOssAccessKeySecret = "rYGRsbjIhoGQbUOl7YdbD7gDqM3zV4"
	}

	//
	GlobalOssBucket, err = beego_utils.HelperConfigSessStringcfg("oss", "bucket")
	if err != nil {
		GlobalOssBucket = "smwyo"
	}
}

/**
*  断点续传文件到 oss， obj 是oss 存放文件名， 如： facesrc/nsb.jpeg , filepath 是文件在本地 地址， 如： d:/faces/nsb.jpeg
**/
func GlobalOssUpoad(obj string, filepath string) (err error) {
	// 创建OSSClient实例。
	client, err := oss.New(GlobalOssEndpoint, GlobalOssAccessKeyId, GlobalOssAccessKeySecret)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(GlobalOssBucket)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 分片大小100K，3个协程并发上传分片，使用断点续传。
	// 其中"<yourObjectName>"为objectKey， "LocalFile"为filePath，100*1024为partSize。
	err = bucket.UploadFile(obj, filepath, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		logs.Error("Error:", err)
		return err
	}
	return nil
}

/**
*  断点续传下载文件， obj 是oss 存放文件名， 如： facesrc/nsb.jpeg , filepath 是文件在本地 地址， 如： d:/faces/nsb.jpeg
**/
func GlobalOssDownload(obj string, filepath string) (err error) {
	// 创建OSSClient实例。
	client, err := oss.New(GlobalOssEndpoint, GlobalOssAccessKeyId, GlobalOssAccessKeySecret)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(GlobalOssBucket)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 分片下载。3个协程并发下载分片，开启断点续传下载。
	// 其中"<yourObjectName>"为objectKey， "LocalFile"为filePath，100*1024为partSize。
	err = bucket.DownloadFile(obj, filepath, 100*1024, oss.Routines(3), oss.Checkpoint(true, ""))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return nil
}

/**
*   上传 bytes 到阿里云 oss 上
**/
func GlobalOssUploadbytes(objName string, sndBytes []byte) (err error) {
	// 创建OSSClient实例。
	client, err := oss.New(GlobalOssEndpoint, GlobalOssAccessKeyId, GlobalOssAccessKeySecret)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(GlobalOssBucket)
	if err != nil {
		logs.Error("Error:", err)
		return err
	}

	// 上传Byte数组。
	err = bucket.PutObject(objName, bytes.NewReader(sndBytes))
	if err != nil {
		logs.Error("snd oss bytes err :: %v", err.Error())
		return err
	}
	return nil
}

/**
*   下载oss 文件 到 bytes 里
**/
func GlobalOssDownloadbytes(objName string) (data []byte, err error) {
	// 创建OSSClient实例。
	client, err := oss.New(GlobalOssEndpoint, GlobalOssAccessKeyId, GlobalOssAccessKeySecret)
	if err != nil {
		logs.Error("Error:", err)
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(GlobalOssBucket)
	if err != nil {
		logs.Error("Error:", err)
		return nil, err
	}

	// 下载文件到缓存。
	body, err := bucket.GetObject(objName)
	if err != nil {
		logs.Error("oss GetObject ( %s ) fail Error :: %v", objName, err.Error())
		return nil, err
	}
	defer body.Close()

	data, err = ioutil.ReadAll(body)
	if err != nil {
		logs.Error("obje readbuf fail, Error: %v", err.Error())
		return nil, err
	}
	return data, nil
}
