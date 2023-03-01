import { EventsHandler, IEventHandler } from '@nestjs/cqrs';
import { UserCreatedEvent } from './user-created.event';
import { TestEvent } from './test.event';
import { EmailService } from '../../email/email.service';
import { CqrsEvent } from './cqrs-event';
import { Inject, Logger, LoggerService } from '@nestjs/common';

@EventsHandler(UserCreatedEvent, TestEvent)
export class UserEventsHandler implements IEventHandler<CqrsEvent> {
  constructor(
    private emailService: EmailService,
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
      case TestEvent.name:
        const { name } = event as TestEvent;
        this.logger.log(`TestEvent occurs with name - ${name}`);
        break;
      default:
        return;
    }
  }
}
