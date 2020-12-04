
// Reading Header Files
#include "Adafruit_MCP9808.h"
#include <Dps310.h>

#include <BLEDevice.h>
#include <BLEUtils.h>
#include <BLEServer.h>
#include <BLE2902.h>
#include <LiquidCrystal_I2C.h>
#include "DHT.h"

#define DHTPIN 14
#define DHTTYPE DHT11
#define SERVICE_UUID        "4fafc201-1fb5-459e-8fcc-c5c9c331914b"
#define CHARACTERISTIC_UUID "beb5483e-36e1-4688-b7f5-ea07361b26a8"

 // Setting sensors
Dps310 Dps310PressureSensor = Dps310();
Adafruit_MCP9808 tempsensor = Adafruit_MCP9808();
DHT dht(DHTPIN, DHTTYPE);

// Setting BLE
BLEServer *pServer;
BLEService* pService;
BLECharacteristic *pCharacteristicPressure;
BLECharacteristic *pCharacteristicTemp;
BLECharacteristic *pCharacteristicIlluminance;
BLECharacteristic *pCharacteristicSolidMoist;
BLECharacteristic *pCharacteristicHumidity;

bool deviceConnected = false;

class MyServerCallbacks: public BLEServerCallbacks {
    void onConnect(BLEServer* pServer) {
      deviceConnected = true;
    };
 
    void onDisconnect(BLEServer* pServer) {
      deviceConnected = false;
    }
};

/**
 * MEMO: 送信されるデータの通信規格について
 *
 * - Humidity: Analog
 * - Illuminance: Analog
 * - Pressure: I2C 
 * - SolidMoisture: Analog
 * - Temperature: I2C
 * 
 *  送信方法: Bluetooth LE で Notify 形式で送信 (今後取得したデータをもとにグラフ化を行い、今後の予測に役立てるため)
*/


void setup() {
  
  Serial.begin(115200);
  Serial.println("ESP32");
  Dps310PressureSensor.begin(Wire);

  int16_t temp_mr = 2;
  int16_t temp_osr = 2;
  int16_t prs_mr = 2;
  int16_t prs_osr = 2;
  int16_t ret = Dps310PressureSensor.startMeasureBothCont(temp_mr, temp_osr, prs_mr, prs_osr);
  
  // Initialize the BLE devices 
  
  BLEDevice::init("Speak Vegetable Device");
  
  pServer = BLEDevice::createServer();
  pServer->setCallbacks(new MyServerCallbacks());
 
  // Create the BLE Service
  pService = pServer->createService(SERVICE_UUID);
 
  // Create a BLE Characteristic
  pCharacteristicPressure= pService->createCharacteristic(
                      CHARACTERISTIC_UUID,
                      BLECharacteristic::PROPERTY_NOTIFY |
                      BLECharacteristic::PROPERTY_INDICATE
                    );
  pCharacteristicTemp= pService->createCharacteristic(
                      CHARACTERISTIC_UUID,
                      BLECharacteristic::PROPERTY_NOTIFY |
                      BLECharacteristic::PROPERTY_INDICATE
                    );
  pCharacteristicIlluminance= pService->createCharacteristic(
                      CHARACTERISTIC_UUID,
                      BLECharacteristic::PROPERTY_NOTIFY |
                      BLECharacteristic::PROPERTY_INDICATE
                    );
  pCharacteristicSolidMoist= pService->createCharacteristic(
                      CHARACTERISTIC_UUID,
                      BLECharacteristic::PROPERTY_NOTIFY |
                      BLECharacteristic::PROPERTY_INDICATE
                    );
  pCharacteristicHumidity =  pService->createCharacteristic(
                      CHARACTERISTIC_UUID,
                      BLECharacteristic::PROPERTY_NOTIFY |
                      BLECharacteristic::PROPERTY_INDICATE
                    );

  BLE2902* ble9202 = new BLE2902();
  ble9202->setNotifications(true);
  pCharacteristicPressure->addDescriptor(ble9202);
  pCharacteristicTemp->addDescriptor(ble9202);
  pCharacteristicIlluminance->addDescriptor(ble9202);
  pCharacteristicSolidMoist->addDescriptor(ble9202);
  pCharacteristicHumidity->addDescriptor(ble9202);
  // Start the service
  pService->start();

  
  BLEAdvertising *pAdvertising = BLEDevice::getAdvertising();
  pAdvertising->addServiceUUID(SERVICE_UUID);
  pAdvertising->setScanResponse(true);
  pAdvertising->setMinPreferred(0x06);  // functions that help with iPhone connections issue
  pAdvertising->setMinPreferred(0x12);
  BLEDevice::startAdvertising();
 
  // Start advertising
  pServer->getAdvertising()->start();
  Serial.println("Waiting a client connection to notify...");
}
void loop() {

  // Illuminance, Soil_Mist Sensor 
  
  uint8_t illuminance,solid_moist,humidity;

  illuminance = analogRead(33);
  solid_moist = analogRead(32);
  humidity = dht.readHumidity();
  
  if (isnan(h) || isnan(t)) {
    Serial.println("Failed to read from DHT sensor!");
  }


  //　Pressure, Temperature Sensor
  uint8_t pressureCount = 20,temperatureCount = 20;

  // need to initialization
  float avrPressure=0.0, avrTemp=0.0; 
  float pressure[pressureCount],temperature[temperatureCount];
  int16_t ret = Dps310PressureSensor.getContResults(temperature, temperatureCount, pressure, pressureCount);

  for (int16_t i = 0; i < temperatureCount; i++){
    avrTemp += temperature[i];
  }
  for (int16_t i = 0; i < pressureCount; i++)
  {
    avrPressure+= pressure[i];
  }              

  // Send Information with BLE 
   avrPressure /= (double)pressureCount;
   avrTemp /= (double)temperatureCount;

   char sendPressureData[100] = {'\0'};
   char sendTempData[100] = {'\0'};
   char sendIlluminanceData[100] = {'\0'};
   char sendSolidMoistData[100] = {'\0'};
   char sendHumidityData[100] = {'\0'};
  if (deviceConnected) {
    sprintf(sendPressureData,"pressure:%f",avrPressure);
    sprintf(sendTempData,"temp:%f",avrTemp);
    sprintf(sendIlluminanceData,"illuminance:%d",illuminance);
    sprintf(sendSolidMoistData,"solid:%d",solid_moist);
    sprintf(sendHumidityData,"humidity:%d",humidity);
    Serial.println(sendPressureData);
    Serial.println(sendTempData);
    Serial.println(sendIlluminanceData);
    Serial.println(sendSolidMoistData);
    Serial.println(sendHumidityData);
    pCharacteristicPressure->setValue(sendPressureData);
    pCharacteristicPressure->notify();
    pCharacteristicTemp->setValue(sendTempData);
    pCharacteristicTemp->notify();
    pCharacteristicIlluminance->setValue(sendIlluminanceData);
    pCharacteristicIlluminance->notify();
    pCharacteristicSolidMoist->setValue(sendSolidMoistData);
    pCharacteristicSolidMoist->notify();
    pCharacteristicHumidity->setValue(sendHumidityData);
    pCharacteristicHumidity->notify();
  }
  
  delay(1000);
  
}
