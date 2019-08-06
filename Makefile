.PHONY: deps clean build package deploy

deps:
	GO111MODULE=on go get -u ./...

clean:
	rm -rf ./image-resize/image-resize

build:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o image-resize/image-resize ./image-resize

package: build
	sam package --template-file template.yaml --s3-bucket ${CodeBucket} --output-template-file packaged.yaml

deploy: package
	aws cloudformation deploy --template-file packaged.yaml \
		--stack-name ${StackName} --capabilities CAPABILITY_IAM \
		--parameter-overrides Subnet1=${Subnet1} Subnet2=${Subnet2} VpcId=${VpcId} BucketName=${Bucket}
