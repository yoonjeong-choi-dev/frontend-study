# Chapter 2. Refactoring Principals
### Contents
1. 리팩터링 정의
2. 리팩터링 규칙(2개의 모자)
3. 리팩터링하는 이유
4. 리팩터링을 해야 하는 시기
5. 리팩터링 시 고려할 문제
6. 리팩터링, 아키텍처, 애그니(YAGNI)
7. 리팩터링과 소프트웨어 개발 프로세스
8. 리팩터링과 성능
9. 리팩터링 자동화


### 리팩터링 정의와 규칙(2개의 모자)
* 리팩터링 정의
  * 명사로서의 정의
    * 겉보기 동작은 그대로 유지
    * 코드를 이해하고 수정하기 쉽도록 내부 구조를 변경하는 **기법**
  * 동사로서의 정의
    * 겉보기 동작은 그대로 유지
    * 여러가지 리팩터링 기법을 적용하여 소프트웨어를 재구성한다
  * 즉, **특정한 방식**에 따라 코드를 정리하는 것만이 리팩터링
    * 코드를 정리하는 모든 행위가 다 리팩터링은 아님
  * 겉보기 동작은 그대로 유지
    * 사용자 관점에서 달라지는 점이 없음(인터페이스 유지)
    * 리팩터링 과정 중에 발견한 버그도 그대로 유지
* 동작을 보존하는 작은 단계들을 거쳐 코드를 수정하고, 전체적으로 큰 변화를 이끄는 행위
  * 각 단계 전후로 동작이 보존
  * 언제든지 리팩터링을 멈출 수 있음
* vs 최적화 
  * 두 행위 모두 코드를 변경하고, 프로그램의 전반적인 기능은 유지
  * 리팩터링
    * 목적: 코드를 이해하고 수정하기 쉽게 만드는 것
    * 결과적으로 성능이 나빠질수도 좋아질수도 있음
  * 최적화
    * 목적: 오로지 속도 개선
    * 코드가 다루기 어렵게 변경될 수 있음


### 리팩터링 규칙(두개의 모자)
* 소프트웨어 개발 시 2개의 모드
  * 기능 추가
  * 리팩터링
* 기능 추가
  * 기존 코드는 건드리지 않고 추가할 기능에만 집중
  * 추가 기능에 대한 테스트를 우선 준비
  * 테스트 통과율을 통해 진척도 확인
* 리팩터링
  * 기능 추가는 절대 하지 않고, 코드 재구성에만 집중
  * 테스트 케이스도 추가하지 않음
    * 리팩터링 과정에서 인터페이스 변경 시에만 테스트 수정


### 리팩터링하는 이유
* 소프트웨어 설계가 좋아진다
  * 프로그램 규모가 커질수록 코드의 아키텍처는 무너지기 쉬어짐
  * 규칙적인 리팩터링은 코드의 아키텍처를 좋은 상태로 유지시켜줌
* 소프트웨어를 이해하기 쉬워진다
  * 프로그래밍에서 중요한 것은 사람
    * 기능 구현에만 신경쓰게되면 나중에 코드를 보는 사람을 배려하지 못함
    * 유지 보수 단계에서 코드를 이해하는데 시간이 더 걸림
  * 기억할 필요가 있는 것들은 최대한 코드에 담아야 함
    * 코드를 이해하는데 사전 지식이 요구되면 안됨
* 버그를 쉽게 찾을 수 있다
  * 코드를 이해하기 쉬워지기 때문에 문제점을 찾는게 빨라짐
* 개발 속도를 높일 수 있다
  * 새로운 기능 추가 시 발생하는 문제
    * 기능 구현 자체는 오래 걸리지 않음
    * 추가하는 기능을 코드베이스에 녹여내는 방법을 찾는데 시간이 더 걸림
  * 설계가 잘된 코드베이스
    * 새로운 기능을 추가할 지점과 어떻게 녹여낼지 쉽게 파악 가능
    * 모듈화가 잘되어 있어 작은 부분만 이해해도 됨
  * 완벽하게 구현해야 한다는 압박감이 사라짐
    * 리팩터링이 없으면, 설계부터 구현까지 완벽하게 끝나야 기능 추가가 완료됨
    * 언제든지 리팩터링을 통해 개선할 수 있음

### 리팩터링을 해야 하는 시기
* 3의 법칙
  * 비슷한 작업을 세번하게 되면 리팩터링을 해야하는 시점
