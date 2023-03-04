import { Injectable } from '@nestjs/common';
import { DataSource, Repository } from 'typeorm';
import { InjectRepository } from '@nestjs/typeorm';
import { UserEntity } from '../entity/user.entity';

import { IUserRepository } from '../../../domain/repository/iuser.repository';
import { User } from '../../../domain/user';
import { UserFactory } from '../../../domain/user.factory';

@Injectable()
export class UserRepository implements IUserRepository {
  constructor(
    @InjectRepository(UserEntity)
    private usersRepository: Repository<UserEntity>,
    private datasource: DataSource,
    private userFactory: UserFactory,
  ) {}

  async findById(id: string): Promise<User | null> {
    const entity = await this.usersRepository.findOne({
      where: { id },
    });

    if (!entity) {
      return null;
    }

    const { email, name, signupVerifyToken, password } = entity;
    return this.userFactory.reconstitute(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
  }

  async findByEmail(email: string): Promise<User | null> {
    const entity = await this.usersRepository.findOne({
      where: { email },
    });

    if (!entity) {
      return null;
    }

    const { id, name, signupVerifyToken, password } = entity;
    return this.userFactory.reconstitute(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
  }

  async findByEmailAndPassword(
    email: string,
    password: string,
  ): Promise<User | null> {
    const entity = await this.usersRepository.findOne({
      where: { email, password },
    });

    if (!entity) {
      return null;
    }

    const { id, name, signupVerifyToken } = entity;
    return this.userFactory.reconstitute(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
  }

  async findBySignupVerifyToken(
    signupVerifyToken: string,
  ): Promise<User | null> {
    const entity = await this.usersRepository.findOne({
      where: { signupVerifyToken },
    });

    if (!entity) {
      return null;
    }

    const { id, name, email, password } = entity;
    return this.userFactory.reconstitute(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
  }

  async save(
    id: string,
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ): Promise<boolean> {
    const queryRunner = this.datasource.createQueryRunner();

    // connect and start transaction
    await queryRunner.connect();
    await queryRunner.startTransaction();
    let isSuccess = true;
    try {
      // saveUser
      const user = new UserEntity();
      user.id = id;
      user.name = name;
      user.email = email;
      user.password = password;
      user.signupVerifyToken = signupVerifyToken;

      // save the entity to user repository
      await queryRunner.manager.save(user);

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
