# Main Project - User Service

### TODO
* 유저 권한 필드 추가
* 권한에 따라 접근 가능한 API 
  * getUsers : 전체 유저 데이터 가져오기
* 유저 비밀번호 해시해서 저장

### Chapter 3. Controller
* Init NestJS Project
* Add User Controller
  * `npx @nestjs/cli g co Users` 
  * Add DTO for Request and Response

### Chapter 4. Provider
* Add User Service(Provider)
  * `npx @nestjs/cli g s Users`
    * `UsersService` is inserted into `Appmodule.providers`
* Add Email Service
  * nodemailer - Email 3rd Party
  * User Service 에서 사용

### Chapter 5. Module
* User 모듈 분리
  * `npx @nestjs/cli g mo Users`
* Email 모듈 분리
  * `npx @nestjs/cli g mo Email`
  * 외부에서 사용해야 하는 모듈 내 `provider`를 `exports`에 추가
* AppModule 변경
  * 추가한 모듈들은 `imports`에 추가
  * `provider` 및 `controller` 더 이상 필요 없음

### Chapter 6. Dynamic Module Config
* 각 실행 환경에 해당하는 구성을 관리하는 모듈
  * 환경 변수를 위한 `@nestjs/config` 라이브러리 설치
  * 모듈에서 현재 환경에 맞는 `.env` 파일 읽어서 설정
* `.env` 파일
  * 환경에 따라 달라지는 설정
    * DB, API HOST 등
  * 외부(코드)에 공개하면 안되는 정보들
    * 3rd Party 계정, API Key
* `start:dev` 스크립트 변경
  * 개발 환경임을 명시
* `src/config` 디렉터리에서 `ConfigModule`에 등록할 관리할 `ConfigFactory`
  * emailConfig
    * Chapter4 에서 만든 Email 서비스 내 하드코딩한 이메일 계정 정보 설정
  * `./env' 
    * 환경 변수 파일들 관리
    * nest-cli.json 변경 필요
      * Nest 기본 빌드 옵션은 .ts 파일을 제외
      * 빌드 시, 현재 환경에 맞는 .env 파일을 넣어주어야 함

### Chapter 7. Validation Pipeline
* `class-validator` 및 `class-transformer` 설치
  * `class-validator`에서 제공하는 데코레이터를 통해 DTO 레벨에서 검증 조건 적용 가능
  * `class-transformer`에서 제공하는 데코레이터를 통해 DTO 레벨에서 pre-processing 가능
  * 요청과 관련된 DTO 제약 조건 추가
* NestJS 에서 제공하는 `ValidationPipe` 파이프를 전역으로 적용
  * 파이프라인은 라우트 핸들러(컨트롤러)로 요청이 전달되기 전에 객체 변환이 가능
  * 특히, 요청 검증에 대한 작업은 전역적으로 발생하므로 해당 파이프라인을 전역으로 적용
* Custom Validation Decorator for `class-validator`
  * https://github.com/typestack/class-validator#custom-validation-decorators
  * `src/utils/decorators/not-in.ts`

### Chapter 8. Database
* 사용하는 데이터베이스 및 ORM
  * MySql & TypeOrm
  * `npm i typeorm@0.3.7 @nestjs/typeorm@9.0.0 mysql2`
* AppModule 변경
  * orm 설정
  * 설정 시 필요한 연결 정보는 환경 변수로 설정
* User Entity
  * 테이블 정의
  * user module 추가 이후, 서비스 구현
* Transaction 설정
  * 방법 1: TypeOrm QueryRunner 이용하여 커넥션 및 커밋,롤백 관리
  * 방법 2: transaction 콜백 함수
* Migration via TypeOrm
  * `package.json`에 스크립트 추가
  * `ormconfig.ts`에 데이터베이스 연결 정보 설정
    * 해당 파일은 ConfigModule 이 환경변수를 읽기 전에 컴파일되어 환경 변수 사용하면 에러 발생
    * 환경변수 사용하기 위해서는 typescript(`tsconfig.json`) 컴파일 옵션 변경 필요
    * TODO: 연결 정보 환경변수로 빼내기
  * 관련 스크립트
    * 테스트 시에는 `synchronize` 옵션 false 설정
    * `npm run typeorm:create src/migrations/CreateUserTable`
      * empty MigrationInterface 구현체 생성
    * `npm run typeorm migration:generate src/migrations/CreateUserTable -- -d ./ormconfig.ts`
      * MigrationInterface 구현체 생성 및 메서드 구현
    * `npm run typeorm migration:run -- -d ./ormconfig.ts`
      * MigrationInterface 구현체를 이용하여 테이블 생성
    * `npm run typeorm migration:revert -- -d ./ormconfig.ts`
      * 마이그레이션 이전 버전으로 revert

### Chapter 10. JWT Auth
* jwt library
  * jwtwebtoken
  * `npm i -D @types/jsonwebtoken`
  * `npm i jsonwebtoken`
* 추가한 파일들
  * authConfig for secret key used in JWT
  * auth module
    * injected to user service
* 구현한 서비스
  * sign-in

### Chapter 11. Logger
* winston library
  * `npm i nest-winston winston`
* 전역적으로 사용하도록 설정
  * `main.ts` 에 로거 설정함으로써 부트스트랩 과정에서도 커스텀 로거 사용 가능
  * 로거를 사용할 모듈에서 `prioviders`에 추가

### Chapter 12. Exception Filter
* exception module
  * `npx @nestjs/cli g module exception`
  * 커스텀 예외 필터(`HttpExceptionFilter`) 정의
  * 전역으로 필터 사용을 위해, exception module 에 provider 설정
* 예외 필터 테스트를 위한 `InternalTestController` 추가

### Chapter 13. Interceptor
* 라우트 핸들러의 요청 처리 전후에 로깅 남기는 로깅 인터셉터

### Chapter 15. Health Check
* Terminus library
  * `npm i @nestjs/terminus`
  * `npm i @nestjs/axios`: HTTP Health Check 시 필요
* Health Check Controller