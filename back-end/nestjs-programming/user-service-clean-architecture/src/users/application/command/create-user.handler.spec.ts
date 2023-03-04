import { UnprocessableEntityException } from '@nestjs/common';
import { Test } from '@nestjs/testing';
import { UserRepository } from 'src/users/infra/db/repository/UserRepository';
import * as ulid from 'ulid';
import * as uuid from 'uuid';
import { UserFactory } from '../../domain/user.factory';
import { CreateUserCommand } from './create-user.command';
import { CreateUserHandler } from './create-user.handler';

jest.mock('uuid');
jest.spyOn(uuid, 'v1').mockReturnValue('0000-0000-0000-0000');
jest.mock('ulid');
jest.spyOn(ulid, 'ulid').mockReturnValue('mock-ulid');

describe('CreateUserHandler', () => {
  let createUserHandler: CreateUserHandler;
  let userFactory: UserFactory;
  let userRepository: UserRepository;

  const id = ulid.ulid();
  const name = 'user_name';
  const email = 'user@example.com';
  const password = 'user_password';
  const signupVerifyToken = uuid.v1();

  beforeEach(async () => {
    const module = await Test.createTestingModule({
      providers: [
        CreateUserHandler,
        {
          provide: UserFactory,
          useValue: {
            create: jest.fn(),
          },
        },
        {
          provide: 'UserRepository',
          useValue: {
            save: jest.fn(),
          },
        },
      ],
    }).compile();

    createUserHandler = module.get(CreateUserHandler);
    userFactory = module.get(UserFactory);
    userRepository = module.get('UserRepository');
  });

  it('should execute CreateUserCommand', async () => {
    // Given : there is no such a user
    userRepository.save = jest.fn().mockReturnValue(true);
    userRepository.findByEmail = jest.fn().mockResolvedValue(null);

    // When
    await createUserHandler.execute(
      new CreateUserCommand(name, email, password),
    );

    // Then
    expect(userRepository.save).toBeCalledWith(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
    expect(userFactory.create).toBeCalledWith(
      id,
      name,
      email,
      password,
      signupVerifyToken,
    );
  });

  it('should throw UnprocessableEntityException when user exists', async () => {
    // Given : there is such a user
    userRepository.findByEmail = jest.fn().mockResolvedValue({
      id,
      name,
      email,
      password,
      signupVerifyToken,
    });

    // When - nothing
    // Then
    await expect(
      createUserHandler.execute(new CreateUserCommand(name, email, password)),
    ).rejects.toThrowError(UnprocessableEntityException);
  });
});
