import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { brandService } from '@/services/brand-service';
import IBrand from '@/types/IBrand';

// 🔹 Получить все бренды
export const useBrands = () => {
  return useQuery({
    queryKey: ['brands'],
    queryFn: () => brandService.getAll().then(res => res.data),
  });
};

// 🔹 Получить бренд по ID
export const useBrand = (id: number) => {
  return useQuery({
    queryKey: ['brand', id],
    queryFn: () => brandService.getById(id).then(res => res.data),
    enabled: !!id,
  });
};

// 🔹 Создать бренд
export const useCreateBrand = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (brand: Omit<IBrand, 'id'>) =>
      brandService.create(brand).then(res => res.data),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['brands'] });
    },
  });
};

// 🔹 Обновить бренд
export const useUpdateBrand = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, brand }: { id: number; brand: Partial<IBrand> }) =>
      brandService.update(id, brand).then(res => res.data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: ['brands'] });
      queryClient.invalidateQueries({ queryKey: ['brand', id] });
    },
  });
};

// 🔹 Удалить бренд
export const useDeleteBrand = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => brandService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['brands'] });
    },
  });
};
