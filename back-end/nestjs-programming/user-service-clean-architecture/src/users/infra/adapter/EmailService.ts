import { EmailService as ExternalEmailService } from 'src/email/email.service';
import { Injectable } from '@nestjs/common';
import { IEmailService } from 'src/users/application/adapter/IEmailService';

@Injectable()
export class EmailService implements IEmailService {
  constructor(private externalEmailService: ExternalEmailService) {}

  sendMemberJoinVerification(email, signupVerifyToken): Promise<void> {
    return this.externalEmailService.sendMemberJoinVerification(
      email,
      signupVerifyToken,
    );
  }
}
