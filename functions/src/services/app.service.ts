import { Injectable } from "@nestjs/common";

@Injectable()
export class AppService {
  public hello(): string {
    return "HEllo World";
  }
}
