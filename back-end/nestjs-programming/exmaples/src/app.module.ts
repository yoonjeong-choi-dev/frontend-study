import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { Chapter3MainModule } from './chapter3-controller/chapter3-main.module';
import { Chapter14SchedulingModule } from './chapter14-scheduling/chapter14-scheduling.module';

@Module({
  controllers: [AppController],
  providers: [AppService],
  imports: [Chapter3MainModule, Chapter14SchedulingModule],
})
export class AppModule {}
