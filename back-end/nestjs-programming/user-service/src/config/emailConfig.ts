import { registerAs } from '@nestjs/config';

// 'email' 토큰으로 ConfigFactory 등록
// => ConfigModule 에서 동적으로 등록
// => 해당 설정을 사용하려면, @Inject(emailConfig.KEY) 를 이용하여 의존성 주입
export default registerAs('email', () => ({
  service: process.env.EMAIL_SERVICE,
  auth: {
    user: process.env.EMAIL_AUTH_USER,
    pass: process.env.EMAIL_AUTH_PASSWORD,
  },
  baseUrl: process.env.EMAIL_BASE_URL,
}));
