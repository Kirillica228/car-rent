// services/auth-service.ts
import { api } from '@/api/base';

interface User {
  user_id : number;
  role : string;
}

export class AuthService {
  async login(email: string, password: string) {
    // сервер возвращает user, токен хранится в HttpOnly cookie
    const res = await api.post('/auth/login', { email, password });
    return res.data; // { user: {...} }
  }

  async register(email: string, password: string) {
    const res = await api.post('/auth/register', { email, password });
    return res.data; // { user: {...} }
  }

  async logout() {
    await api.post('/auth/logout'); // сервер удаляет cookie
  }

  async me() {
    // проверка текущего пользователя
    const res = await api.get('/authorize/me');
    return res.data; // { user: {...} } или 401
  }
}

export const authService = new AuthService();
