import { Injectable } from '@nestjs/common';
import { AuthService as ExternalAuthService } from 'src/auth/auth.service';
import ExternalAuthUser from 'src/auth/AuthUser';
import { IAuthService } from 'src/users/application/adapter/IAuthService';

@Injectable()
export class AuthService implements IAuthService {
  constructor(private externalAuthService: ExternalAuthService) {}

  signIn(id: string, name: string, email: string): string {
    const authUser: ExternalAuthUser = {
      id,
      name,
      email,
    };
    return this.externalAuthService.signIn(authUser);
  }
}
