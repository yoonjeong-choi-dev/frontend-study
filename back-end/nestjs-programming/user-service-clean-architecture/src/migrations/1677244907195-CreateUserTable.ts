import { MigrationInterface, QueryRunner } from 'typeorm';

export class CreateUserTable1677244907195 implements MigrationInterface {
  name = 'CreateUserTable1677244907195';

  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`CREATE TABLE \`user\`
                             (
                                 \`id\`                varchar(255) NOT NULL,
                                 \`name\`              varchar(30)  NOT NULL,
                                 \`email\`             varchar(60)  NOT NULL,
                                 \`password\`          varchar(30)  NOT NULL,
                                 \`signupVerifyToken\` varchar(60)  NOT NULL,
                                 PRIMARY KEY (\`id\`)
                             ) ENGINE=InnoDB`);
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(`DROP TABLE \`user\``);
  }
}
