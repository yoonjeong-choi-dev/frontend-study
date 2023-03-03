import { Injectable, Logger } from '@nestjs/common';
import { Cron, SchedulerRegistry } from '@nestjs/schedule';
import { CronJob } from 'cron';

const dynamicTaskName = 'cronSample';

@Injectable()
export class TaskService {
  private readonly logger = new Logger(TaskService.name);

  constructor(private readonly schedulerRegistry: SchedulerRegistry) {
    this.addCronJob();
  }

  private addCronJob() {
    const job = new CronJob('* * * * * *', () => {
      this.logger.warn(`Run! ${dynamicTaskName}`);
    });

    this.schedulerRegistry.addCronJob(dynamicTaskName, job);
    this.logger.warn(`job {dynamicTaskName} added!`);
  }

  @Cron('0 * * * * *', { name: 'cronTask' })
  handleCron() {
    this.logger.log('Task Called for every minute!');
  }

  startJob() {
    const job = this.schedulerRegistry.getCronJob(dynamicTaskName);
    job.start();
    this.logger.log(`Job Started: ${job.lastDate()}`);
  }

  stopJon() {
    const job = this.schedulerRegistry.getCronJob(dynamicTaskName);
    job.stop();
    this.logger.log(`Job Stopped: ${job.lastDate()}`);
  }
}
