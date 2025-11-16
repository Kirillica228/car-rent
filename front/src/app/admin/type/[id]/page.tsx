"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import apiFetch from "../../../../utils/api";

type CarType = {
  ID: number;
  name: string;
  is_visible: boolean;
};

export default function TypeDetailPage() {
  const params = useParams();
  const typeId = params?.id as string;

  const [type, setType] = useState<CarType | null>(null);
  const [form, setForm] = useState<Partial<CarType>>({});
  const [editMode, setEditMode] = useState(false);

  async function fetchType() {
    try {
      const data = await apiFetch(`/admin/types/${typeId}`);
      setType(data);
      setForm(data);
    } catch (err) {
      console.error(err);
    }
  }

  useEffect(() => {
    if (typeId) fetchType();
  }, [typeId]);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value, type, checked } = e.target;
    setForm((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleSave = async () => {
    try {
      await apiFetch(`/admin/types/${typeId}`, {
        method: "PUT",
        body: JSON.stringify({
          name: form.name,
          is_visible: form.is_visible,
        }),
      });
      setEditMode(false);
      fetchType();
    } catch (err) {
      console.error(err);
      alert("Ошибка при сохранении изменений");
    }
  };

  if (!type)
    return (
      <p className="p-6 text-gray-400 text-center">
        ⏳ Загружаем данные типа...
      </p>
    );

  return (
    <div className="p-8 max-w-2xl mx-auto space-y-8">
      {/* Заголовок */}
      <div className="flex justify-between items-center">
        <h1 className="text-3xl font-extrabold bg-gradient-to-r from-blue-500 to-purple-600 text-transparent bg-clip-text">
          📦 Тип #{type.ID}
        </h1>
        {!editMode ? (
          <button
            onClick={() => setEditMode(true)}
            className="px-5 py-2 rounded-xl bg-indigo-600 hover:bg-indigo-700 text-white shadow"
          >
            ✏️ Редактировать
          </button>
        ) : (
          <button
            onClick={handleSave}
            className="px-5 py-2 rounded-xl bg-green-600 hover:bg-green-700 text-white shadow"
          >
            💾 Сохранить
          </button>
        )}
      </div>

      {/* Форма */}
      <div className="bg-white/10 backdrop-blur-lg border border-gray-700 rounded-2xl p-6 shadow-xl space-y-6">
        {/* Название */}
        <div>
          <label className="block text-gray-300 font-medium mb-2">
            Название типа
          </label>
          <input
            name="name" // <-- исправлено
            value={form.name || ""}
            onChange={handleChange}
            disabled={!editMode}
            placeholder="Название типа"
            className="w-full p-3 rounded-xl border border-gray-300 bg-white/5 text-black placeholder-gray-400 focus:ring-2 outline-none"
          />
        </div>

        {/* Чекбокс is_visible */}
        <div className="flex items-center gap-3">
          <input
            id="is_visible"
            name="is_visible" // <-- исправлено
            type="checkbox"
            checked={form.is_visible ?? false}
            onChange={handleChange}
            disabled={!editMode}
            className="w-5 h-5 text-blue-600 bg-gray-800 border-gray-600 rounded focus:ring-2 focus:ring-blue-500 cursor-pointer disabled:opacity-60"
          />
          <label htmlFor="is_visible" className="text-gray-300 text-sm">
            Отображать тип в каталоге
          </label>
        </div>
      </div>
    </div>
  );
}
