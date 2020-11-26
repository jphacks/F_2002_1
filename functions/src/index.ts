import * as admin from "firebase-admin";
import * as functions from "firebase-functions";

admin.initializeApp()

// Send Message to the server
exports.post = functions.database.ref("/device_data/{deviceId}/{phone_udid}").onWrite(async (snap, context) => {
  const deviceId = context.params.deviceId;
  const registedData = snap.after.val();
  const registedHumidityData = registedData.humidity;
  const registedTemperatureData = registedData.temperature;
  const registedPressureData =registedData.pressure;
  const registedSolidMoistureData = registedData.solid_moisture;
  const registedIlluminanceData = registedData.illuminance;
  if(
    registedHumidityData.status!=="ok"||
    registedTemperatureData.status!=="ok" ||
    registedPressureData.status!=="ok" ||
    registedSolidMoistureData.status!=="ok" ||
    registedIlluminanceData.status!=="ok"
  ){
      // TODO: ここに通知機能や外部APIなどの連携を実装する
      functions.logger.info("通知を送ります");
  }
  
  
});

