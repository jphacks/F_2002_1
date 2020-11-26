import bluepy
import requests
import os,sys
import firebase_admin
from firebase_admin import db,credentials
import datetime
import logging
from os.path import join,dirname
from dotenv import load_dotenv

## 環境設定変数についての読み込み

load_dotenv(verbose=True)
dotenv_path = join(dirname(__file__),".env-sample")
load_dotenv(dotenv_path)

## Firebase のシークレットキー情報を読み込み(この部分の変数は秘匿してます。)
firebaseKey = join(dirname(__file__),"accessKey.json")
FirebasePlugin =credentials.Certificate(firebaseKey)

# TODO: このUDID は各デバイスごとに設定すること
SERVICE_UUID = os.environ.get("SERVICE_UUID")
DEVICE_ADDRESS = os.environ.get("DEVICE_ADDRESS")
PHONE_UDID = os.environ.get("PHONE_UDID")


pressure=""
temperature=""
illuminance=""
solidMoisture=""

firebase_admin.initialize_app(FirebasePlugin,{
  "databaseURL": "https://speak-vegetable.firebaseio.com"
})
device_db = db.reference("/device_data")

# BLE の通信方法: Notify 
# データ受信方法:  devKey:value
# devKey: pressure, temp, illuminance, solid , humid
# value に値を送信する
class MyDelegate(bluepy.btle.DefaultDelegate):
    def __init__(self, params):
        bluepy.btle.DefaultDelegate.__init__(self)

    def handleNotification(self, cHandle, data):
        global pressure, temperature,illuminance, solidMoisture
        data = data.decode(encoding='utf-8').split(":")
        if(data[0] =="pressure"):
          pressure =data[1]
        elif(data[0]=="temp"):
          temperature =data[1]
        elif(data[0] =="illuminance"):
          illuminance =data[1]
        elif(data[0]=="solid"):
          solidMoisture =data[1]

def main():
    try:
      peri = bluepy.btle.Peripheral()
      peri.connect(DEVICE_ADDRESS, bluepy.btle.ADDR_TYPE_PUBLIC)
      peri.withDelegate(MyDelegate(bluepy.btle.DefaultDelegate))
      
      nowDatetime = datetime.datetime.now()
      createdAt = nowDatetime.strftime('%Y/%m/%d-%H:%M:%S')
      while(pressure=="" or temperature=="" or illuminance =="" or solidMoisture==""):
          if peri.waitForNotifications(1.0):
              continue
      
    except Exception as e:
      logging.error(e)
      peri.disconnect()
    
    peri.disconnect()
    print(nowDatetime)
    print("全データを取得")
    print("気圧:",pressure)
    print("気温:",temperature)
    print("照度:",illuminance)
    print("土壌:",solidMoisture)
    device_db.child(SERVICE_UUID).child(PHONE_UDID).set({
      "createdAt":createdAt,
      "cultivationId":"xxxxxxxxxx",
      "humidity":{
        "value":"80",
       "status":"ok"
      },
      "pressure":{
       "value": pressure,
       "status":"ok"
       },
      "illuminance":{
        "value":illuminance,   
       "status":"ng"
      },
      "pressure":{
        "value":pressure,
       "status":"ok"
      },
      "solid_moisture":{
        "value":solidMoisture,  
       "status":"ok"
      },
      "temperature":{
        "value":temperature,
       "status":"ok"
      }
      })
    # TODO:　ラズパイから Firebase に直接データを送信するようにして
    # getData = requests.post(BASE_URL)
    # print(getData)

if __name__ == "__main__":
    main()
