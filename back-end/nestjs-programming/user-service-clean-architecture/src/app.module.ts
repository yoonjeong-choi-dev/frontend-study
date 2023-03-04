import { Module } from '@nestjs/common';
import { UsersModule } from './users/users.module';
import { ConfigModule } from '@nestjs/config';
import emailConfig from './config/emailConfig';
import { envValidationSchema } from './config/validationSchema';
import { TypeOrmModule } from '@nestjs/typeorm';
import authConfig from './config/authConfig';
import { ExceptionModule } from './exception/exception.module';
import { InternalTestController } from './internel-test/internal-test.controller';
import { LoggingModule } from './logging/logging.module';
import { HealthCheckController } from './health-check/health-check.controller';
import { TerminusModule } from '@nestjs/terminus';
import { HttpModule } from '@nestjs/axios';
import { CustomHealthCheckIndicator } from './health-check/CustomHealthCheckIndicator';

@Module({
  imports: [
    UsersModule,
    ConfigModule.forRoot({
      envFilePath: [`.${process.env.NODE_ENV}.env`], // 현재 환경에 대한 환경 변수 파일
      load: [emailConfig, authConfig], // ConfigFactory 등록
      isGlobal: true, // 환경 설정에 관련된 모듈이므로 전역으로 사용
      validationSchema: envValidationSchema,
    }),
    TypeOrmModule.forRoot({
      type: 'mysql',
      host: process.env.DATABASE_HOST,
      port: 3306,
      username: process.env.DATABASE_USERNAME,
      password: process.env.DATABASE_PASSWORD,
      database: process.env.DATABASE_NAME,
      entities: [__dirname + '/**/*.entity{.ts,.js}'],
      synchronize: process.env.DATABASE_SYNCHRONIZE === 'true',
      migrations: [__dirname + '/**/migrations/*.js'],
      migrationsTableName: 'migrations',
    }),
    ExceptionModule,
    LoggingModule,
    TerminusModule,
    HttpModule,
  ],
  controllers: [InternalTestController, HealthCheckController],
  providers: [CustomHealthCheckIndicator],
})
export class AppModule {}
