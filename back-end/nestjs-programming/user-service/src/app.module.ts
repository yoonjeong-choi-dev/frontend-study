import { Module } from '@nestjs/common';
import { UsersModule } from './users/users.module';
import { ConfigModule } from '@nestjs/config';
import emailConfig from './config/emailConfig';
import { envValidationSchema } from './config/validationSchema';

@Module({
  imports: [
    UsersModule,
    ConfigModule.forRoot({
      envFilePath: [`${__dirname}/config/env/.${process.env.NODE_ENV}.env`], // 현재 환경에 대한 환경 변수 파일
      load: [emailConfig], // ConfigFactory 등록
      isGlobal: true, // 환경 설정에 관련된 모듈이므로 전역으로 사용
      validationSchema: envValidationSchema,
    }),
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
