import * as uuid from 'uuid';
import {
  Injectable,
  InternalServerErrorException,
  UnprocessableEntityException,
} from '@nestjs/common';
import { EmailService } from '../email/email.service';
import { UserInfo } from './UserInfo';
import { UserEntity } from './entity/user.entity';
import { DataSource, Repository } from 'typeorm';
import { ulid } from 'ulid';
import { InjectRepository } from '@nestjs/typeorm';

@Injectable()
export class UsersService {
  constructor(
    private emailService: EmailService,
    @InjectRepository(UserEntity)
    private usersRepository: Repository<UserEntity>,
    private datasource: DataSource,
  ) {}

  async createUser(name: string, email: string, password: string) {
    const userExist = await this.checkUserExists(email);
    if (userExist) {
      throw new UnprocessableEntityException(
        'Already registered by this email',
      );
    }

    const signupVerifyToken = uuid.v1();
    // await this.saveUser(name, email, password, signupVerifyToken);
    const isSaved = await this.saveUserUsingQueryRunner(
      name,
      email,
      password,
      signupVerifyToken,
    );

    if (!isSaved) {
      throw new InternalServerErrorException('Database Save Error');
    }

    await this.sendMemberJoinEmail(email, signupVerifyToken);
  }

  private async checkUserExists(email: string): Promise<boolean> {
    const user = await this.usersRepository.findOne({
      where: { email },
    });
    return user !== null;
  }

  // no transaction
  private async saveUser(
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ) {
    const user = new UserEntity();
    user.id = ulid();
    user.name = name;
    user.email = email;
    user.password = password;
    user.signupVerifyToken = signupVerifyToken;

    // save the entity to user repository
    await this.usersRepository.save(user);
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

  private async saveUserUsingTransaction(
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ) {
    await this.datasource.transaction(async (manager) => {
      // saveUser
      const user = new UserEntity();
      user.id = ulid();
      user.name = name;
      user.email = email;
      user.password = password;
      user.signupVerifyToken = signupVerifyToken;

      // save the entity to user repository
      await manager.save(user);
    });
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
