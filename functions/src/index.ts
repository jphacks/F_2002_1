import { NestFactory } from "@nestjs/core";
import { ExpressAdapter } from "@nestjs/platform-express";
import { AppModule } from "./interfaces/gateways/modules/app.module";
import * as express from "express";
import * as functions from "firebase-functions";
import * as helmet from "helmet";

const server = express();

export const createNestServer = async (expressInstance: any) => {
  const app = await NestFactory.create(
    AppModule,
    new ExpressAdapter(expressInstance)
  );

  // ここにセキュリティについての設定を追加する
  app.use(helmet());
  app.enableCors();

  console.log("the server is starting @ firebase");
  return app.init();
};

createNestServer(server)
  .then(v => console.log("Nest Ready"))
  .catch(err => console.error("Nest broken", err));

export const api = functions.https.onRequest(server);
