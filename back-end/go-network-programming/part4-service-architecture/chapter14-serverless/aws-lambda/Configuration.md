# AWS Lambda 설정

### Config for AWS CLI
* `aws configure`
  * Access Key Pair 생성 후 설정
  * 지역은 ap-northeast-2
* lambda 서비스에 대한 신뢰 정책 등록
  * 해당 파일과 동일한 경로에서 실행
  * `aws iam create-role --role-name "yj-golang-network-lambda" --assume-role-policy-document file://trust-policy.json`
  * 호출 output으로 나오는 `Arn` 필드 확인

### Deploy to AWS Lambda
* Reference: [AWS Lambda Docs](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
* Build & Compress
  * `GOOS=linux GOARCH=amd64 go build -o main aws-main.go` 
  * `zip main.zip main`
* Deploy
  *  lambda 서비스에 대한 신뢰 정책 등록의 output `Arn` 필요
  * `aws lambda create-function --function-name yj-go-network-serverless-app --runtime go1.x --role arn:aws:iam::085771716532:role/yj-golang-network-lambda --handler main --zip-file fileb://main.zip`
* Update
  * `aws lambda update-function-code --function-name "yj-go-network-serverless-app" --zip-file "fileb://main.zip"`
* Test Script
  * `./command-line-test.sh`

### Clean up
* https://ap-northeast-2.console.aws.amazon.com/lambda 에서 삭제