* 준비를 위한 리팩터링
  * 기능을 쉽게 추가하게 만들기
  * 리팩터링하기 가장 좋은 시점은 기능을 새로 추가하기 직전
  * 버그 픽스하는 시점도 좋은 리팩터링 시기
    * 버그가 발생하는 코드들을 한 곳에 합치는 방식
    * 한 곳에 집중되어 있기 때문에 버그 픽스하기 편해짐
* 이해를 위한 리팩터링
  * 코드를 수정하기 위해서는 코드를 이해해야 함
    * 이해하는 과정에서 해당 코드의 의도를 더 명확하게 할 수 있는지 확인
    * 자신이 이해한 내용을 코드에 옮겨 담을수 있도록 노력
  * 리팩터링으로 코드가 이해하기 쉬워지면 눈에 보이지 않았던 개선점도 발견 가능
* 쓰레기 줍기 리팩터링
  * 코드를 파악하는 과정에서 비효율적인 코드를 발견한 상황
    * 로직이 쓸데없이 복잡
    * 비슷한 동작을 하는 함수가 여러개(함수 매개변수화하기로 개선)
  * 현재 작업과 리팩터링 사이에서의 절충안
    * 수정이 간단한 경우에는 리팩터링 후 작업 진행
    * 시간이 걸리는 리팩터링인 경우 원래 작업 진행 후 리팩터링
  * 중요한 것은 처음 봤을 때보다 조금이나마 개선되어야 한다는 것(캠핑 규칙)
* 계획된 리팩터링과 수시로 하는 리팩터링
  * 저자는 리팩터링 계획을 따로 잡기보다는 수시로 하는 편
    * 프로그래밍 과정에 자연스럽게 녹임
    * 준비를 위한/이해를 위한/쓰레기 줍기 리팩터링 기회가 주어질 때마다 진행
    * 리팩터링은 프로그래밍(개발)과 구분되는 별개의 활동이 아님
  * 잘 작성된 코드도 수많은 리팩터링 과정을 거쳐야 함
  * 코드를 수정할 때, 수정하기 쉽게 정돈하고 그다음 쉽게 수정하자(켄트백)
    * 기능 추가도 코드의 수정 -> 추가하기 쉽게 코드를 수정
  * 계획된 리팩터링을 하게 되는 상황을 최대한 피하도록 해야 함
    * 계획된 리팩터링이 필요하다는 것은 이미 코드베이스가 무너졌다는 것
* 리팩터링과 기능추가 커밋 구분
  * 기능 추가와 리팩터링이 밀접하게 엮인 경우가 많음
    * 준비를 위한 리팩터링
    * 둘을 구분하게 되면, 리팩터링에 대한 맥락 정보가 사라짐
    * ```
      질문
      - PR 기준 : 우리는 지라 이슈 기준
      - 대규모 리팩토링을 접한적이 있는지?
      ```
* 리팩터링에 대한 설득
  * 리팩터링을 해야한다는 것을 관리자에게 설득해야 하는 상황
  * 개발자의 임무는 효과적으로 소프트웨어를 **빠르게** 만드는 것
    * 리팩터링은 개발 속도를 홀리는데 매우 효과적
    * 리팩터링은 코드를 쉽게 이해하게 만들어 줌
    * 기능을 추가하거나 버그를 고치는 작업을 빠르게 하기 위해서는 코드를 빠르게 파악해야 함
* 리팩터링하면 안되는 상황
  * 지저분한 코드를 발견하여도 수정할 필요가 없는 상황
    * 해당 코드를 현재 외부 API 사용하듯 사용하는 상황
    * 즉 지저분한 코드를 이해할 필요가 없는 경우
    * 해당 코드를 이해해야 하는 상황이 왔을 때 리팩터링을 해야 효과적
  * 처음부터 새로 작성하는 것이 더 쉬운 상황

### 리팩터링 시 고려할 문제
* 무언가를 언제 어디에 적용할지는 판단하는 기준
  * 리팩터링, 테스트와 같은 작업은 오로지 경제적인 이유
  * 건강한 코드베이스는 빠른 개발 업무로 나아간다
  * 궁극적인 목적은 개발 속도를 높여서, 더 적은 노력으로 더 많은 가치를 창출하는 것
