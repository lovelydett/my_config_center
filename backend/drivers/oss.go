/**
* @description Implement a customized OSS connector
* @author Yuting Xie
* @date Nov 15, 2024
 */

package drivers

import (
	"log"
	"wolf/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// The bucket is to be exported
var Bucket *oss.Bucket

func init() {
	ossConfig := config.GetDeployConfig().OSS
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		log.Fatalf("Failed to create credentials provider: %v", err)
	}
	clientOptions := []oss.ClientOption{oss.SetCredentialsProvider(&provider)}
	clientOptions = append(clientOptions, oss.Region(ossConfig.Region))
	clientOptions = append(clientOptions, oss.AuthVersion(oss.AuthV4))

	// First try internal endpoint
	client, err := oss.New(ossConfig.InternalEndpoint, "", "", clientOptions...)
	if err != nil {
		log.Println("Unable to connect to OSS via internal endpoint, try external instead")
		client, err = oss.New(ossConfig.ExternalEndpoint, "", "", clientOptions...)
	}
	if err != nil {
		log.Fatalf("Failed to connect to both internal and external OSS endpoints: %v", err)
	}

	Bucket, err = client.Bucket(ossConfig.BucketName)
	if err != nil {
		log.Fatalf("Failed to get Bucket: %v", err)
	}
}
