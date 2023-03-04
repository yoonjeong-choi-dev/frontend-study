export class User {
  constructor(
    private _id: string,
    private _name: string,
    private _email: string,
    private password: string,
    private signupVerifyToken: string,
  ) {}

  get id(): string {
    return this._id;
  }

  get name(): string {
    return this._name;
  }

  get email(): string {
    return this._email;
  }
}
