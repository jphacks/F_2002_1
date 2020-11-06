import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import {AppModule} from "../../interfaces/gateways/modules/app.module";
import {AppController} from "../../interfaces/controllers/app.controller";
import {AppService} from "../../infrastructure/services/app.service";
import * as request from 'supertest';
// import { ConfigModule } from '@nestjs/config';

describe("E2E テストについて",()=>{
    
    let app: INestApplication;
    beforeEach(async () => {
        const moduleFixture: TestingModule = await Test.createTestingModule({
          imports: [
              // テスト用の dotenv ファイルである .env.test をコンフィグとして読み込む
            AppModule,
          ],
          controllers: [AppController],
          providers: [AppService],
        }).compile();
    
        app = moduleFixture.createNestApplication();
        await app.init();
    });
    describe("URL：/", ()=>{
        it("正常：200",async ()=>{
            
            const getHelloWorld  = await request(app.getHttpServer()).get("/");
            expect(getHelloWorld.status).toEqual(200);
            expect(getHelloWorld.text).toBe("Hello World!")
        })
    });
    describe("URL: /:roomId",()=>{
        it("正常: 200",async()=>{
            
        })
    })
    afterEach(async () => {
        await app.close();
    });
})