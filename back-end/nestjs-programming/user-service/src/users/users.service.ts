import * as uuid from 'uuid';
import { Injectable } from '@nestjs/common';
import { EmailService } from '../email/email.service';
import { UserInfo } from './UserInfo';

@Injectable()
export class UsersService {
  constructor(private emailService: EmailService) {}

  async createUser(name: string, email: string, password: string) {
    console.log('Create User');
    await this.checkUserExists(email);

    const signupVerifyToken = uuid.v1();
    await this.saveUser(name, email, password, signupVerifyToken);

    await this.sendMemberJoinEmail(email, signupVerifyToken);
  }

  private checkUserExists(email: string) {
    // TODO: connect with DB
    console.log('checkUserExists');
    return false;
  }

  private saveUser(
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ): void {
    // TODO: connect with DB
    console.log('saveUser');
    return;
  }

  private async sendMemberJoinEmail(email: string, signupVerifyToken: string) {
    await this.emailService.sendMemberJoinVerification(
      email,
      signupVerifyToken,
    );
  }

  async verifyEmail(signupVerifyToken: string): Promise<string> {
    // TODO: check user existence by token & publish token to sign in
    throw new Error('Not Implemented');
  }

  async signIn(email: string, password: string): Promise<string> {
    // TODO: check user existence by email, password & publish token to sign in
    throw new Error('Not Implemented');
  }

  async getUserInfo(userId: string): Promise<UserInfo> {
    // TODO: check user existence by userId
    return {
      id: userId,
      name: 'temp user',
      email: 'temp@example.com',
    };
  }
}
