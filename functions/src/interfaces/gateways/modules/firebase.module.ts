import { Module } from '@nestjs/common';
import { ApiService } from "@/infrastructure/services/user/user.service";
import { ApiController } from "@/interfaces/controllers/user/user.controller";

// Firebase のログイン状況を取得する
@Module({
    providers:[ApiService],
    controllers:[ApiController]
})
export class ApiModule {}
