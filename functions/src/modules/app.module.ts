import { Module } from "@nestjs/common";
import { AppService } from "../services/app.service";
import { AppController } from "../controllers/app.controller";
@Module({
  providers: [AppService],
  controllers: [AppController]
})
export class AppModule {}
