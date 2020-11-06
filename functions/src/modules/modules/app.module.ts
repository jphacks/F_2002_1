import { Module } from "@nestjs/common";
import { AppService } from "../../services/app.service";
import { AppController } from "../../controllers/app.controller";

import { ApiModule } from './firebase.module';
@Module({
  providers: [AppService],
  controllers: [AppController],
  imports: [ApiModule]
})
export class AppModule {}
