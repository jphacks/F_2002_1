import { Injectable } from "@nestjs/common";
import {database} from "../plugins/firebase.config";


type Chat = {
  id?:string,
  from: string,
  message: string,
  createdAt:string,
}

@Injectable()
export class AppService {
  public hello(): string {
    return "Hello World!";
  }

  /**
   * テキストメッセージをサーバーに送信する
   * @param roomId :string, 送信先のチャットIDを送信する
   * @param message:string, 
  */
  public sendText(roomId:string, chat:Chat){

    // database.ref(); // ここで Firestore へデータを送信する
    return;
  }
  
  public 


}
