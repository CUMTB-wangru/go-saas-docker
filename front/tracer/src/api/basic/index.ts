import { post, get } from "/@/api/http"
export const getImage = (data: {}) => post("/api", data);
export const register = (data: {}) => post("/api/user/register", data)
export const sendSms = (data: {}) => post("/api/user/sms", data)
export const login = (data: {}) => post("/api/user/login", data)
export const loginSms = (data: {}) => post("/api/user/login_sms", data)