* 코드 소유권
  * 리팩터링을 하고자 하는 코드를 변경하지 못하는 상황
    * 리팩터링을 하고자 하는 코드에 소유권이 없는 경우
    * 바꾸려는 함수가 공개된 인터페이스(Public API)
      * 클라이언트 측 코드를 변경할 수 없음
  * 코드 소유권은 팀 단위로 두는 것이 좋음
    * 팀원이라면 누구든지 팀이 소유하는 코드 수정 가능
    * 팀원마다 각자가 책임지는 영역이 있을 수는 있음
      * 해당 영역의 변경 사항을 관리하는 것을 의미
      * ```
        질문
        - 같은 코드베이스에 다수의 코드 오너가 있는 경우가 있었는지?
        - 우리는 영역에 따라서 코드 오너 설정 -> PR 시 반드시 오너 승인 필요
        ```
* 브랜치
  * fdaf
  * ```
    질문
    1. 머지 vs 통합
    - 전회사 : fork 따서 작업하고 PR(open source style)
    - 지금 회사 : branch 따서 작업하고 PR
    2. branch 기준
    - 피처 단위
    ```
* 테스팅
  * 리팩터링의 특징은 소프트웨어의 겉보기 동작은 유지된다는 것
    * 작은 단계 별로 리팩터링을 하기 때문에 문제가 생겨도 오류가 되는 코드의 양은 적음
    * 원인을 찾지 못해도 롤백하면 됨
  * 중요한 것은 오류가 발생했다는 것을 빠르게 인지하는 것
    * 오류를 파악하기 위한 테스트 스위트가 필요
    * 리팩터링을 하기 위해서는 자가 테스트 코드가 필수적
  * 테스트 코드의 효과
    * 안전하게 리팩터링 가능
    * 안전하게 새로운 기능 추가 가능
    * CI 과정에서도 테스팅을 통해 안전하게 진행 가능
  * ```
    질문
    1. 각자 회사에서의 테스트 코드?
    - 저번에 내가 말한 거는 e2e 쪽
    - 유닛테스트(리액트 jest) 나 UI 동작 관련 테스트는 없는지
    2. 각자 회사에서의 CI/CD 환경
    - 나는 깃헙 액션 사용
    - 코드 오너 승인 + 필수 테스트 통과해야 머지 가능
    - vercel 통해서 빌드 -> 빌드한거로 테스트하고 깃헙 통해서 staging/prod 배포 가능
    ```
* 레거시 코드
  * 레거시 코드의 특징
    * 대체로 복잡
    * 테스트가 없는 경우가 많음
  * 레거시 코드 리팩터링
    * 레거시 시스템을 파악할 때 리팩터링이 도움이 됨
    * 보통 테스트 코드가 없으므로, 리팩터링하기 위해서는 테스트 보강이 필요
  * 레거시 시스템에서의 테스트 보강
    * 테스트를 염두해 둔 시스템만이 테스트가 쉽다
    * 테스트를 염두하지 않은 레거시 시스템에서는 테스트 보강이 어려움
    * 테스트를 할 틈새를 찾아 테스트를 추가하는 것이 주된 해결책
    * 테스트 할 틈새를 만들 때 리팩터링 활용 가능
* 데이터베이스(TODO)
  * 진화형 데이터베이스 설계 및 데이터베이스 리팩터링
    * 큰 변경들을 쉽게 조합하고 다룰수 있는 데이터 마이그레이션 스크립트 작성
    * 접근 코드와 스키마에 대한 구조적 변경을 위 스크립트를 통해 처리하게 통

### 리팩터링, 아키텍처, 애그니(YAGNI)
* 소프트웨어 아키텍처 w/o 리팩터링
  * 선제적인 아키텍처
  * 개발 시작 전에 아키텍처 설계가 완료되어야 함
  * 아키텍처를 설계하기 위해서는 요구사항을 모두 파악하고 있어야 함
    * 실무에서는 요구 사항을 사전에 모두 파악이 불가능
    * 운영 단계에서 변경되거나 추가되는 요구 사항이 많음
  * 유연성 메커니즘
    * 추후 변경될 요구 사항을 유연하게 대처하기 위한 방법
    * 예상 시나리오에 대응하는 매개변수를 모두 추가하는 방식
    * 추측을 통한 구현이므로 당장의 쓰임에 비해 구현이 복잡해짐
* 소프트웨어 아키텍처 with 리팩터링
  * 리팩터링은 요구사항 변화에 쉽게 대응하도록 코드 베이스를 설계하게 해줌
  * 현재까지 파악한 요구사항만을 해결하는 소프트웨어 설계
    * **미래의 불확실한 요구사항을 추측하여 유연성 메커니즘을 심지 말아야 함**
    * 요구사항을 더 잘 이해하게 되면 그때 리팩터링
  * 아키텍처의 복잡도에 지장을 주지 않는 선에서만 유연성 메커니즘을 도입
    * 복잡도를 높일 수 있는 경우에는 검증 이후 도입
  * 구현의 유연함 정도를 적정선에서 절충해야 함
    * 예상되는 변경을 미리 반영하지 않았을 때, 이후 리팩터링이 더 힘들다면 구현이 더 유연해져야 함
