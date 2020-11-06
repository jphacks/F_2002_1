import { Test, TestingModule } from '@nestjs/testing';
import { AppController } from './app.controller';
import { AppService } from '../../infrastructure/services/app.service';

describe('AppController', () => {
  let appController: AppController

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService],
    }).compile();

    appController = app.get<AppController>(AppController);
  });

  describe('Check the sample code', () => {
    it('should return "Hello World!"', () => {
      expect(appController.hello()).toBe('Hello World!');
    });
  });
});