import { ulid } from 'ulid';
import * as uuid from 'uuid';
import {
  Inject,
  Injectable,
  InternalServerErrorException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { UserFactory } from '../../domain/user.factory';
import { IUserRepository } from '../../domain/repository/iuser.repository';
import { CreateUserCommand } from './create-user.command';

@Injectable()
@CommandHandler(CreateUserCommand)
export class CreateUserHandler implements ICommandHandler<CreateUserCommand> {
  constructor(
    private userFactory: UserFactory,
    @Inject('UserRepository')
    private userRepository: IUserRepository,
  ) {}

  // UsersService.createUser
  async execute(command: CreateUserCommand): Promise<any> {
    const { name, email, password } = command;

    const userExist = await this.checkUserExists(email);
    if (userExist) {
      throw new UnprocessableEntityException(
        'Already registered by this email',
      );
    }

    const id = ulid();
    const signupVerifyToken = uuid.v1();
    const isSaved = await this.userRepository.save(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );

    if (!isSaved) {
      throw new InternalServerErrorException('Database Save Error');
    }

    this.userFactory.create(id, name, email, password, signupVerifyToken);
  }

  private async checkUserExists(email: string): Promise<boolean> {
    const user = await this.userRepository.findByEmail(email);
    return user !== null;
  }
}
