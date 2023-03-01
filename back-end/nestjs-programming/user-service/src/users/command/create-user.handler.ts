import {
  Injectable,
  InternalServerErrorException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { CommandHandler, EventBus, ICommandHandler } from '@nestjs/cqrs';
import { CreateUserCommand } from './create-user.command';
import { InjectRepository } from '@nestjs/typeorm';
import { UserEntity } from '../entity/user.entity';
import { DataSource, Repository } from 'typeorm';
import { ulid } from 'ulid';
import * as uuid from 'uuid';
import { UserCreatedEvent } from '../event/user-created.event';
import { TestEvent } from '../event/test.event';

@Injectable()
@CommandHandler(CreateUserCommand)
export class CreateUserHandler implements ICommandHandler<CreateUserCommand> {
  constructor(
    @InjectRepository(UserEntity)
    private usersRepository: Repository<UserEntity>,
    private datasource: DataSource,
    private eventBus: EventBus,
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

    const signupVerifyToken = uuid.v1();
    const isSaved = await this.saveUserUsingQueryRunner(
      name,
      email,
      password,
      signupVerifyToken,
    );

    if (!isSaved) {
      throw new InternalServerErrorException('Database Save Error');
    }

    this.eventBus.publish(new UserCreatedEvent(email, signupVerifyToken));
    this.eventBus.publish(new TestEvent(email));
  }

  private async checkUserExists(email: string): Promise<boolean> {
    const user = await this.usersRepository.findOne({
      where: { email },
    });
    return user !== null;
  }

  private async saveUserUsingQueryRunner(
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ) {
    const queryRunner = this.datasource.createQueryRunner();

    // connect and start transaction
    await queryRunner.connect();
    await queryRunner.startTransaction();
    let isSuccess = true;
    try {
      // saveUser
      const user = new UserEntity();
      user.id = ulid();
      user.name = name;
      user.email = email;
      user.password = password;
      user.signupVerifyToken = signupVerifyToken;

      // save the entity to user repository
      await queryRunner.manager.save(user);

      // transaction test
      //throw new Error();

      console.log('id', user.id);

      // commit
      await queryRunner.commitTransaction();
    } catch (e) {
      // rollback
      await queryRunner.rollbackTransaction();
      isSuccess = false;
    } finally {
      await queryRunner.release();
    }

    return isSuccess;
  }
}
