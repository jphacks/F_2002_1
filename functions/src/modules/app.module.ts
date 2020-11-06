import { Module } from "@nestjs/common";
import { AppService } from "../services/app.service";
import { AppController } from "../controllers/app.controller";

import { FirebaseModule } from './firebase.module';
import {ConfigModule, } from "@nestjs/config";
@Module({
  providers: [AppService],
  controllers: [AppController],
  imports: [FirebaseModule,ConfigModule.forRoot({
    isGlobal:true,
    envFilePath:".env"
  })]
})
export class AppModule {}
