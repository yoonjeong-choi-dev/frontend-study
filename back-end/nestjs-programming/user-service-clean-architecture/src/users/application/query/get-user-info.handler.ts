import { Inject, NotFoundException } from '@nestjs/common';
import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { GetUserInfoQuery } from './get-user-info.query';
import { IUserRepository } from 'src/users/domain/repository/iuser.repository';
import { UserInfo } from '../../interface/UserInfo';

@QueryHandler(GetUserInfoQuery)
export class GetUserInfoHandler implements IQueryHandler<GetUserInfoQuery> {
  constructor(
    @Inject('UserRepository')
    private userRepository: IUserRepository,
  ) {}

  // UsersService.getUserInfo
  async execute(query: GetUserInfoQuery): Promise<UserInfo> {
    const { userId } = query;
    const user = await this.userRepository.findById(userId);

    if (!user) {
      throw new NotFoundException('there is no such a user');
    }

    return {
      id: user.id,
      name: user.name,
      email: user.email,
    };
  }
}
