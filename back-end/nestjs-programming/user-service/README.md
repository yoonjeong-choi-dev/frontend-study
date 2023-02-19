# Main Project - User Service

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


