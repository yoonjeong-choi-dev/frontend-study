namespace Ch8EventEmitter {
  type ConnectEvent = {
    ready: void,
    error: Error,
    reconnecting: {
      attempt: number,
      delay: number,
    },
  };

  // ConnectEvent[E] : 특정 이벤트를 가리킴
  // => args 는 특정 이벤트에 대한 정보
  type RedisClient = {
    on<E extends keyof ConnectEvent>(
      event: E,
      f: (arg: ConnectEvent[E]) => void
    ): void,
    emit<E extends keyof ConnectEvent>(
      event: E,
      args: ConnectEvent[E]
    ): void
  }

  function useRedisClientExample(client: RedisClient) {
    client.on('ready', () => console.log('Client is ready'));
    client.on('error', (e) => console.error('An error occurred :', e.message));
    client.on('reconnecting', params => {
      console.log(`Reconnection... ${params.attempt} times with ${params.delay}(s)`);
    });
  }

}
