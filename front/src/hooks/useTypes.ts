import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { typeService } from '@/services/type-service';
import IType from '@/types/IType';

// 🔹 Получить все типы
export const useTypes = () => {
  return useQuery({
    queryKey: ['types'],
    queryFn: () => typeService.getAll().then(res => res.data),
  });
};

// 🔹 Получить тип по ID
export const useType = (id: number) => {
  return useQuery({
    queryKey: ['type', id],
    queryFn: () => typeService.getById(id).then(res => res.data),
    enabled: !!id,
  });
};

// 🔹 Создать тип
export const useCreateType = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (type: Omit<IType, 'id'>) =>
      typeService.create(type).then(res => res.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['types'] });
    },
  });
};

// 🔹 Обновить тип
export const useUpdateType = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, type }: { id: number; type: Partial<IType> }) =>
      typeService.update(id, type).then(res => res.data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: ['types'] });
      queryClient.invalidateQueries({ queryKey: ['type', id] });
    },
  });
};

// 🔹 Удалить тип
export const useDeleteType = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => typeService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['types'] });
    },
  });
};
