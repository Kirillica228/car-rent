"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import apiFetch from "../../../../utils/api";
import ImageUploader from "@/components/imageUploader";

type Brand = { ID: number; name: string };
type CarType = { ID: number; name: string };
type Image = { ID: number; CarID: number; URL: string };
type CarDetail = {
  ID: number;
  Brand: Brand;
  Type: CarType;
  Model: string;
  Year: number;
  LicensePlate: string;
  Status: string;
  Price: number;
  IsVisible: boolean;
  Color: string;
  Description: string;
  Transmission: string;
  Images: Image[];
};

export default function CarDetailPage() {
  const params = useParams();
  const carId = params?.id as string;

  const [car, setCar] = useState<CarDetail | null>(null);
  const [form, setForm] = useState<Partial<CarDetail>>({});
  const [editMode, setEditMode] = useState(false);
  const [brands, setBrands] = useState<Brand[]>([]);
  const [types, setTypes] = useState<CarType[]>([]);
  const [files, setFiles] = useState<File[]>([]);
  const [removedExisting, setRemovedExisting] = useState<string[]>([]);
  const [loading, setLoading] = useState(false);

  // Загрузка данных
  useEffect(() => {
    async function fetchData() {
      if (!carId) return;
      const [carRes, brandsRes, typesRes] = await Promise.all([
        apiFetch(`/admin/cars/${carId}`),
        apiFetch("/admin/brands/list"),
        apiFetch("/admin/types/list"),
      ]);
      setCar(carRes);
      setForm(carRes);
      setBrands(Array.isArray(brandsRes) ? brandsRes : []);setTypes(Array.isArray(typesRes) ? typesRes : []);
    }
    fetchData();
  }, [carId]);

  // Обработка изменения полей
  const handleChange = (
  e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>
) => {
  const target = e.target;
  const { name, value, type, checked } = target as HTMLInputElement;

  setForm((prev) => ({
    ...prev,
    [name]: type === "checkbox" ? checked : value,
  }));
};

  // Сохранение изменений
  const handleSave = async () => {
    try {
      setLoading(true);

      const formData = new FormData();
      formData.append("brand_id", String(form.Brand?.ID));
      formData.append("type_id", String(form.Type?.ID));
      formData.append("model", form.Model ?? "");
      formData.append("year", String(form.Year ?? ""));
      formData.append("license_plate", form.LicensePlate ?? "");
      formData.append("status", form.Status ?? "");
      formData.append("price", String(form.Price ?? ""));
      formData.append("is_visible", String(form.IsVisible ?? false));
      formData.append("color", form.Color ?? "");
      formData.append("description", form.Description ?? "");
      formData.append("transmission", form.Transmission ?? "АКПП");
      formData.append("status", form.Status ?? "Готова к аренде");

      // Новые фото
      files.forEach((file) => formData.append("photos", file));

      // Удалённые фото
      removedExisting.forEach((url) =>
        formData.append("removed_photos", url)
      );
      await apiFetch(`/admin/cars/${carId}`, {
        method: "PUT",
        body: formData,
      });

      setEditMode(false);
      const updated = await apiFetch(`/admin/cars/${carId}`);
      setCar(updated);
      setForm(updated);
      setFiles([]);
      setRemovedExisting([]);
    } catch (err) {
      console.error(err);
      alert("Ошибка при сохранении изменений");
    } finally {
      setLoading(false);
    }
  };

  if (!car)
    return (
      <p className="p-6 text-gray-400 text-center">⏳ Загружаем данные...</p>
    );

  return (
    <div className="max-w-3xl mx-auto bg-white border border-gray-300 rounded-2xl p-8 mt-8 space-y-6">
      {/* Заголовок */}
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-bold text-gray-800">
          ✏️ Редактирование машины #{car.ID}
        </h1>
        {!editMode ? (
          <button
            onClick={() => setEditMode(true)}
            className="bg-blue-600 hover:bg-blue-700 text-white font-medium px-5 py-2 rounded-lg transition"
          >
            Редактировать
          </button>
        ) : (
          <button
            onClick={handleSave}
            disabled={loading}
            className="bg-green-600 hover:bg-green-700 text-white font-medium px-5 py-2 rounded-lg transition disabled:opacity-60"
          >
            {loading ? "⏳ Сохраняем..." : "💾 Сохранить"}
          </button>
        )}
      </div>

      {/* Основная форма */}
      <div className="flex flex-col gap-4">
        {/* Бренд */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Бренд</label>
          <select
            name="Brand.ID"
            value={form.Brand?.ID ?? car.Brand.ID}
            onChange={(e) =>
              setForm((prev) => ({
                ...prev,
                Brand: { ID: Number(e.target.value), name: "" },
              }))
            }
            disabled={!editMode}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
          >
            <option value="">Выберите бренд</option>
            {brands.map((b) => (
              <option key={b.ID} value={b.ID}>
                {b.name}
              </option>
            ))}
          </select>
        </div>

        {/* Тип */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Тип</label>
          <select
            name="Type.ID"
            value={form.Type?.ID ?? car.Type.ID}
            onChange={(e) =>
              setForm((prev) => ({
                ...prev,
                Type: { ID: Number(e.target.value), name: "" },
              }))
            }
            disabled={!editMode}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
          >
            <option value="">Выберите тип</option>
            {types.map((t) => (
              <option key={t.ID} value={t.ID}>
                {t.name}
              </option>
            ))}
          </select>
        </div>
        {/* Статус */}
        <div>
          <label className="block text-gray-700 font-medium mb-1">Статус</label>
          <select
            name="Status"
            value={form.Status ?? car.Status ?? "Готова к аренде"}
            onChange={handleChange}
            disabled={!editMode}
            className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
          >
            <option value="На диагностике">На диагностике</option>
            <option value="Готова к аренде">Готова к аренде</option>
            <option value="Недоступна">Недоступна</option>
          </select>
        </div>

        {/* Остальные поля */}
        {[
          { label: "Модель", name: "Model", type: "text", placeholder: "Введите модель" },
          { label: "Год выпуска", name: "Year", type: "number", placeholder: "2025" },
          { label: "Номерной знак", name: "LicensePlate", type: "text", placeholder: "A123BC" },
          { label: "Цена аренды ($)", name: "Price", type: "number", placeholder: "100" },
        ].map(({ label, name, type, placeholder }) => (
          <div key={name}>
            <label className="block text-gray-700 font-medium mb-1">{label}</label>
            <input
              name={name}
              type={type}
              value={(form as any)[name] ?? (car as any)[name] ?? ""}
              onChange={handleChange}
              disabled={!editMode}
              placeholder={placeholder}
              className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
            />
          </div>
        ))}

        {/* Видимость */}
        <div className="flex items-center gap-2">
          <input
            name="IsVisible"
            type="checkbox"
            checked={form.IsVisible ?? car.IsVisible ?? false}
            onChange={(e) =>
              setForm((prev) => ({ ...prev, IsVisible: e.target.checked }))
            }
            disabled={!editMode}
            className="w-4 h-4 accent-blue-600 cursor-pointer disabled:opacity-60"
          />
          <label className="text-gray-700 text-sm">Видимая</label>
        </div>
      </div>
      {/* Цвет */}
      <div>
        <label className="block text-gray-700 font-medium mb-1">Цвет</label>
        <input
          name="Color"
          value={form.Color ?? car.Color ?? ""}
          onChange={handleChange}
          disabled={!editMode}
          placeholder="Введите цвет"
          className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
        />
      </div>

      {/* Описание */}
      <div>
        <label className="block text-gray-700 font-medium mb-1">Описание</label>
        <textarea
          name="Description"
          value={form.Description ?? car.Description ?? ""}
          onChange={handleChange}
          disabled={!editMode}
          placeholder="Введите описание"
          className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
        />
      </div>

      {/* Коробка передач */}
      <div>
        <label className="block text-gray-700 font-medium mb-1">Коробка передач</label>
        <select
          name="Transmission"
          value={form.Transmission ?? car.Transmission ?? "АКПП"}
          onChange={handleChange}
          disabled={!editMode}
          className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 outline-none disabled:opacity-60"
        >
          <option value="АКПП">АКПП</option>
          <option value="МКПП">МКПП</option>
        </select>
      </div>
      {/* Фото */}
      <div>
        <h2 className="text-lg font-semibold text-gray-800 mb-2">Фотографии</h2>
        <ImageUploader
          initialImages={car.Images.map((img) => img.URL)}
          onChange={(newFiles, removed) => {
            setFiles(newFiles);
            setRemovedExisting(removed);
          }}
        />
      </div>
    </div>
  );
}
