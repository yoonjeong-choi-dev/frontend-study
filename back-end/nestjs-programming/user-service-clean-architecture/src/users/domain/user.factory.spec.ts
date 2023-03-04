import { EventBus } from '@nestjs/cqrs';
import { Test } from '@nestjs/testing';
import { User } from './user';
import { UserFactory } from './user.factory';

describe('UserFactory', () => {
  let userFactory: UserFactory;
  let eventBus: jest.Mocked<EventBus>;

  // beforeAll 로 하는 경우, EventBus.publish 호출횟수가 초기화되지 않음
  beforeEach(async () => {
    const module = await Test.createTestingModule({
      providers: [
        UserFactory,
        {
          provide: EventBus,
          useValue: {
            publish: jest.fn(),
          },
        },
      ],
    }).compile();

    userFactory = module.get(UserFactory);
    eventBus = module.get(EventBus);
  });

  describe('create', () => {
    it('should create user', () => {
      // Given - empty
      // When
      const user = userFactory.create(
        'user_id',
        'user_name',
        'user@example.com',
        'user_password',
        'verify-token-example',
      );

      // Then
      const expected = new User(
        'user_id',
        'user_name',
        'user@example.com',
        'user_password',
        'verify-token-example',
      );

      expect(expected).toEqual(user);
      expect(eventBus.publish).toBeCalledTimes(1);
    });
  });

  describe('reconstitute', () => {
    it('should reconstitute user', () => {
      // Given - empty
      // When
      const user = userFactory.reconstitute(
        'user_id',
        'user_name',
        'user@example.com',
        'user_password',
        'verify-token-example',
      );

      // Then
      const expected = new User(
        'user_id',
        'user_name',
        'user@example.com',
        'user_password',
        'verify-token-example',
      );

      expect(expected).toEqual(user);
      expect(eventBus.publish).toBeCalledTimes(0);
    });
  });
});
