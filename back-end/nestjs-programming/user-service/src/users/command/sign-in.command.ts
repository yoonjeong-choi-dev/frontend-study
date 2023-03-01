import { ICommand } from '@nestjs/cqrs';

export class SignInCommand implements ICommand {
  constructor(readonly email: string, readonly password: string) {}
}
