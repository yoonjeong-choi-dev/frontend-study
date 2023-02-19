import {
  BadRequestException,
  Controller,
  Get,
  Header,
  HttpCode,
  Param,
  Patch,
  Put,
  Req,
  Res,
} from '@nestjs/common';
import { Request, Response } from 'express';

@Controller('chapter3-controller')
export class Chapter3MainController {
  @Get()
  getIndex(): string {
    return '<h1>This is index of Chapter 3</h1>';
  }

  @Get('/hi')
  sayHi(): string {
    return 'Hi~';
  }

  @Get('/req')
  getRequest(@Req() req: Request): string {
    console.log(req);
    return `This is your request from "${req.url}"`;
  }

  @Header('Custom', 'from @Header')
  @Get('/header')
  getCustomHeader(@Res() res: Response) {
    res.header('Custom2', 'from Response Object');
    return res.status(200).send('Customized Response Object - See headers');
  }

  @HttpCode(202)
  @Patch('')
  updateTest(): string {
    return 'Resource Updated! - Check status code';
  }

  @Patch(':id')
  updateById(@Param('id') id: string) {
    return `User id(${id}) is updated`;
  }

  @Put(':id')
  updateBadRequest(@Param('id') id: string) {
    if (parseInt(id) < 0) {
      throw new BadRequestException('id must be non-negative');
    }
    return `User id(${id}) is updated`;
  }
}