* 간결한 설계, 점신적 설계, YAGNI
  * 리팩터링을 전제로 하는 아키텍처 설계 방식
  * YAGNI(you aren't going to need)
    * 당장에 필요한 기능만으로 최대한 간결하게 설계하라
    * 미래의 필요할 것이라고 예상하고 만든 기능 상당수는 쓰이지 않거나, 미래의 요구사항을 제대로 반영하지 못하는 경우가 많음

### 리팩터링과 소프트웨어 개발 프로세스
* 개발 프로세스의 중요한 3가지 요소
  * 자가 테스트 코드
  * 지속적 통합(CI)
  * 리팩터링
* 자가 테스트 코드의 필요성
  * 리팩터링의 결과를 검증해야 함(겉보기 동작은 그대로 유지되는가?)
  * 지속적 통합에서 통합된 결과를 검증해야 함
* 지속적 통합의 필요성
  * 다른 사람의 작업을 방해하지 않고 언제든지 리팩터링을 할 수 있어야 함
  * 리팩터링 결과를 빠르게 공유해야 함
  * 소프트웨어를 언제든지 릴리즈할 수 있는 상태로 유지
    * 비즈니스 요구 사항을 빠르게 반영 가능
    * 문제 발생 시, 롤백을 통해 위험 요소 감소 가능

### 리팩터링과 성능
* 리팩터링을 하면 소프트웨어가 느려질 수도 있음
  * 저자는 이해하기 쉽게 리팩터링을 하기 위해 성능을 포기하는 방향으로 수정하는 경우가 많았음
  * BUT 리팩터링 결과 수정이 쉬어져 성능 튜닝하기 편해짐
  * 리얼타임 시스템을 제외한 시스템 성능 개선 비결
    * 튜닝하기 쉽게 만들고(리팩터링)
    * 원하는 속도까지 튜닝
* 빠른 소프트웨어를 만드는 방법
  * 시간 예산 분배 방식
  * 끊임없이 관심을 기울이자
  * 성능 신경쓰지 말자
* 방법 1: 시간 예산 분배 방식
  * 성능이 중요한 하드 리얼타임 시스템에서 사용
  * 컴포넌트마다 자원(시간과 공간) 예산을 할당
  * 컴포넌트는 할당받은 예산을 초과하지 않도록 구현되어야 함
* 방법 2: 끊임없이 관심을 기울이자
  * 개발자라면 누구나 높은 성능을 유지하기 위해 노력함
  * 성능 개선을 위해 코드를 수정하다보면 코드베이스가 다루기 어렵게 변함
  * 개선이 특정 동작에만 관련된 경우가 많음
    * 즉 컴파일러, 런타임 엔진, 하드웨어의 동작에 대해서는 알지 못하고 코드가 변경
    * 제대로 된 성능 개선이 안됨
* 방법 3: 성능 신경쓰지 말자
  * 성능에 대한 사실
    * 프로그램 전체 코드 중 극히 일부에서 대부분의 시간을 소비함
    * 코드 전체에 대해서 성능 개선을 하는 것은 의미가 없음
  * 성능 최적화를 반드시 해야하는 상황이 아니라면 코드를 다루기 쉽게 개발
  * 성능 최적화에서는 다음과 같은 절차를 따름
    * 프로파일러를 이용하여 코드 분석
    * 분석 결과를 통해 성능에 영향을 주는 부분만 집중해서 최적화
    * 리팩터링처럼 최적화를 위한 수정도 작은 단계로 나눠서 진행

### 리팩터링 자동화
* 자동 리팩터링 기능을 지원하는 IDE
  * 코드를 구문 트리(syntax tree)로 변환하여 단순 에디터 이상으로 리팩터링 지원
  * 코드와 관련된 주석도 함께 변경해주는 기능을 지원하기도 함
* 정적 타입 언어의 경우 제공되는 리팩터링 기능이 많아짐
  * 동일한 메서드 이름이어도 현재 이름 변경을 하는 클래스에 대해서만 변경
* 자동화 툴을 사용한다하여도, 테스트를 통해 제대로 되었는지 검증 필요
* 언어 서버
  * 구문트리를 구성해서 에디터에 API 형태로 제공하는 소프트웨어
  * 자신의 에디터를 리팩터링 기능을 위해 커스터마이징 가능

















