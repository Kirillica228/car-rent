"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import apiFetch from "../../../../utils/api";

export default function AddBrandPage() {
  const router = useRouter();
  const [name, setName] = useState("");
  const [isVisible, setIsVisible] = useState(true); // 👈 чекбокс состояния

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      await apiFetch("/admin/brands/create", {
        method: "POST",
        body: JSON.stringify({
          name,
          is_visible: isVisible, // 👈 передаём флаг
        }),
      });
      router.push("/admin/brand");
    } catch (err) {
      console.error(err);
      alert("Ошибка при добавлении бренда");
    }
  };

  return (
    <div className="w-full max-w-md bg-white border border-gray-300 rounded-lg p-6 shadow-sm">
  <h1 className="text-2xl font-bold text-gray-800 text-center mb-6">
    🏷️ Добавить бренд
  </h1>

  <form onSubmit={handleSubmit} className="flex flex-col gap-5">
    {/* Название */}
    <div>
      <label className="block text-gray-700 font-medium mb-2">
       name
      </label>
      <input
        type="text"
        placeholder="Введите название бренда"
        value={name}
        onChange={(e) => setName(e.target.value)}
        className="w-full px-3 py-2 border border-gray-300 rounded-md text-gray-900 placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:border-transparent disabled:bg-gray-100"
        required
      />
    </div>

    {/* Чекбокс видимости */}
    <div className="flex items-center gap-2">
      <input
        id="is_visible"
        type="checkbox"
        checked={isVisible}
        onChange={(e) => setIsVisible(e.target.checked)}
        className="w-4 h-4 accent-blue-600 cursor-pointer"
      />
      <label htmlFor="is_visible" className="text-gray-700 text-sm">
        is_visible
      </label>
    </div>

    {/* Кнопки */}
    <div className="flex justify-between gap-3">
    
      <button
        type="submit"
        className="flex-1 bg-green-600 hover:bg-green-700 text-white font-medium px-4 py-2 rounded-md transition"
      >
        💾 Сохранить
      </button>
    </div>
  </form>
</div>

  );
}
