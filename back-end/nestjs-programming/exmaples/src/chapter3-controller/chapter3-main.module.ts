import { Module } from '@nestjs/common';
import { ApiController } from './api/api.controller';
import { Chapter3MainController } from './chapter3-main.controller';

@Module({
  controllers: [ApiController, Chapter3MainController],
})
export class Chapter3MainModule {}
