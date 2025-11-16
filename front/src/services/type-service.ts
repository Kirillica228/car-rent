import { api } from "@/api/base";
import IType from "@/types/IType";
import axios from "axios";

class TypeService {
  private baseURL = `catalog/types`;

  getAll() {
    return api.get<IType[]>(`${this.baseURL}/list`);
  }

  getById(id: number) {
    return api.get<IType>(`${this.baseURL}/${id}`);
  }

  create(type: Omit<IType, 'id'>) {
    return api.post<IType>(`${this.baseURL}/create`, type);
  }

  update(id: number, brand: Partial<IType>) {
    return api.put<IType>(`${this.baseURL}/${id}`, brand);
  }

  delete(id: number) {
    return api.delete<void>(`${this.baseURL}/${id}`);
  }
}

export const typeService = new TypeService();