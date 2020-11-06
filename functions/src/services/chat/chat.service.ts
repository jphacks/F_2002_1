import { Injectable } from '@nestjs/common';
import * as firebase from "firebase-admin";

@Injectable()
export class ChatService {
    public async getChatMessage(roomId:string){
        
        try {
            const chatRoom= firebase.database().ref("chat_room").child(roomId)
            const getData = await chatRoom.on("value",snapshot=>snapshot.val());
            return getData;
        } catch(e){
            console.error(e);
            throw e;
        }
    }
    public async sendChatMessage(roomId:string,userId:string,message:string){
        try {
            const chatRoom = firebase.database().ref("chat_room").child("roomId");
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
