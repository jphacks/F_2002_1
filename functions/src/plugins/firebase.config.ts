import * as firebase from 'firebase-admin'

/*
const firebasePlugin:firebase.ServiceAccount = {
  projectId: process.env.project_id,
  privateKey: process.env.private_key,
  clientEmail: process.env.client_email,
}
*/

firebase.initializeApp({
  credential: firebase.credential.applicationDefault(),
  databaseURL: "https://speak-vegetable.firebaseio.com"
})



export const database = firebase.database();
export const firestore = firebase.firestore();
export const storage = firebase.storage();
