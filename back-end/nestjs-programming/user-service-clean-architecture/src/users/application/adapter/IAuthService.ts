export class IAuthService {
  signIn: (id: string, name: string, email: string) => string;
}
