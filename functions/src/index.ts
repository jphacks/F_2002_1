import * as admin from "firebase-admin";
import * as functions from "firebase-functions";

admin.initializeApp()

// Send Message to the server
exports.post = functions.database.ref("/device_data/{deviceId}/{phone_udid}").onWrite(async (snap, context) => {
  const deviceId = context.params.deviceId;
  const phone_udid = context.params.phone_udid;
  const registedData = snap.after.val();
  const cultivationId = registedData.cultivationId
  const getThreshold = await admin.database().ref(`/threshold/${cultivationId}`).once("value")
  const thresholdValues= getThreshold.val();

  // この部分でセンサーの値 の 閾値の判定を行う
  // 値が閾値より大きければ ok
  const registedHumidityData = registedData.humidity;
  registedHumidityData.status =Number(registedHumidityData.value) <Number(thresholdValues.humidity)?"ng":"ok";
  const registedTemperatureData = registedData.temperature;
  registedTemperatureData.status =Number(registedHumidityData.value) < Number(thresholdValues.temperature)?"ng":"ok";
  const registedPressureData =registedData.pressure;
  registedPressureData.status = Number(registedPressureData.value) < Number(thresholdValues.pressure)?"ng":"ok";
  const registedSolidMoistureData = registedData.solid_moisture;

  // 値が閾値より小さければ ok 
  registedSolidMoistureData.status = Number(thresholdValues.solid_moisture) < Number(registedSolidMoistureData.value)  ?"ng":"ok";
  const registedIlluminanceData = registedData.illuminance;
  registedIlluminanceData.status = Number(thresholdValues.illuminance) < Number(registedIlluminanceData.value) ?"ng":"ok";

  const setSensorData = {
    ...registedData,
    humidity: registedHumidityData,
    illuminance:registedIlluminanceData,
    pressure:registedPressureData,
    solid_moisture:registedSolidMoistureData,
    temperature: registedTemperatureData
  }
  await admin.database().ref(`/device_data/${deviceId}/${phone_udid}`).update(setSensorData).then(()=>{
    functions.logger.log("値の変更が完了しました!")
  }).catch(e=>{
    functions.logger.log("値の変更に失敗しました。エラーログを確認してください");
    functions.logger.error(e);
  })
  // データをupdate する必要がある
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

