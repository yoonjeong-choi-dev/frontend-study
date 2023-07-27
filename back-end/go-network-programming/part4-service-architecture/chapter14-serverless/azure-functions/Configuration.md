# Azure Functions 설정

### Config for Azure CLI
* `az login`
  * 브라우저를 통해 SDK 인증
* azure cli 툴은 Azure Functions 기능을 지원하지 않아 다음의 명령어를 통해 도구 설치 필요
  * Reference: [Azure Functions](https://learn.microsoft.com/ko-kr/azure/azure-functions/functions-run-local?tabs=v4%2Cmacos%2Cjava%2Cportal%2Cbash#v2)
  * ```
    brew tap azure/functions
    brew install azure-functions-core-tools@4
    brew link --overwrite azure-functions-core-tools@4
    ```

### Development
* Create Custom Handler
  * `func init --worker-runtime custom`
    * custom 인자를 통해 커스텀 핸들러 초기화
  * `customHandler.enableForwardingHttpRequest` 필드를 true 로 설정하여 HTTP 요청 포워딩 활성화
    * Azure 서비스가 클라이언트와 커스텀 핸들러 사이의 프락시 역할
    * `customHandler.description.defaultExecutablePath` 필드가 프락시 진입점
* 커스텀 핸들러에 대한 구성 설정
  * 커스텀 핸들러 구성 폴더 이름이 애저 Function Name
  * `CustomHandlerFunction/function.json`
  * 해당 파일은 Azure Function 에 들어오는 HTTP 요청 트리거에 바인딩되는 설정들로 구성
* 메인 파일
  * 다른 클라우드 서비스와 다르게 핸들러가 아닌 서버 자체에 대한 구현이 필요
  * 메인 핸들러는 gcp 배포 시에 구현한 핸들러 사용
* Local Test
  * `go build -o main.exe main.go`
    * `customHandler.description.defaultExecutablePath` 필드 값과 동일한 실행파일 생성
  * `func start`

### Deploy to Azure Functions
* Build
  * `GOOS=windows go build -o main.exe main.go`
    * 코드는 Azure 운영체제인 windows에서 동작할 예정
* Deploy
  * Create Azure Function App
    * Create Pre-requisite Entities
      * `az group create --name yjGoWithServerless --location eastus`
        *  Resource Group
      * `az storage account create --name yjgoserverless --location eastus --resource-group yjGoWithServerless --sku Standard_LRS`
        * Storage name
        * `--name`: 스토리지 이름은 소문자로만 구성되어야 함
  * Create Function App
    ```
      az functionapp create --resource-group yjGoWithServerless \
      --consumption-plan-location eastus --runtime custom \
      --functions-version 3 \
      --storage-account yjgoserverless \
      --name yj-azure-serverless
    ```
  * `func azure functionapp publish yj-azure-serverless --no-build --force`
* Test Script
  * `./command-line-test.sh`

### Clean up
* https://ap-northeast-2.console.aws.amazon.com/lambda 에서 삭제