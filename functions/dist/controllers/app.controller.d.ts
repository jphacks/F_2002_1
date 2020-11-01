import { AppService } from "@/services/app.service";
export declare class AppController {
    private appsService;
    constructor(appsService: AppService);
    hello(): string;
}
