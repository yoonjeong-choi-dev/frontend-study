import { Inject, Injectable, NotFoundException } from '@nestjs/common';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { IAuthService } from 'src/users/application/adapter/IAuthService';
import { IUserRepository } from 'src/users/domain/repository/iuser.repository';
import { VerifyEmailCommand } from './verify-email.command';

@Injectable()
@CommandHandler(VerifyEmailCommand)
export class VerifyEmailHandler implements ICommandHandler<VerifyEmailCommand> {
  constructor(
    @Inject('UserRepository')
    private userRepository: IUserRepository,
    @Inject('AuthService')
    private authService: IAuthService,
  ) {}

  // UsersService.verifyEmail
  async execute(command: VerifyEmailCommand): Promise<string> {
    const { signupVerifyToken } = command;
    const user = await this.userRepository.findBySignupVerifyToken(
      signupVerifyToken,
    );

    if (!user) {
      throw new NotFoundException('there is no such a user');
    }

    return this.authService.signIn(user.id, user.name, user.email);
  }
}
