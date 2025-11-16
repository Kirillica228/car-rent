import IBrand from "@/types/IBrand";
import { api } from '@/api/base';

class BrandService {
  private baseURL = `catalog/types`;

  getAll() {
    return api.get<IBrand[]>(`${this.baseURL}/list`);
  }

  getById(id: number) {
    return api.get<IBrand>(`${this.baseURL}/${id}`);
  }

  create(car: Omit<IBrand, 'id'>) {
    return api.post<IBrand>(`${this.baseURL}/create`, car);
  }

  update(id: number, brand: Partial<IBrand>) {
    return api.put<IBrand>(`${this.baseURL}/${id}`, brand);
  }

  delete(id: number) {
    return api.delete<void>(`${this.baseURL}/${id}`);
  }
}

export const brandService = new BrandService();