import {
  Controller,
  Get,
  InternalServerErrorException,
  UseInterceptors,
} from '@nestjs/common';
import { ErrorsInterceptor } from './errors.interceptor';

@Controller('internal-test')
export class InternalTestController {
  @Get('/filter-test')
  makeError(test: any) {
    return test.noMethod();
  }

  @UseInterceptors(ErrorsInterceptor)
  @Get('/intercept-test')
  interceptTest() {
    throw new InternalServerErrorException();
  }
}
