import Mail from 'nodemailer/lib/mailer';
import * as nodemailer from 'nodemailer';

import { Inject, Injectable } from '@nestjs/common';
import emailConfig from '../config/emailConfig';
import { ConfigType } from '@nestjs/config';

interface EmailOptions {
  to: string;
  subject: string;
  html: string;
}

@Injectable()
export class EmailService {
  private transporter: Mail;

  constructor(
    @Inject(emailConfig.KEY) private config: ConfigType<typeof emailConfig>,
  ) {
    this.transporter = nodemailer.createTransport({
      service: config.service,
      auth: {
        user: config.auth.user,
        pass: config.auth.pass,
      },
    });
  }

  async sendMemberJoinVerification(
    emailAddress: string,
    signupVerifyToken: string,
  ) {
    // url for request to verify this email
    const baseUrl = this.config.baseUrl;
    const url = `${baseUrl}/users/email-verify?signupVerifyToken=${signupVerifyToken}`;

    const mailOptions: EmailOptions = {
      to: emailAddress,
      subject: 'Sign-up Email',
      html: `
        Click the 'Confirm' button then you can join the service.</br>
        <form action="${url}" method="POST">
          <button>Confirm</button>
          </form>
      `,
    };

    console.log('Send mail');
    return await this.transporter.sendMail(mailOptions);
  }
}
