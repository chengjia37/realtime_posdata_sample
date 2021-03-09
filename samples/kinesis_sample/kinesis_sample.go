package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kinesis"
)

const deaultRegion = "ap-northeast-1"
const profileName = "kinesis"

func main() {
	mySession := session.Must(
		session.NewSessionWithOptions(session.Options{Profile: profileName}))
	// _ = mySession
	service := kinesis.New(mySession, aws.NewConfig().WithRegion(deaultRegion))
	data := "Hello from GO Lang"
	record := &kinesis.PutRecordInput{
		Data:         []byte(data),
		PartitionKey: aws.String("1"),
		StreamName:   aws.String("pos"),
	}
	kinesisResp, err := service.PutRecord(record)
	if err != nil {
		panic(err)
	}
	fmt.Println(*kinesisResp.SequenceNumber)
	fmt.Println(*kinesisResp.ShardId)
}
