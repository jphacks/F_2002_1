import { Controller, Get, Param, Post ,Body} from "@nestjs/common";
import { AppService } from "../services/app.service";
import {IChat} from "../interfaces/chat";

@Controller()
export class AppController {
  constructor(private appsService: AppService) {}
  @Get()
  hello(): string {
    return this.appsService.hello();
  }
  // チャット
  @Post("/chat/:roomId")
  responseText(roomId:string,chat:IChat){
    return this.appsService.responseText(roomId,chat)
  }

}
