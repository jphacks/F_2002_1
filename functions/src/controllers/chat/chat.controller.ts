import { Body, Controller, Get,Param,Post} from '@nestjs/common';
import {ChatService} from "../../services/chat/chat.service";

@Controller('chat')
export class ChatController {
    constructor(private chatService:ChatService){}
    @Get(":roomId")
    public getChat(@Param("roomId") roomId:string){
        return this.chatService.getChatMessage(roomId);
    }
    @Post(":roomId/:userId")
    public sendMessage(@Param() params, @Body() body){
        const roomId = params.roomId;
        const userId= params.userId;
        const sendMessage = body.message;
        return this.chatService.sendChatMessage(roomId,userId,sendMessage)
    }

}
