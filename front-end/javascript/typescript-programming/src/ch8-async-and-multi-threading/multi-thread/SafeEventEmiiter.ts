import EventEmitter from 'events';

export default class SafeEmitter<Events extends Record<string | symbol, unknown[]>> {
  private emitter = new EventEmitter();

  emit<K extends keyof Events>(channel: K, ...data: Events[K]) {
    // @ts-ignore
    return this.emitter.emit(channel, ...data);
  }

  on<K extends keyof Events>(channel: K, listener: (...data: Events[K]) => void) {
    // @ts-ignore
    return this.emitter.on(channel, listener);
  }
}
