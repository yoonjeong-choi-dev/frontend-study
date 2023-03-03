import { Controller, Post } from '@nestjs/common';
import { TaskService } from './TaskService';

@Controller('chapter14-scheduling')
export class Chapter14SchedulingController {
  constructor(private taskService: TaskService) {}

  @Post('/start')
  start() {
    this.taskService.startJob();
    return 'start job';
  }

  @Post('/stop')
  stop() {
    this.taskService.stopJon();
    return 'end job';
  }
}
