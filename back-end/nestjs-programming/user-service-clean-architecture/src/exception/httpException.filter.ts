import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  InternalServerErrorException,
  Logger,
} from '@nestjs/common';
import { Response, Request } from 'express';

@Catch()
export class HttpExceptionFilter implements ExceptionFilter {
  constructor(private logger: Logger) {}

  catch(exception: any, host: ArgumentsHost): any {
    const ctx = host.switchToHttp();
    const res = ctx.getResponse<Response>();
    const req = ctx.getRequest<Request>();
    const stack = exception.stack;

    // HttpException 은 nest 에서 처리
    if (!(exception instanceof HttpException)) {
      exception = new InternalServerErrorException();
    }

    const response = (exception as HttpException).getResponse();
    this.logger.log(this.createLog(req, response, stack));

    res.status((exception as HttpException).getStatus()).json(response);
  }

  private createLog(req: Request, res: object | string, stack: any) {
    return {
      timestamp: new Date(),
      url: req.url,
      response: res,
      stack,
    };
  }
}
