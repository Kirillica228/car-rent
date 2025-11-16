"use client";
import { useEffect } from "react";

export default function FindUs() {
  useEffect(() => {
    document.title = "CarRent — где нас найти";
  }, []);

  const branches = [
    {
      city: "Ижевск",
      address: "ул. Красная, 269",
      phone: "+7 (495) 123-45-67",
      hours: "Пн–Вс: 9:00–21:00",
    },
  ];

  return (
    <div className="min-h-screen bg-black text-white font-sans flex flex-col">
      {/* HERO */}
      <section className="flex flex-col justify-center items-center py-32 text-center bg-gradient-to-b from-gray-900 to-neutral-900">
        <h1 className="text-6xl font-light tracking-tight mb-6 uppercase">
          Где нас найти
        </h1>
        <p className="text-gray-400 max-w-2xl text-lg leading-relaxed">
          Один город. Один стиль.  
          Добро пожаловать в <span className="text-white">CarRent</span> — там, где начинается твоя дорога.
        </p>
      </section>

      {/* MAP SECTION */}
      <section className="relative bg-gray-900 py-20">
        <div className="max-w-6xl mx-auto px-6 flex flex-col md:flex-row gap-12 items-center">
          {/* Карта */}
          <div className="flex-1 rounded-2xl overflow-hidden ">
            <iframe
              src="https://yandex.ru/map-widget/v1/?um=constructor%3Af5cba16226714c9949df28ca04f6100809809b43fa7310c05af0269eb0df633e&amp;source=constructor"
              width="100%"
              height="400"
              className="border-0 "
              allowFullScreen
            ></iframe>
          </div>

          {/* Текст справа */}
          <div className="flex-1 text-left space-y-6">
            <h2 className="text-4xl font-light mb-4">Наши офисы</h2>
            <p className="text-gray-400 text-lg leading-relaxed">
              Каждый наш филиал — не просто место, где ты берёшь авто.
              Это точка старта, откуда начинается твоя дорога.
            </p>
            <div className="w-24 h-[2px] bg-white opacity-40"></div>
            <p className="text-gray-500 italic text-sm">
              CarRent — скорость. стиль. уверенность.
            </p>
            {branches.map((b, i) => (
            <div
              key={i}
              className=" "
            >
              <h3 className="text-2xl font-semibold mb-3">{b.city}</h3>
              <p className="mb-1">{b.address}</p>
              <p className="mb-1">{b.phone}</p>
              <p className="text-gray-500">{b.hours}</p>
            </div>
          ))}
          </div>
        </div>
      </section>

    
    </div>
  );
}
