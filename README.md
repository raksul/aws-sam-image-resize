# image-resize

## Requirements

* go1.12
* docker

## Setup process

### Installing dependencies

```shell
make deps
```

### Building

```shell
make build
```

### Local development

**Invoking function locally**

```bash
sam local start-lambda
```

## Packaging and deployment

```yaml
make deploy \
  CodeBucket=raksul-image-resize-code \
  StackName=raksul-image-resize \
  VpcId=vpc-xxxxxxxx \
  Subnet1=subnet-xxxxxxxx \
  Subnet2=subnet-xxxxxxxx \
  Bucket=raksul-image-resize
```
