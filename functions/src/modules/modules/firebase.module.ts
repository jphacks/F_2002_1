import { Module } from '@nestjs/common';
import { ChatService } from "@/services/chat/chat.service";
import { ChatController } from "@/controllers/chat/chat.controller";

// Firebase のログイン状況を取得する
@Module({
    providers:[ChatService],
    controllers:[ChatController]
})
export class ApiModule {}
