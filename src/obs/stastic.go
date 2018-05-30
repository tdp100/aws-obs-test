package obs

import "github.com/aws/aws-sdk-go/aws"
import "github.com/aws/aws-sdk-go/service/s3"
import "github.com/aws/aws-sdk-go/aws/session"
import "github.com/aws/aws-sdk-go/aws/credentials"
import "net/http"
import "net/url"
import "fmt"

const (
	AccessKey = "2zJSTd6dzx8lRFXbb6hH"
	SecretKey = "XV1XvUA7ErF5FhOWHHdGklZszpuqR32yEz51oq0C"
	DefaultRegion ="cn-north-1"
	DefaultRegionEndpoint = "obs.myhwclouds.com"
	ProxyURL = "http://127.0.0.1:8087"
	UseProxy = false
)

type S3client struct {
	S3srv *s3.S3
}

//NewS3Client 生成s3实例
func NewS3Client() *S3client {
	proxy := func(_ *http.Request) (*url.URL, error) {
		if UseProxy {
			return url.Parse(ProxyURL)
		}
		return nil, nil
	}
	
	httpClient:= &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
		},
	}

	ss:=session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(AccessKey, SecretKey, ""),
		DisableSSL: aws.Bool(true),
		HTTPClient: httpClient,
		Region: aws.String(DefaultRegion),
		Endpoint: aws.String(DefaultRegionEndpoint),
	})

	s3Instance:= s3.New(ss)
	return &S3client{
		S3srv: s3Instance,
	}
}

func (self *S3client) headBucket(bucketName string) error {
	_, err:=self.S3srv.HeadBucket(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Printf("Failed to head bucket, bucket:%s, err:%s", bucketName, err.Error())
		return err
	}
	return nil
}

func (self *S3client) getBucketLocation(bucketName string) (string, error) {
	out, err:=self.S3srv.GetBucketLocation(&s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		fmt.Printf("Failed to get bucket location, bucket:%s, err:%s",bucketName, err.Error())
		return "", err
	}
	return aws.StringValue(out.LocationConstraint), nil
}