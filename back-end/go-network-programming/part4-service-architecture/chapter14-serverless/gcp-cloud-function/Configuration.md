# GCP Cloud Function 설정

### Config for GCP
* `gcloud init`
  * 브라우저를 통해 SDK 인증
  * 프로젝트 이름은 "yj-go-network-serverless"으로 새로 생성
* 결제 활성화
  * 해당 프로젝트에서 결제 계정 등록
    * [등록 페이지](https://console.cloud.google.com/billing/01AA77-3A3B69-2782F2?hl=ko&project=yj-go-network-serverless)
* `gcloud projects list`
  * 인증된 계정에 있는 프로젝트 리스트 출력

### Development
* `go.mod`
  * 모듈 이름은 반드시 URL 형태로 지정해야함
    * 현재는 `yj-test.com/gcp`
* handler
  * AWS 와 다르게, 진입점에 해당하는 핸들러를 Export 해야함
    * 즉, 메인 핸들러의 이름은 대문자로 시작
    * 현재는 `MainHandler`로 설정

### Deploy to GCP Cloud Function
* Deploy
  * `gcloud functions deploy MainHandler --source ./ --runtime go119 --trigger-http --allow-unauthenticated`
  * `deploy`: 진입점에 해당하는 exported function name(현재는 `MainHandler`)
  * `--trigger-http`: 
  * Output URL
    * `httpsTrigger.url`: HTTP 요청에 대해서 해당 핸들러를 호출한다는 옵션
    * `allow-unauthenticated`: HTTP 요청에 대해서 인증을 요구하지 않는다는 public 옵션
* Test Script
  * `./command-line-test.sh`

### Clean up
* `gcloud functions delete MainHandler`