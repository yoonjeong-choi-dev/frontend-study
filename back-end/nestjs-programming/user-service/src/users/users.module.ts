import { Logger, Module } from '@nestjs/common';
import { TypeOrmModule } from '@nestjs/typeorm';
import { CqrsModule } from '@nestjs/cqrs';

import { UsersController } from './users.controller';
// import { UsersService } from './users.service';
import { EmailModule } from '../email/email.module';
import { UserEntity } from './entity/user.entity';
import { AuthModule } from '../auth/auth.module';
import { CreateUserHandler } from './command/create-user.handler';
import { VerifyEmailHandler } from './command/verify-email.handler';
import { SignInHandler } from './command/sign-in.handler';
import { UserEventsHandler } from './event/user-events.handler';
import { GetUserInfoHandler } from './query/get-user-info.handler';

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
    // UsersService,
    Logger,
    ...queryHandlers,
    ...commandHandlers,
    ...eventHandlers,
  ],
})
export class UsersModule {}
