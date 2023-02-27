# Simple Client
* 가상의 서버(pkg)에 대한 client 코드
* 가상 서버는 테스트 코드에서 Mock Server 로 구현되어 테스트

## 가상 서버 API
* GET simple JSON Data
  * Response - [{name, version}]
* POST simple JSON Data 
  * Request - {name, version}
  * Response - {id}
* POST multipart data
  * Request - {name, version, filedata, Binary Data}
  * Response - {id, filename, size int}