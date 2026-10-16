import axios from "axios";
import {API_KEY} from "./constants.ts";

axios.interceptors.request.use(
    config => {
        config.headers['Authorization'] = `Bearer ${API_KEY}`;
        return config;
    },
    error => {
        return Promise.reject(error);
    }
)