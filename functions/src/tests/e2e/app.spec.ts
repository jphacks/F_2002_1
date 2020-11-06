import { Test, TestingModule } from '@nestjs/testing';
import { INestApplication } from '@nestjs/common';
import {AppModule} from "../../modules/app.module";
import {AppController} from "../../controllers/app.controller";
import {AppService} from "../../services/app.service";
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
    afterEach(async () => {
        await app.close();
    });
})