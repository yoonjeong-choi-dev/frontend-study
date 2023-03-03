import { Module } from '@nestjs/common';
import { Chapter14SchedulingController } from './chapter14-scheduling.controller';
import { ScheduleModule } from '@nestjs/schedule';
import { TaskService } from './TaskService';

@Module({
  imports: [ScheduleModule.forRoot()],
  controllers: [Chapter14SchedulingController],
  providers: [TaskService],
})
export class Chapter14SchedulingModule {}
