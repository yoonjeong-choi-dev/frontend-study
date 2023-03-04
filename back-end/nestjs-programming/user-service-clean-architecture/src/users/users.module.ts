import { Logger, Module } from '@nestjs/common';
import { CqrsModule } from '@nestjs/cqrs';
import { TypeOrmModule } from '@nestjs/typeorm';
import { AuthModule } from 'src//auth/auth.module';
import { EmailModule } from 'src/email/email.module';
import { CreateUserHandler } from 'src/users/application/command/create-user.handler';
import { SignInHandler } from 'src/users/application/command/sign-in.handler';
import { VerifyEmailHandler } from 'src/users/application/command/verify-email.handler';
import { UserEventsHandler } from 'src/users/application/event/user-events.handler';
import { GetUserInfoHandler } from 'src/users/application/query/get-user-info.handler';
import { UserFactory } from 'src/users/domain/user.factory';
import { AuthService } from 'src/users/infra/adapter/AuthService';
import { EmailService } from 'src/users/infra/adapter/EmailService';
import { UserEntity } from 'src/users/infra/db/entity/user.entity';
import { UserRepository } from 'src/users/infra/db/repository/UserRepository';

import { UsersController } from 'src/users/interface/users.controller';

const factories = [UserFactory];

const repositories = [
  {
    provide: 'UserRepository',
    useClass: UserRepository,
  },
  {
    provide: 'EmailService',
    useClass: EmailService,
  },
  {
    provide: 'AuthService',
    useClass: AuthService,
  },
];

const queryHandlers = [GetUserInfoHandler];
const commandHandlers = [CreateUserHandler, VerifyEmailHandler, SignInHandler];
const eventHandlers = [UserEventsHandler];

@Module({
  imports: [
    TypeOrmModule.forFeature([UserEntity]),
    EmailModule,
    AuthModule,
    CqrsModule,
  ],
  controllers: [UsersController],
  providers: [
    Logger,
    ...repositories,
    ...factories,
    ...queryHandlers,
    ...commandHandlers,
    ...eventHandlers,
  ],
})
export class UsersModule {}
