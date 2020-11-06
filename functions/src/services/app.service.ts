import { Injectable } from "@nestjs/common";

<<<<<<< HEAD
import {IChat} from "../entity/chat";


=======
>>>>>>> origin
@Injectable()
export class AppService {
  public hello(): string {
    return "Hello World!";
  }
<<<<<<< HEAD

  /**
   * テキストメッセージをサーバーに送信する
   * @param roomId :string, 送信先のチャットIDを送信する
   * @param message:string
  */

  
  public async responseText(roomId:string, inputMessage:IChat){
    let responseMessage:string

    // MEMO: この処理をCSV ファイルで検索できるようにする(index を利用して検索する)
    switch(inputMessage.message){
      case "水やり":
        return "";
      default:
        responseMessage = "すみません。このメッセージには対応していません。"
    }
    try {
      console.log("action")
      // ここでチャットを流す
      // database.ref("chat_room").child(roomId).push(); // ここで Firestore へデータを送信する
    } catch(e){
      console.error(e);
      throw new Error(e);
    }

    
    return;
  }
  


=======
>>>>>>> origin
}
