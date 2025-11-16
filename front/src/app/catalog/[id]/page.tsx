"use client";

import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import Cookies from "js-cookie";
import apiFetch from "../../../utils/api";
import Link from "next/link";

const API_URL = "http://localhost:8080/api/catalog";

type Images = { ID: number; CarID: number; URL: string };
type Brand = { ID: number; name: string };
type Type = { ID: number; name: string };

type Car = {
  ID: number;
  Brand: Brand;
  Type: Type;
  Model: string;
  Year: number;
  Price: number;
  Images: Images[];
  Color: string;
  Description: string;
  Transmission: string;
};

export default function CarPage() {
  const { id } = useParams<{ id: string }>();
  const [car, setCar] = useState<Car | null>(null);
  const [loading, setLoading] = useState(true);
  const token = Cookies.get("token");
  const [currentIndex, setCurrentIndex] = useState(0);

  useEffect(() => {
    (async () => {
      try {
        const res = await apiFetch(`/catalog/cars/${id}`, { method: "GET" });
        setCar(res);
      } catch (err) {
        console.error("Ошибка загрузки машины:", err);
      } finally {
        setLoading(false);
      }
    })();
  }, [id]);

  if (loading)
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-900 text-gray-400">
        Загрузка...
      </div>
    );

  if (!car)
    return (
      <div className="min-h-screen flex items-center justify-center bg-gray-900 text-red-400">
        Машина не найдена
        <button onClick={() => console.log(car)}>fffff</button>
      </div>
    );

  const images =
    car.Images?.length > 0
      ? car.Images
      : [{ ID: 0, CarID: car.ID, URL: "/placeholder-car.jpg" }];

  const prevSlide = () =>
    setCurrentIndex((i) => (i === 0 ? images.length - 1 : i - 1));
  const nextSlide = () =>
    setCurrentIndex((i) => (i === images.length - 1 ? 0 : i + 1));

  return (
    <div className="min-h-screen bg-gray-900 text-white py-16 px-6">
      <div className="max-w-6xl mx-auto bg-gray-800 border border-gray-700 rounded-3xl overflow-hidden shadow-[0_0_40px_rgba(255,255,255,0.05)]">
        {/* Фото блока */}
        <div className="relative w-full h-[500px] overflow-hidden">
          <img
            src={`${API_URL}${images[currentIndex].URL}`}
            alt={`${car.Brand.name} ${car.Model}`}
            className="w-full h-full object-cover object-center transition-transform duration-700 ease-out hover:scale-105"
          />

          {/* Переключатели */}
          {images.length > 1 && (
            <>
              <button
                onClick={prevSlide}
                className="absolute left-6 top-1/2 -translate-y-1/2 bg-black/40 hover:bg-black/70 text-white px-4 py-2 rounded-full transition"
              >
                ‹
              </button>
              <button
                onClick={nextSlide}
                className="absolute right-6 top-1/2 -translate-y-1/2 bg-black/40 hover:bg-black/70 text-white px-4 py-2 rounded-full transition"
              >
                ›
              </button>

              <div className="absolute bottom-6 left-1/2 -translate-x-1/2 flex gap-2">
                {images.map((_, i) => (
                  <div
                    key={i}
                    onClick={() => setCurrentIndex(i)}
                    className={`w-2.5 h-2.5 rounded-full cursor-pointer ${
                      i === currentIndex ? "bg-white" : "bg-gray-500"
                    }`}
                  />
                ))}
              </div>
            </>
          )}
        </div>

       {/* Контент */}
<div className="p-10 md:p-14 space-y-6">
  {/* Основные характеристики */}
  <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
    <div className="space-y-2">
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Бренд:</span> {car.Brand.name}
      </p>
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Модель:</span> {car.Model}
      </p>
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Тип:</span> {car.Type.name}
      </p>
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Год выпуска:</span> {car.Year}
      </p>
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Цвет:</span> {car.Color}
      </p>
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Коробка передач:</span> {car.Transmission}
      </p>
    </div>

    <div className="text-left md:text-right space-y-2">
      <p className="text-gray-400 text-sm">
        <span className="font-semibold">Цена аренды:</span> {car.Price} р / день
      </p>
    </div>
  </div>

  {/* Описание */}
  <div className="bg-gray-800 border border-gray-700 rounded-xl p-4 text-gray-300">
    Описание: 
    <p>{car.Description || "-"}</p>
  </div>

  {/* Кнопка */}
  {token ? (
    <Link
      href={`/rent/${car.ID}`}
      className="px-8 py-3 bg-white text-black rounded-xl font-semibold hover:bg-gray-200 transition"
    >
      Забронировать
    </Link>
  ) : (
    <Link
      href="/login"
      className="px-8 py-3 border border-gray-500 text-gray-300 rounded-xl font-medium hover:bg-gray-800 transition"
    >
      Войти, чтобы забронировать
    </Link>
  )}
</div>
      </div>
    </div>
  );
}
