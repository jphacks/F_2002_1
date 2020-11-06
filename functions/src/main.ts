import { NestFactory } from '@nestjs/core';
import { AppModule } from './modules/modules/app.module';
import * as helmet from "helmet";

async function bootstrap() {
  const app = await NestFactory.create(AppModule);

  // ここにセキュリティについての設定を追加する
  app.use(helmet());

  // NestJS はこれでCORS の設定を追加することができる
  app.enableCors();
  await app.listen(3000);
}
bootstrap();