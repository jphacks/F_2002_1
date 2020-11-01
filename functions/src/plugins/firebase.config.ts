import * as firebase from 'firebase-admin'

firebase.initializeApp({
  credential: firebase.credential.applicationDefault(),
})

export const database = firebase.database();
export const firestore = firebase.firestore();
export const storage = firebase.storage();
