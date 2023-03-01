import { CqrsEvent } from './cqrs-event';
import { IEvent } from '@nestjs/cqrs';

export class UserCreatedEvent extends CqrsEvent implements IEvent {
  constructor(readonly email: string, readonly signupVerifyToken: string) {
    super(UserCreatedEvent.name);
  }
}
