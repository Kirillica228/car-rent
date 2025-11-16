"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import apiFetch from "../../../../utils/api";

type Brand = {
  ID: number;
  name: string;
  is_visible: boolean;
};

export default function BrandDetailPage() {
  const params = useParams();
  const brandId = params?.id as string;

  const [brand, setBrand] = useState<Brand | null>(null);
  const [form, setForm] = useState<Partial<Brand>>({});
  const [editMode, setEditMode] = useState(false);

  async function fetchBrand() {
    try {
      const data = await apiFetch(`/admin/brands/${brandId}`);
      setBrand(data);
      setForm(data);
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    if (brandId) fetchBrand();
  }, [brandId]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSave = async () => {
    try {
      await apiFetch(`/admin/brands/${brandId}`, {
        method: "PUT",
        body: JSON.stringify({
          name: form.name,
          is_visible: form.is_visible,
        }),
      });
      setEditMode(false);
      fetchBrand();
    } catch (err) {
      console.error(err);
      alert("Ошибка при сохранении изменений");
    }
  };

  if (!brand)
    return (
      <p className="p-6 text-gray-500 text-center">
        ⏳ Загружаем данные бренда...
      </p>
    );

  return (
    <div className="p-6 max-w-xl mx-auto space-y-6">
      {/* Заголовок */}
      <div className="flex justify-between items-center">
        <h1 className="text-2xl font-bold text-gray-800">
          Бренд #{brand.ID}
        </h1>

        {!editMode ? (
          <button
            onClick={() => setEditMode(true)}
            className="px-4 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg"
          >
            ✏️ Редактировать
          </button>
        ) : (
          <button
            onClick={handleSave}
            className="px-4 py-2 bg-green-600 hover:bg-green-700 text-white rounded-lg"
          >
            💾 Сохранить
          </button>
        )}
      </div>

      {/* Форма */}
      <div className="border border-gray-300 rounded-lg p-4 bg-white space-y-4">
        {/* Название */}
        <div>
          <label className="block text-gray-700 font-medium mb-2">
            Название бренда
          </label>
          <input
            name="name" // ✅ исправлено (раньше было Name)
            value={form.name || ""}
            onChange={handleChange}
            disabled={!editMode}
            placeholder="Введите название бренда"
            className="w-full p-2 border border-gray-300 rounded-md focus:ring-2 focus:ring-blue-400 focus:outline-none disabled:bg-gray-100"
          />
        </div>

        {/* Чекбокс */}
        <div className="flex items-center gap-2">
          <input
            id="is_visible"
            name="is_visible" // ✅ исправлено
            type="checkbox"
            checked={form.is_visible ?? false}
            onChange={handleChange}
            disabled={!editMode}
            className="w-4 h-4 cursor-pointer disabled:opacity-60"
          />
          <label htmlFor="is_visible" className="text-gray-700 text-sm">
            Отображать бренд в каталоге
          </label>
        </div>
      </div>
    </div>
  );
}
