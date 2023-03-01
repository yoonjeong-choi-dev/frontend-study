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
import AuthGuard from '../auth/auth.guard';
import { CreateUserCommand } from './command/create-user.command';
import { SignInCommand } from './command/sign-in.command';
import { VerifyEmailCommand } from './command/verify-email.command';
import { GetUserInfoQuery } from './query/get-user-info.query';
// import { UsersService } from './users.service';

@Controller('users')
export class UsersController {
  // IOC for services
  constructor(
    // private usersService: UsersService,
    // @Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: WinstonLogger,
    @Inject(Logger) private readonly logger: LoggerService,
    private commandBus: CommandBus,
    private queryBus: QueryBus,
  ) {}

  @Post()
  async createUser(@Body() dto: CreateUserDto): Promise<void> {
    // console.log('createUser', dto);
    this.logger.log(`createUser - ${JSON.stringify(dto)}`);

    const { name, email, password } = dto;
    //await this.usersService.createUser(name, email, password);
    const command = new CreateUserCommand(name, email, password);
    return this.commandBus.execute(command);
  }

  @Post('/email-verify')
  async verifyEmail(@Query() dto: VerifyEmailDto): Promise<string> {
    // console.log('verifyEmail', dto);
    this.logger.log(`verifyEmail - ${JSON.stringify(dto)}`);

    const { signupVerifyToken } = dto;
    // return await this.usersService.verifyEmail(signupVerifyToken);
    const command = new VerifyEmailCommand(signupVerifyToken);
    return await this.commandBus.execute(command);
  }

  @Post('/sign-in')
  async signIn(@Body() dto: UserLoginDto): Promise<string> {
    // console.log('signIn', dto);
    this.logger.log(`signIn - ${JSON.stringify(dto)}`);

    const { email, password } = dto;
    // return await this.usersService.signIn(email, password);
    const command = new SignInCommand(email, password);
    return await this.commandBus.execute(command);
  }

  @UseGuards(AuthGuard)
  @Get('/:id')
  async getUserInfo(
    @Headers() headers: any,
    @Param('id') userId: string,
  ): Promise<UserInfo> {
    // console.log('getUserInfo', userId);
    this.logger.log(`getUserInfo - ${userId}`);

    // return await this.usersService.getUserInfo(userId);
    const query = new GetUserInfoQuery(userId);
    return await this.queryBus.execute(query);
  }

  // @Get('/:id')
  // async getUserInfo(
  //   @Headers() headers: any,
  //   @Param('id') userId: string,
  // ): Promise<UserInfo> {
  //   console.log('getUserInfo', userId);
  //
  //   const jwtString = headers.authorization.split('Bearer ')[1];
  //
  //   this.authService.verify(jwtString);
  //
  //   return await this.usersService.getUserInfo(userId);
  // }
}
