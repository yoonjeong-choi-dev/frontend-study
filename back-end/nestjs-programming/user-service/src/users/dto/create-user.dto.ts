import {
  IsEmail,
  IsString,
  Matches,
  MaxLength,
  MinLength,
} from 'class-validator';
import { Transform } from 'class-transformer';
import { NotIn } from '../../utils/decorators/not-in';

export class CreateUserDto {
  @Transform((params) => params.value.trim())
  @IsString()
  @MinLength(2)
  @MaxLength(30)
  @NotIn('password', { message: 'password cannot contain name' })
  readonly name: string;

  @Transform((params) => params.value.trim())
  @IsString()
  @IsEmail()
  @MaxLength(60)
  readonly email: string;

  @IsString()
  @Matches(/^[A-Za-z\d!@#$%^&*()]{8,30}$/, {
    message: 'password must consist of alphabets, numbers and !@#$%^&*()',
  })
  readonly password: string;
}
