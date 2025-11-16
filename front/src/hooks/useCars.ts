import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { carService } from '@/services/car-service';
import ICar from "@/types/ICar"


// 🔹 Получить все машины
export const useCars = () => {
  return useQuery({
    queryKey: ['cars'],
    queryFn: () => carService.getAll().then(res => res.data),
  });
};


// 🔹 Получить конкретную машину по ID
export const useCar = (id: number) => {
  return useQuery({
    queryKey: ['car', id],
    queryFn: () => carService.getById(id).then(res => res.data),
    enabled: !!id, // не делает запрос, если id нет
  });
};


// 🔹 Создать новую машину
export const useCreateCar = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (formData: FormData) => carService.create(formData).then(res => res.data),
    onSuccess: () => {
      // после создания — обновляем список
      queryClient.invalidateQueries({ queryKey: ['cars'] });
    },
  });
};


// 🔹 Обновить машину
export const useUpdateCar = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: ({ id, formData }: { id: number; formData: FormData }) =>
      carService.update(id, formData).then(res => res.data),
    onSuccess: (_, { id }) => {
      queryClient.invalidateQueries({ queryKey: ['car', id] });
      queryClient.invalidateQueries({ queryKey: ['cars'] });
    },
  });
};


// 🔹 Удалить машину
export const useDeleteCar = () => {
  const queryClient = useQueryClient();

  return useMutation({
    mutationFn: (id: number) => carService.delete(id),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['cars'] });
    },
  });
};