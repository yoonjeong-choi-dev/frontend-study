# Terraform 설정

### CLI
* 한국 지역에서 사용 가능한 쿠버네티스 버전 목록
  * `az aks get-versions --location koreasouth -o table`
* azure 인증을 위한 서비스 주체 생성
  * `az ad sp create-for-rbac --role="Contributor" --scopes="/subscriptions/{account_id}"`
    * account_id: `az account show` 출력의 id 필드
  * variable.tf 파일의 client_id 및 client_secret 입력 시 사용
* 쿠버네티스 클러스터 인증 정보 설정
  * `az aks get-credentials --resource-group yjtube --name yjtube`
* Dashboard
  * `kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml`
  * `kubectl proxy`
  * http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login

### .tf 파일 설명
* providers.tf
  * 테라폼 플러그인 정의 및 설정
  * 플러그인들의 버전 관리를 위한 파일
* variables.tf
  * 테라폼 스크립트들에서 사용할 전역 변수들
  * `location` 값 확인하는 곳: [link]("https://github.com/claranet/terraform-azurerm-regions/blob/master/REGIONS.md") 

