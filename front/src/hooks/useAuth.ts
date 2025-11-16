import { useMutation, useQuery } from '@tanstack/react-query';
import { authService } from '@/services/auth-service';

// Типы
interface LoginVariables {
  email: string;
  password: string;
}

interface Role {
  id: number;
  name: string;
}

interface User {
  user_id: string;
  role: string; // теперь роль есть
}

interface AuthResponse {
  user: User;
}

// 🔹 Логин
export const useLogin = () =>
  useMutation<AuthResponse, Error, LoginVariables>({
    mutationFn: ({ email, password }) => authService.login(email, password),
  });

// 🔹 Регистрация
export const useRegister = () =>
  useMutation<AuthResponse, Error, LoginVariables>({
    mutationFn: ({ email, password }) => authService.register(email, password),
  });

// 🔹 Текущий пользователь
export const useMe = () =>
  useQuery<User>({
    queryKey: ['me'],
    queryFn: () => authService.me(),
    retry: (failureCount, error: any) => {
      if (error.response?.status === 401) return false; // не повторять при 401
      return failureCount < 2; // максимум 2 повтора для остальных ошибок
    },
  });

// 🔹 Логаут
export const useLogout = () =>
  useMutation({
    mutationFn: () => authService.logout(),
  });