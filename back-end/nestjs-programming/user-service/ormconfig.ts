import { DataSource } from 'typeorm';

const AppDataSource = new DataSource({
  type: 'mysql',
  host: 'localhost',
  port: 3306,
  username: 'nestJsBackendProgrammingUser',
  password: 'nestJsBackendProgrammingPassword',
  database: 'nestJsBackendProgramming',
  entities: [__dirname + '/**/*.entity{.ts,.js}'],
  synchronize: false,
  migrations: [__dirname + '/**/migrations/*.js'],
  migrationsTableName: 'migrations',
  migrationsRun: false, // cli 명령어로만 수행
});

export default AppDataSource;
