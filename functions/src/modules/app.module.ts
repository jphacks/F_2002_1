import { Module } from "@nestjs/common";
import { AppService } from "../services/app.service";
import { AppController } from "../controllers/app.controller";

import { FirebaseModule } from './firebase.module';

@Module({
  providers: [AppService],
  controllers: [AppController],
  imports: [FirebaseModule]
})
export class AppModule {}
