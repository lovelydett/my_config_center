/**
* @description Implement a customized OSS connector
* @author Yuting Xie
* @date Nov 15, 2024
 */

package drivers

import (
	"bytes"
	"io"
	"log"
	"sync"
	"wolf/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OSSConnector struct {
	client *oss.Client
	Bucket *oss.Bucket
}

var ossOnce sync.Once
var ossConn OSSConnector

func GetOSSConnector() *OSSConnector {
	ossOnce.Do(func() {
		ossConfig := config.GetDeployConfig().OSS
		provider, err := oss.NewEnvironmentVariableCredentialsProvider()
		if err != nil {
			log.Fatalf("Failed to create credentials provider: %v", err)
		}
		clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
		clientOptions = append(clientOptions, oss.Region(ossConfig.Region))
		clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))

		// First try internal endpoint
		ossConn.client, err = oss.New(ossConfig.InternalEndpoint, "", "", clientOptions...)
		if err != nil {
			log.Println("Unable to connect to OSS via internal endpoint, try external instead")
			ossConn.client, err = oss.New(ossConfig.ExternalEndpoint, "", "", clientOptions...)
		}
		if err != nil {
			log.Fatalf("Failed to connect to both internal and external OSS endpoints: %v", err)
		}

		ossConn.Bucket, err = ossConn.client.Bucket(ossConfig.BucketName)
		if err != nil {
			log.Fatalf("Failed to get Bucket: %v", err)
		}
	})

	return &ossConn
}

func (oss *OSSConnector) UploadFile(objKey string, filePath string) error {
	return oss.Bucket.PutObjectFromFile(objKey, filePath)
}

func (oss *OSSConnector) UploadBytes(objKey string, buffer []byte) error {
	return oss.Bucket.PutObject(objKey, bytes.NewReader(buffer))
}

func (oss *OSSConnector) UploadStream(objKey string, stream io.Reader) error {
	return oss.Bucket.PutObject(objKey, stream)
}

func (oss *OSSConnector) AppendBytes(objKey string, buffer []byte, offset int64) error {
	_, err := oss.Bucket.AppendObject(objKey, bytes.NewReader(buffer), offset)
	return err
}
