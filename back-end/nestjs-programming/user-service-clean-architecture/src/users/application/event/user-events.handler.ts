import { EventsHandler, IEventHandler } from '@nestjs/cqrs';
import { IEmailService } from 'src/users/application/adapter/IEmailService';
import { UserCreatedEvent } from '../../domain/user-created.event';
import { CqrsEvent } from '../../domain/cqrs-event';
import { Inject, Logger, LoggerService } from '@nestjs/common';

@EventsHandler(UserCreatedEvent)
export class UserEventsHandler implements IEventHandler<CqrsEvent> {
  constructor(
    @Inject('EmailService') private emailService: IEmailService,
    @Inject(Logger) private readonly logger: LoggerService,
  ) {}

  async handle(event: CqrsEvent) {
    switch (event.name) {
      case UserCreatedEvent.name:
        const { email, signupVerifyToken } = event as UserCreatedEvent;
        this.logger.log(`UserCreatedEvent occurs with email - ${email}`);
        console.log(`UserCreatedEvent occurs with email - ${email}`);
        await this.emailService.sendMemberJoinVerification(
          email,
          signupVerifyToken,
        );
        break;
      default:
        return;
    }
  }
}
