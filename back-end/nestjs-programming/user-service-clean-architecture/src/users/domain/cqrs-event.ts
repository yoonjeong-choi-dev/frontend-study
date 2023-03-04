export abstract class CqrsEvent {
  protected constructor(readonly name: string) {}
}
