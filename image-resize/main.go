package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	sess = session.Must(session.NewSession())
	svc  = s3.New(sess)
)

func handler(req events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	c := NewConfig(req)
	if c == nil {
		return NewErrResponse(http.StatusBadRequest), nil
	}

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(c.Bucket),
		Key:    aws.String(c.ObjectKey),
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchBucket:
				fallthrough
			case s3.ErrCodeNoSuchKey:
				return NewErrResponse(http.StatusNotFound), nil
			}
		}

		fmt.Printf("Get S3 Err: %#v\n", err)
		return NewErrResponse(http.StatusInternalServerError), nil
	}
	defer resp.Body.Close()

	res, err := Convert(resp.Body, c)
	if err != nil {
		fmt.Printf("Resize Err: %#v\n", err)
		return NewErrResponse(http.StatusInternalServerError), nil
	}

	return events.ALBTargetGroupResponse{
		Body: res,
		Headers: map[string]string{
			"Content-Type": "image/png",
		},
		StatusCode:      200,
		IsBase64Encoded: true,
	}, nil
}

func NewErrResponse(code int) events.ALBTargetGroupResponse {
	return events.ALBTargetGroupResponse{
		Body: http.StatusText(code),
		Headers: map[string]string{
			"Content-Type": "text/plain",
		},
		StatusCode: code,
	}
}

func main() {
	lambda.Start(handler)
}
