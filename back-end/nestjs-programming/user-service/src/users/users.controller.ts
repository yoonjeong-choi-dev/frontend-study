import {
  Body,
  Controller,
  Get,
  Param,
  Post,
  Query,
  Headers,
  UseGuards,
} from '@nestjs/common';
import { CreateUserDto } from './dto/create-user.dto';
import { VerifyEmailDto } from './dto/verify-email.dto';
import { UserLoginDto } from './dto/user-login.dto';
import { UserInfo } from './UserInfo';
import { UsersService } from './users.service';
import { AuthService } from '../auth/auth.service';
import AuthGuard from '../auth/auth.guard';

@Controller('users')
export class UsersController {
  // IOC for services
  constructor(
    private usersService: UsersService,
    private authService: AuthService,
  ) {}

  @Post()
  async createUser(@Body() dto: CreateUserDto): Promise<void> {
    console.log('createUser', dto);
    const { name, email, password } = dto;
    await this.usersService.createUser(name, email, password);
  }

  @Post('/email-verify')
  async verifyEmail(@Query() dto: VerifyEmailDto): Promise<string> {
    console.log('verifyEmail', dto);
    const { signupVerifyToken } = dto;
    return await this.usersService.verifyEmail(signupVerifyToken);
  }

  @Post('/sign-in')
  async signIn(@Body() dto: UserLoginDto): Promise<string> {
    console.log('signIn', dto);
    const { email, password } = dto;
    return await this.usersService.signIn(email, password);
  }

  @UseGuards(AuthGuard)
  @Get('/:id')
  async getUserInfo(
    @Headers() headers: any,
    @Param('id') userId: string,
  ): Promise<UserInfo> {
    console.log('getUserInfo', userId);
    return await this.usersService.getUserInfo(userId);
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
