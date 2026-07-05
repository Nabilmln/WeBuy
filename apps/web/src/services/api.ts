import axios from "axios";

export const userServiceApi = axios.create({
  baseURL: import.meta.env.VITE_USER_SERVICE_URL || "http://localhost:8081",
});

export const productServiceApi = axios.create({
  baseURL: import.meta.env.VITE_PRODUCT_SERVICE_URL || "http://localhost:8082",
});