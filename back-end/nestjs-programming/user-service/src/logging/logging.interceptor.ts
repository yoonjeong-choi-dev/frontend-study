import {
  CallHandler,
  ExecutionContext,
  Injectable,
  Logger,
  NestInterceptor,
} from '@nestjs/common';
import { Observable } from 'rxjs';
import { tap } from 'rxjs/operators';

@Injectable()
export class LoggingInterceptor implements NestInterceptor {
  constructor(private readonly logger: Logger) {}

  intercept(
    context: ExecutionContext,
    next: CallHandler<any>,
  ): Observable<any> | Promise<Observable<any>> {
    const { method, url, body } = context.getArgByIndex(0);
    this.logger.log(`Request to ${method} via ${url}`);

    const start = Date.now();

    return next.handle().pipe(
      tap((data) =>
        this.logger.log(`Response from ${method} via ${url}
          response: ${JSON.stringify(data)}
          duration(ms): ${Date.now() - start}`),
      ),
    );
  }
}
