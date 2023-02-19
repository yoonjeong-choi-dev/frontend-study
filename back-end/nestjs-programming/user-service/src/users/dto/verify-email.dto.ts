import { IsString } from 'class-validator';

export class VerifyEmailDto {
  @IsString()
  signupVerifyToken: string;
}
