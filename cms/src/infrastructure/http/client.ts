import axios from "axios";
import { API_KEY } from "@/application/config/constants";

export const httpClient = axios.create();

httpClient.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${API_KEY}`;
    return config;
});
