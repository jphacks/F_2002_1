import { Module } from "@nestjs/common";
import { AppService } from "../services/app.service";
import { AppController } from "../controllers/app.controller";
<<<<<<< HEAD

import { FirebaseModule } from './firebase.module';
import {ConfigModule, } from "@nestjs/config";
@Module({
  providers: [AppService],
  controllers: [AppController],
  imports: [FirebaseModule,ConfigModule.forRoot({
    isGlobal:true,
    envFilePath:".env"
  })]
=======
@Module({
  providers: [AppService],
  controllers: [AppController]
>>>>>>> origin
})
export class AppModule {}
