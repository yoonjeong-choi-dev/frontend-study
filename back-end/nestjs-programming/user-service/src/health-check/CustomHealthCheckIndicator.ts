import { Injectable } from '@nestjs/common';
import {
  HealthCheckError,
  HealthIndicator,
  HealthIndicatorResult,
} from '@nestjs/terminus';

export interface CustomEvent {
  name: string;
  type: string;
}

@Injectable()
export class CustomHealthCheckIndicator extends HealthIndicator {
  private events: CustomEvent[] = [
    { name: 'event1', type: 'OK' },
    { name: 'event2', type: 'BAD' },
  ];

  async check(key: string): Promise<HealthIndicatorResult> {
    const badEvents = this.events.filter((event) => event.type === 'BAD');
    const isHealthy = badEvents.length === 0;
    const result = this.getStatus(key, isHealthy, { badEvents });

    if (isHealthy) {
      return result;
    }

    throw new HealthCheckError('Event check failed', result);
  }
}
