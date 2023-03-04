import { User } from '../user';

export interface IUserRepository {
  findById: (id: string) => Promise<User | null>;
  findByEmail: (email: string) => Promise<User | null>;
  findByEmailAndPassword: (
    email: string,
    password: string,
  ) => Promise<User | null>;
  findBySignupVerifyToken: (signupVerifyToken: string) => Promise<User | null>;
  save: (
    id: string,
    name: string,
    email: string,
    password: string,
    signupVerifyToken: string,
  ) => Promise<boolean>;
}
