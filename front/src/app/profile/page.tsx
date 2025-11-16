"use client";

import { useEffect, useState } from "react";
import apiFetch from "../../utils/api";
import Link from "next/link";
import Image from "next/image";

type Rent = {
  ID: number;
  CarID: number;
  Name: string;
  Image: string;
  StartDate: string;
  EndDate: string;
  Status: string;
  TotalPrice: number;
  CreatedAt: string;
};

export default function ProfilePage() {
  const [rents, setRents] = useState<Rent[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    (async () => {
      try {
        const res = await apiFetch("/rents/list", { method: "GET" });
        setRents(Array.isArray(res) ? res : []);
      } catch (err) {
        console.error("Ошибка загрузки аренд:", err);
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 p-8">
      <h1 className="text-4xl font-extrabold bg-gradient-to-r from-blue-500 to-purple-600 text-transparent bg-clip-text mb-8">
        👤 Мой профиль
      </h1>

      <div className="bg-white/10 backdrop-blur-lg border border-gray-700 rounded-2xl shadow-xl p-6">
        <h2 className="text-2xl font-bold text-white mb-6">📋 История аренд</h2>

        {loading ? (
          <p className="text-gray-400">⏳ Загружаем аренды...</p>
        ) : rents.length === 0 ? (
          <p className="text-gray-400">🚫 У вас пока нет аренд</p>
        ) : (
          <div className="space-y-6">
            {rents.map((rent) => {
  const imageUrl =
    rent.Image && rent.Image.length > 0
      ? `http://localhost:8080/api/catalog${rent.Image}`
      : "/placeholder-car.jpg";

  const handleCancel = async () => {
  if (!confirm("Вы уверены, что хотите отменить эту аренду?")) return;

  try {
    const res = await apiFetch(`/rents/cancel/${rent.ID}`, { method: "PUT" });

    // если apiFetch выбрасывает ошибку — перехватим в catch
    if (res?.status && res.status !== 204 && res.status !== 200) {
      switch (res.status) {
        case 403:
          alert("Нельзя отменить аренду менее чем за 1 день до начала.");
          break;
        case 404:
          alert("Аренда не найдена или уже отменена.");
          break;
        case 409:
          alert("Аренда уже отменена ранее.");
          break;
        default:
          alert("Произошла ошибка при отмене аренды.");
      }
      return;
    }

    // ✅ Успешная отмена
    setRents((prev) => prev.filter((r) => r.ID !== rent.ID));
    alert("✅ Аренда успешно отменена!");
  } catch (err: any) {
    console.error("Ошибка при отмене аренды:", err);

    // Проверим, если apiFetch возвращает объект ошибки с кодом
    const status = err?.status;
    switch (status) {
      case 403:
        alert("Нельзя отменить аренду менее чем за 1 день до начала.");
        break;
      case 404:
        alert("Аренда не найдена или уже отменена.");
        break;
      case 409:
        alert("Аренда уже отменена ранее.");
        break;
      default:
        alert(`Ошибка при отмене аренды: ${err?.message || err}`);
    }
  }
};


  return (
    <div
      key={rent.ID}
      className="flex flex-col md:flex-row items-center gap-6 bg-white/5 border border-gray-700 rounded-2xl p-5 hover:shadow-2xl hover:scale-[1.01] transform transition duration-300"
    >
      <div className="flex-1 space-y-2">
        <Link
          href={`/catalog/${rent.CarID}`}
          className="text-xl font-bold text-white hover:text-purple-400 transition"
        >
          {rent.Name}
        </Link>

        <p className="text-gray-400">
          Период:{" "}
          <span className="text-white">
            {new Date(rent.StartDate).toLocaleDateString()} –{" "}
            {new Date(rent.EndDate).toLocaleDateString()}
          </span>
        </p>

        <p className="text-gray-400">
          Статус:{" "}
          <span
            className={(() => {
              switch (rent.Status) {
                case "Новый":
                  return "text-blue-400 font-semibold"; // активная аренда
                case "Завершён":
                  return "text-green-400 font-semibold"; // успешно завершена
                case "Отменён":
                  return "text-red-400 font-semibold"; // отменена пользователем
                default:
                  return "text-gray-400 font-semibold"; // на всякий случай
              }
            })()}
          >
            {rent.Status}
          </span>
        </p>


        <p className="text-gray-400">
          Дата создания: {new Date(rent.CreatedAt).toLocaleDateString("ru-RU")}
        </p>

        <p className="text-green-400 font-semibold text-lg">
          {rent.TotalPrice} р
        </p>

        {/* Кнопка отмены */}
        {rent.Status.toLowerCase() === "новый" && (
          <button
            onClick={handleCancel}
            className="mt-2 px-4 py-2 bg-red-600 hover:bg-red-700 text-white rounded-lg font-medium transition"
          >
            Отменить аренду
          </button>
        )}
      </div>
    </div>
  );
})}
          </div>
        )}
      </div>
    </div>
  );
}
