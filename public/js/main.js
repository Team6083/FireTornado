// Initialize Firebase
var config = {
    apiKey: "AIzaSyBdsKcG7NeEsqwKDZNX3NKVQPer9euzGx0",
    authDomain: "overproperty-frc6083.firebaseapp.com",
    databaseURL: "https://overproperty-frc6083.firebaseio.com",
    projectId: "overproperty-frc6083",
    storageBucket: "overproperty-frc6083.appspot.com",
    messagingSenderId: "115432376160"
};
firebase.initializeApp(config);

// Initialize Cloud Firestore through Firebase
var db = firebase.firestore();

// Disable deprecated features
db.settings({
    timestampsInSnapshots: true
});

