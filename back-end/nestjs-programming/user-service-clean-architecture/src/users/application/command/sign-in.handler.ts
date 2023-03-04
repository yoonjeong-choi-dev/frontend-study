import { Inject, Injectable, NotFoundException } from '@nestjs/common';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { InjectRepository } from '@nestjs/typeorm';
import { IAuthService } from 'src/users/application/adapter/IAuthService';
import { Repository } from 'typeorm';
import { UserEntity } from '../../infra/db/entity/user.entity';
import { SignInCommand } from './sign-in.command';

@Injectable()
@CommandHandler(SignInCommand)
export class SignInHandler implements ICommandHandler<SignInCommand> {
  constructor(
    @InjectRepository(UserEntity)
    private usersRepository: Repository<UserEntity>,
    @Inject('AuthService')
    private authService: IAuthService,
  ) {}

  // UsersService.signIn
  async execute(command: SignInCommand): Promise<string> {
    const { email, password } = command;
    const user = await this.usersRepository.findOne({
      where: { email, password },
    });

    if (!user) {
      throw new NotFoundException('there is no such a user');
    }

    return this.authService.signIn(user.id, user.name, user.email);
  }
}
