import axios from "axios";
import { API_KEY } from "@/application/config/constants";
import { normalizeApiError } from "@/infrastructure/http/errors";

export const httpClient = axios.create();

httpClient.interceptors.request.use((config) => {
    config.headers.Authorization = `Bearer ${API_KEY}`;
    return config;
});

httpClient.interceptors.response.use(undefined, (error: unknown) => {
    return Promise.reject(normalizeApiError(error));
});
