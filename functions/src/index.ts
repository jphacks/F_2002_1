import { NestFactory } from "@nestjs/core";
import { ExpressAdapter } from "@nestjs/platform-express";
import { AppModule } from "./modules/app.module";
import * as express from "express";
import * as functions from "firebase-functions";
import * as helmet from "helmet";

import { ConfigService } from '@nestjs/config';
import * as firebase from 'firebase-admin'
import { ServiceAccount } from "firebase-admin";


const server = express();

export const createNestServer = async (expressInstance: any) => {
  const app = await NestFactory.create(
    AppModule,
    new ExpressAdapter(expressInstance)
  );
  const configService: ConfigService = app.get(ConfigService);
  const adminConfig: ServiceAccount = {
    "projectId"  : configService.get<string>('FIREBASE_PROJECT_ID'),
    "privateKey" : configService.get<string>('FIREBASE_PRIVATE_KEY').replace(/\\n/g, '\n'),
    "clientEmail": configService.get<string>('FIREBASE_CLIENT_EMAIL'),
  };
  // ここにセキュリティについての設定を追加する
  app.use(helmet());
  app.enableCors();

  firebase.initializeApp({
    credential: firebase.credential.cert(adminConfig),
    databaseURL: "https://speak-vegetable.firebaseio.com"
  })
  return app.init();
};

createNestServer(server)
  .then(v => console.log("Nest Ready"))
  .catch(err => console.error("Nest broken", err));

export const api = functions.https.onRequest(server);
