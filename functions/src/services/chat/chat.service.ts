import { Injectable } from '@nestjs/common';
import * as firebase from "firebase-admin";

@Injectable()
export class ChatService {
    public async getChatMessage(roomId:string){
        try {
            const chatRoom= firebase.database().ref("chat_room").child(roomId)
            let getData: any = await chatRoom.once("value")
            if (getData) {
                const rootList = getData.val();
                const  list = [];
                // データオブジェクトを配列に変更する
                if(rootList !== null) {
                    Object.keys(rootList).forEach((val) => {
                        list.push(rootList[val]);
                    })
                }
                getData = list;
            }
            return getData;
        } catch(e){
            console.error(e);
            throw e;
        }
    }
    public async sendChatMessage(roomId:string,userId:string,message:string){
        try {
            const chatRoom = firebase.database().ref("chat_room").child(roomId);
            const chatKey = await chatRoom.push().key
            const sendMessage = await chatRoom.child(chatKey).set({
                id:chatKey,
                speaker:userId,
                createdAt:Date.now(),
                text:message
            })
            return sendMessage;
        } catch(e){
            console.error(e);
            throw e;
        }
    }
}
