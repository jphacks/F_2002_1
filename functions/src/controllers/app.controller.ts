import { Controller, Get } from "@nestjs/common";
import { AppService } from "@/services/app.service";
@Controller()
export class AppController {
  constructor(private appsService: AppService) {}
  @Get()
  hello(): string {
    return this.appsService.hello();
  }
}
