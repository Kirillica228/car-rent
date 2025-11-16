"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import ImageUploader from "@/components/imageUploader";
import { useBrands } from "@/hooks/useBrands";
import { useTypes } from "@/hooks/useTypes";
import { useCreateCar } from "@/hooks/useCars";


type CarDetail = {
  brand_id: number;
  type_id: number;
  model: string;
  year: string;
  license_plate: string;
  status: string;
  price: string;
  is_visible: boolean;
  color: string;
  description: string;
  transmission: string;
};

export default function AddCarPage() {
  const router = useRouter();
  const { data: brands = [], isLoading: brandsLoading } = useBrands();
  const { data: types = [], isLoading: typesLoading } = useTypes();
  const createCar = useCreateCar();

  const [form, setForm] = useState<CarDetail>({
    brand_id: 0,
    type_id: 0,
    model: "",
    year: "",
    license_plate: "",
    status: "Готова к аренде",
    price: "",
    is_visible: true,
    color: "",
    description: "",
    transmission: "АКПП",
  });

  const [files, setFiles] = useState<File[]>([]);

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
  ) => {
    const { name, value, type, checked } = e.target as HTMLInputElement;

    setForm((prev) => ({
      ...prev,
      [name]:
        type === "checkbox"
          ? checked
          : name === "brand_id" || name === "type_id"
          ? Number(value)
          : value,
    }));
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    // Формируем FormData для отправки файлов
    const formData = new FormData();
    Object.entries(form).forEach(([key, value]) => {
      formData.append(key, value.toString());
    });

    files.forEach((file) => formData.append("photos", file));

    createCar.mutate(formData, {
      onSuccess: () => {
        router.push("/admin/car");
      },
      onError: (err) => {
        console.error(err);
        alert("Ошибка при добавлении машины");
      },
    });
  };

  if (brandsLoading || typesLoading) {
    return <div className="text-center py-10 text-gray-500">Загрузка списков...</div>;
  }

  return (
    <div className="w-full max-w-2xl bg-white border border-gray-300 rounded-lg p-8 mx-auto">
      <h1 className="text-3xl font-bold text-center text-gray-800 mb-6">
        Добавить машину
      </h1>

      <form onSubmit={handleSubmit} className="flex flex-col gap-5">
        {/* Бренд */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Бренд</label>
          <select
            name="brand_id"
            value={form.brand_id}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-400"
            required
          >
            <option value={0}>Выберите бренд</option>
            {brands.map((b) => (
              <option key={b.id} value={b.id}>
                {b.name}
              </option>
            ))}
          </select>
        </div>

        {/* Тип */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Тип</label>
          <select
            name="type_id"
            value={form.type_id}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none"
            required
          >
            <option value={0}>Выберите тип</option>
            {types.map((t) => (
              <option key={t.id} value={t.id}>
                {t.name}
              </option>
            ))}
          </select>
        </div>

        {/* Модель */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Модель</label>
          <input
            name="model"
            placeholder="Введите модель"
            value={form.model}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 outline-none"
            required
          />
        </div>

        {/* Год */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Год выпуска</label>
          <input
            name="year"
            type="number"
            placeholder="2025"
            value={form.year}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
            required
          />
        </div>

        {/* Номерной знак */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Номерной знак</label>
          <input
            name="license_plate"
            placeholder="A123BC"
            value={form.license_plate}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
            required
          />
        </div>

        {/* Остальные поля (status, price, color, description, transmission) */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Статус</label>
          <select
            name="status"
            value={form.status}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-400"
            required
          >
            <option value="На диагностике">На диагностике</option>
            <option value="Готова к аренде">Готова к аренде</option>
            <option value="Недоступна">Недоступна</option>
          </select>
        </div>

        <div>
          <label className="block text-gray-700 font-medium mb-1">Цена аренды (руб)</label>
          <input
            name="price"
            type="number"
            value={form.price}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
            required
          />
        </div>

        <div className="flex items-center gap-2">
          <input
            name="is_visible"
            type="checkbox"
            checked={form.is_visible}
            onChange={handleChange}
            className="w-4 h-4 accent-blue-600 cursor-pointer"
          />
          <label className="text-gray-700 text-sm">Видимая</label>
        </div>

        <div>
          <label className="block text-gray-700 font-medium mb-1">Цвет</label>
          <input
            name="color"
            value={form.color}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
        </div>

        <div>
          <label className="block text-gray-700 font-medium mb-1">Описание</label>
          <textarea
            name="description"
            value={form.description}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
        </div>

        <div>
          <label className="block text-gray-700 font-medium mb-1">Коробка передач</label>
          <select
            name="transmission"
            value={form.transmission}
            onChange={handleChange}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 focus:outline-none focus:ring-2 focus:ring-blue-400"
          >
            <option value="АКПП">АКПП</option>
            <option value="МКПП">МКПП</option>
          </select>
        </div>

        {/* Фотографии */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Фотографии</label>
          <ImageUploader initialImages={[]} onChange={setFiles} />
        </div>

        <button
          type="submit"
          className="flex-1 bg-green-600 hover:bg-green-700 text-white font-medium px-4 py-2 rounded-md transition"
        >
          💾 Сохранить
        </button>
      </form>
    </div>
  );
}
