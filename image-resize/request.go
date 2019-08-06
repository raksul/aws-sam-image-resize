package main

import (
	"os"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
)

var (
	Bucket = os.Getenv("BUCKET_NAME")
)

func NewConfig(r events.ALBTargetGroupRequest) *Config {
	width, _ := strconv.Atoi(r.QueryStringParameters["w"])
	height, _ := strconv.Atoi(r.QueryStringParameters["h"])
	objectKey := r.Path

	if width <= 0 && height <= 0 {
		return nil
	}

	return &Config{
		Bucket:    Bucket,
		ObjectKey: objectKey,
		Width:     width,
		Height:    height,
	}
}
