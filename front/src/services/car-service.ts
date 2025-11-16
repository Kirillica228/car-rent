import { api } from "@/api/base";
import ICar from "@/types/ICar";
import axios from "axios";

class CarService {
  private baseURL = `catalog/cars`;
  getAll() {
    return api.get<ICar[]>(`${this.baseURL}/list`);
  }

  getById(id: number) {
    return api.get<ICar>(`${this.baseURL}/${id}`);
  }

  create(formData: FormData) {
    return api.post<ICar>(`${this.baseURL}/create`, formData, { headers: { "Content-Type": "multipart/form-data" } });
  }

  update(id: number, formData: FormData) {
    return api.put<ICar>(`${this.baseURL}/${id}`, formData, { headers: { "Content-Type": "multipart/form-data" } });
  }

  delete(id: number) {
    return api.delete<void>(`${this.baseURL}/${id}`);
  }
}

export const carService = new CarService();