import * as admin from "firebase-admin";
import * as functions from "firebase-functions";


admin.initializeApp()

// Send Message to the server
exports.postChat = functions.database.ref("") .onCreate((snap, context) => {
  const appOptions = JSON.parse(process.env.FIREBASE_CONFIG);
  appOptions.databaseAuthVariableOverride = context.auth;
  const app = admin.initializeApp(appOptions, 'app');
  const uppercase = snap.val();
  
});
