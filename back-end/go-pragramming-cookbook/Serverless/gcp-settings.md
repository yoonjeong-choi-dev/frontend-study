### App Engine Install & Setting
* https://cloud.google.com/appengine/docs/flexible/go/create-app?hl=ko
* gcloud CLI 설치
  * https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-cli-422.0.0-darwin-arm.tar.gz?hl=ko
  * `/Users/yjchoi/google-cloud-sdk`에 압축 해제한 파일들 복사
  * `./google-cloud-sdk/install.sh`
  * `source .bash_profile`
  * `gcloud init`
    * Please enter your numeric choice:  1
    * Enter project ID you would like to use:  yj-golang
  * `gcloud auth login`
* `gcloud components install app-engine-go`
  * Go App Engine Component 설치
* yj-golang 프로젝트에 데이터스토어 생성
  * https://console.cloud.google.com/datastore/


### App Engine Deploy
* `gcloud config set project yj-golang`
* `gcloud auth application-default login`
* Local Test
  * `export GCLOUD_DATASET_ID=yj-golang`
  * `go build .;./go-app-engine`
  * http://localhost:7166/?message= 형태로 GET 테스트
* Deploy
  * `gcloud app deploy`
    * 배포 지역 에러는 https://console.cloud.google.com/home/activity?project=yj-golang 에서 확인
    * Cloud Build API 허용
      * 결제 카드 등록 필요
      * https://console.developers.google.com/apis/api/cloudbuild.googleapis.com/overview?project=yj-golang
  * `gcloud app browse`
    * 배포가 완료된 웹페이지로 이동