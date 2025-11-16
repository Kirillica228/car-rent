// api/base.ts
import axios, { AxiosInstance } from 'axios';

export const api: AxiosInstance = axios.create({
  baseURL: "http://localhost:80/api/",
  withCredentials: true, // важно! браузер отправляет HttpOnly cookie автоматически
});

// опционально response interceptor для автоматического refresh
// api.interceptors.response.use(
//   res => res,
//   async err => {
//     if (err.response?.status === 401) {
//       // можно вызвать /auth/refresh на бэке, сервер вернёт новую cookie
//       // и повторить запрос один раз
//       try {
//         await api.post('/auth/refresh');
//         return api.request(err.config);
//       } catch {
//         // если refresh не удался — редирект на логин
//         window.location.href = '/login';
//         return Promise.reject(err);
//       }
//     }
//     return Promise.reject(err);
//   }
// );
api.interceptors.response.use(
  res => res,
  err => {
    return Promise.reject(err);
  }
);