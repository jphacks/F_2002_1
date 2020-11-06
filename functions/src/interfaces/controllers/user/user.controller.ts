import { ApiService } from '@/infrastructure/services/user/user.service';
import { Controller, Get } from '@nestjs/common';
//import {ApiService} from "../../services/api/api.service";

@Controller('api')
export class ApiController {
    constructor(private apiService: ApiService){}

    @Get("/")
    public sayHello(){
        return "Hello From API"
    }
}
