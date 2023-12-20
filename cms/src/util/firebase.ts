// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
  apiKey: "AIzaSyC59xS5P27qzafvucTLTPkAm6gwrWRfMQ8",
  authDomain: "momentum-ministry.firebaseapp.com",
  projectId: "momentum-ministry",
  storageBucket: "momentum-ministry.appspot.com",
  messagingSenderId: "649626638654",
  appId: "1:649626638654:web:2c65617d5ee8ed16cab46f",
  measurementId: "G-F5RPPG148R"
};

const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);
