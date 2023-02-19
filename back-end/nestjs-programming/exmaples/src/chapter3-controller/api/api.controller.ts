import { Controller, Get } from '@nestjs/common';

@Controller({
  host: 'api.yj',
  path: 'chapter3-controller',
})
export class ApiController {
  // GET api.yj:7166/chapter3-controller
  @Get()
  apiTest() {
    return 'This is a API host test';
  }
}
