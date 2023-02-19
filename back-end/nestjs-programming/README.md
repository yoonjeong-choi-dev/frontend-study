# 교재 정보
- NestJS로 배우는 백엔드 프로그래밍(저자:한용재, 출판:제이펍)
- Spring가 비슷한 형태를 가지는 NestJS가 궁금하여 시작
- 리프레시 용

# 파일 정보
- 주의 사항
  - `npx @nestjs/cli` 사용해야 함
  - KT Issue
    - 현재 KT 관련해서 npm 이슈 존재([reference](https://velog.io/@librarian/ts-jest-%EC%84%A4%EC%B9%98-%EC%95%88%EB%90%98%EB%8A%94-%ED%98%84%EC%83%81))
    - `npm config set registry https://registry.npmjs.cf/` 
    - 복구: `npm config set registry https://registry.npmjs.org/`
- exmaples
  - 각 챕터 별 예제 코드
  - 챕터 단위로 라우팅하여 구현 예정
- user-service
  - 교재 전반적으로 구현하는 유저 서비스
  - 회원 가입, 로그인, 유저 정보 조회 기능 구현
