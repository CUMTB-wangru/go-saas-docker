import axios from "axios"; // 引用axios
let baseURL = import.meta.env.VITE_API_LOCAL_URL

const instance = axios.create({
    baseURL: String(baseURL),
    timeout: 60000,
});

//get请求
export function get(url: string, params = {}) {
    return new Promise((resolve, reject) => {
        instance.get(url, { params: params, }).then((response) => {
            resolve(response);
        }).catch((err) => {
            reject(err);
        });
    });
}

//post请求
export function post(url: string, data = {}) {
    return new Promise((resolve, reject) => {
        instance.post(url, data).then((response) => {
            resolve(response.data);
        }).catch((err) => {
            reject(err.response.data);
        });
    });
}
