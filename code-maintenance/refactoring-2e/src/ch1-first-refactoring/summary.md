# Chapter 1. Refactoring - First Example
* `index.mjs`를 통해 해당 기능 사용
  * 단계별로 리팩터링하는 과정을 `step/step*.mjs` 소스코드에 기록
* `data/*` 에 `json` 데이터 저장
  * `plays.json` 공연 기본 정보
    * 이름 및 장르
  * `invoices.json` 공연 요청 정보
    * 고객 이름, [공연 이름, 관객 수]

### Step 0
* 요구 사항
  * 공연 기본 정보 및 공연 요청 입력
  * 공연료 및 보너스 포인트를 계산하여 청구 내역 출력
* 제약 사항
  * 현재 공연 장르는 단 2개
* 문제점 : 새로운 기능을 추가하기 어려운 구조
  * 추가 기능 1: 공연 정보를 특정 포멧에 맞춰 출력하는 기능
    * 예를 들어 HTML 출력 기능
    * 현재는 공연 정보가 문자열로 반환되어 추가 기능을 위해서는 함수 로직을 복사해서 새로운 함수를 만들어야 함
    * 이유 : 출력을 위한 정보를 추출+계산하는 로직과 출력을 하는 로직이 합쳐져 있음
  * 추가 기능 2: 공연 장르나 공연료 계산 정책 변경
    * 현재는 장르에 대한 공연료 및 보너스 포인트 계산이 하드 코딩
    * 하나의 엔티티 계산 내에 조건부 로직이 있음 => 코드 파악이 어려움 

### 리팩터링 첫 단계 : 테스트 코드 작성
* 테스트 코드의 필요성
  * 리팩터링 이전 원본 코드에 대한 결과와 비교하는 과정이 필요
  * 리팩터링 결과 로직 자체에 대한 변경이 없음을 확신할 수 있는 도구
* 각 단계마다 실행해야하는 것
  * 테스트 코드 실행
    * 해당 단계의 리팩터링이 제대로 되었는지 확인 가능
  * 깃 커밋을 통한 로깅
* 각 단계(프로그램 수정)은 매우 작은 작업으로 구성해야 한다
  * 중간의 실수가 있어도 빠르게 롤백 가능
  * 조금씩 변경하고 테스트하는 것이 핵심

### Step 1 : 함수 추출하기
* 각 장르에 대한 공연료(`currentAmount`) 계산하는 부분
  * 현재는 `switch` 구문을 분석을 통해 해당 로직을 파악해야 한다
  * 별도의 함수로 추출함으로써, 나중에 다시 볼 때 빠르게 해당 로직을 파악 가능
* 추출하는 로직에서 사용하는 변수들 파악
  * 로직 자체에서만 사용하는 변수들(`performance`,`play`)은 함수의 매개변수로 설정
  * 로직 내에서 값이 변하는 변수(`currentAmount`)
    * 로직의 목표가 `currentAmount`를 계산하는 것
    * 추출한 함수의 반환값으로 사용
    * 함수의 반환을 의미하도록 변수명을 `result`로 변경
* 비슷한 방식으로 각 장르에 대한 보너스 포인트(`volumeCredits`) 계산하는 부분도 함수로 추출 가능
  * step2 에서 진행
  * `getVolumeCredit` 함수
  

### Step 2 : 로컬 변수 제거
* 로컬(임시) 변수를 질의 함수로 바꾸거나 함수로 추출하여 제거하는 이유
  * 추출 작업이 쉬워 진다(See Step3)
* 제거하는 로컬 변수
  * `play` from `getAmountFor(performance, play)`
  * `currentAmount`
  * `totalAmount`
  * `volumeCredits`
* 지역 변수를 무조건 질의 함수로 바꾸는게 맞나?
  * 가장 상단에 상수로 저장해놓고 처리하는 경우가 많지 않나?

### 중간 점검 for Step 1~2
* 기능 분리 및 로컬 변수 제거를 위한 중첩 함수들


### Step 3 : 단계 쪼개기
* 추가 기능 : 다양한 출력 방식
  * HTML format 으로 출력하기
* 현재 함수는 2단계로 구성
  * 필요한 데이터 처리
  * 처리한 데이터를 이용하여 특정 포멧(문자열, HTML) 출력
* 모델과 뷰의 분리
  * 하나의 모델(처리한 데이터)에 대한 다양한 뷰를 만들 수 있음
  * 이 단계에서 중요한 것은 로우 데이터를 어떤 식으로 구성하여 뷰에게 전달할지 i.e 중간 데이터 설계
* 단계를 쪼개는 방법 : 함수 추출하기
  * 전체 로직을 크게 두 함수로 나눈다
  * 이후, 각 함수에 사용하는 로직들을 분리
    * Step2 작업으로 인해 로직 분리가 쉬워짐

### Step 4 : 조건부 로직을 다형성으로 바꾸기
* 문제가 되는 부분
  * `createStatementData.mjs` 의 `getAmountFor` & `getVolumeCredit`
    * 장르에 따라 계산 방식이 다름
    * 구현 내에 조건부 로직이 많음
  * 새로운 장르가 추가되거나 계산 정책이 달라지는 경우
    * 위 두 함수에 모두 접근하여 로직 추가/변경이 필요
    * 장르에 대한 새로운 계산 로직이 추가되는 경우, 위 두 함수처럼 새로운 함수를 만들고 그 안에 조건부 로직을 구현해야 함
    * 계산 함수가 많아질수록 추가/변경해야 하는 코드가 많아짐
* 다형성을 이용하여 계산 로직 분리
  * 장르(타입 코드) 대신 서브 클래스 사용하기
  * 타입에 해당 하는 서브 클래스 객체를 생성하는 팩토리 함수 패턴 적용