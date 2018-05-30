package obs

import "testing"
import "fmt"

func TestQueryBucketInfo(t *testing.T) {
	srv:=NewS3Client()
	err:=srv.headBucket("tdpdfv")
	if err !=nil {
		t.Error(err.Error())
	}
	fmt.Print("success. \n")
	t.Log("success.")
}

func TestGetBucketLocation(t *testing.T) {
	srv:=NewS3Client()
	location, err:=srv.getBucketLocation("tdpdfv")
	if err !=nil {
		t.Error(err.Error())
	}
	fmt.Printf("success. location:%s \n",location)
	t.Log("success.")
}