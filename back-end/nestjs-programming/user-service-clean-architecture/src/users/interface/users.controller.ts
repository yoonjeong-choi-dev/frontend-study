import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Query,
  Headers,
  UseGuards,
  Inject,
  Logger,
  LoggerService,
} from '@nestjs/common';
import { CommandBus, QueryBus } from '@nestjs/cqrs';
import { CreateUserDto } from './dto/create-user.dto';
import { VerifyEmailDto } from './dto/verify-email.dto';
import { UserLoginDto } from './dto/user-login.dto';
import { UserInfo } from './UserInfo';
import AuthGuard from '../../auth/auth.guard';
import { CreateUserCommand } from '../application/command/create-user.command';
import { SignInCommand } from '../application/command/sign-in.command';
import { VerifyEmailCommand } from '../application/command/verify-email.command';
import { GetUserInfoQuery } from '../application/query/get-user-info.query';

@Controller('users')
export class UsersController {
  // IOC for services
  constructor(
    @Inject(Logger) private readonly logger: LoggerService,
    private commandBus: CommandBus,
    private queryBus: QueryBus,
  ) {}

  @Post()
  async createUser(@Body() dto: CreateUserDto): Promise<void> {
    this.logger.log(`createUser - ${JSON.stringify(dto)}`);

    const { name, email, password } = dto;
    const command = new CreateUserCommand(name, email, password);
    return this.commandBus.execute(command);
  }

  @Post('/email-verify')
  async verifyEmail(@Query() dto: VerifyEmailDto): Promise<string> {
    this.logger.log(`verifyEmail - ${JSON.stringify(dto)}`);

    const { signupVerifyToken } = dto;
    const command = new VerifyEmailCommand(signupVerifyToken);
    return await this.commandBus.execute(command);
  }

  @Post('/sign-in')
  async signIn(@Body() dto: UserLoginDto): Promise<string> {
    this.logger.log(`signIn - ${JSON.stringify(dto)}`);

    const { email, password } = dto;
    const command = new SignInCommand(email, password);
    return await this.commandBus.execute(command);
  }

  @UseGuards(AuthGuard)
  @Get('/:id')
  async getUserInfo(
    @Headers() headers: any,
    @Param('id') userId: string,
  ): Promise<UserInfo> {
    this.logger.log(`getUserInfo - ${userId}`);

    const query = new GetUserInfoQuery(userId);
    return await this.queryBus.execute(query);
  }
}
