import { Test, TestingModule } from '@nestjs/testing';
import { AppController } from './app.controller';
import { AppService } from '../services/app.service';

describe('AppController', () => {
  let appController: AppController

  beforeEach(async () => {
    const app: TestingModule = await Test.createTestingModule({
      controllers: [AppController],
      providers: [AppService],
    }).compile();

    appController = app.get<AppController>(AppController);
  });

<<<<<<< HEAD
  describe('Check the sample code', () => {
=======
  describe('root', () => {
>>>>>>> origin
    it('should return "Hello World!"', () => {
      expect(appController.hello()).toBe('Hello World!');
    });
  });
});