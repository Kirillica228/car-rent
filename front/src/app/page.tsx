"use client"
import { useEffect } from "react";

export default function Home() {
  useEffect(() => {
    document.title = "CarRent — аренда автомобилей премиум-класса";
  }, []);

  const features = [
    { title: "Премиальные авто", text: "Только тщательно отобранные модели для безупречного комфорта." },
    { title: "Прозрачные условия", text: "Никаких скрытых платежей — всё честно и ясно." },
    { title: "Моментальное бронирование", text: "Бронируйте за пару кликов — без звонков и ожидания." },
  ];

  const plans = [
    { name: "Базовый", price: "от 2500₽ / день", desc: "Для коротких поездок и тест-драйвов." },
    { name: "Комфорт", price: "от 4500₽ / день", desc: "Оптимальный вариант для города и трассы." },
    { name: "Бизнес", price: "от 9000₽ / день", desc: "Премиум-класс для тех, кто выбирает лучшее." },
  ];

  return (
    <div className="flex flex-col min-h-screen bg-white text-black font-sans">

      {/* HERO */}
      <section className="flex flex-col items-center justify-center py-40 px-6 bg-gray-900 text-white text-center">
        <h1 className="text-6xl font-light tracking-tight mb-6">CarRent</h1>
        <p className="text-xl font-light max-w-2xl mb-10">
          Аренда автомобилей бизнес-класса. Без компромиссов. Без ожидания.
        </p>
        <button className="border border-white px-8 py-3 rounded-full text-white hover:bg-white hover:text-black transition-all">
          Забронировать авто
        </button>
      </section>

      {/* О нас */}
      <section className="py-24 px-6 text-center">
        <h2 className="text-4xl font-semibold mb-6">О компании</h2>
        <p className="text-gray-700 max-w-3xl mx-auto leading-relaxed">
          CarRent — это сервис, созданный для тех, кто ценит стиль, скорость и комфорт.  
          Мы предоставляем автомобили премиум-класса для путешествий, встреч и вдохновения.
        </p>
      </section>

      {/* Преимущества */}
      <section className="py-24 bg-gray-900 text-white">
        <h2 className="text-4xl font-semibold text-center mb-16">Наши преимущества</h2>
        <div className="grid md:grid-cols-3 gap-12 max-w-6xl mx-auto px-6">
          {features.map((f, i) => (
            <div key={i} className="border border-gray-700 rounded-xl p-10 hover:bg-white hover:text-black transition-colors">
              <h3 className="text-2xl mb-4 font-medium">{f.title}</h3>
              <p className="text-gray-300">{f.text}</p>
            </div>
          ))}
        </div>
      </section>

      {/* Процесс */}
      <section className="py-24 px-6 bg-white text-center">
        <h2 className="text-4xl font-semibold mb-16">Как это работает</h2>
        <div className="max-w-5xl mx-auto grid md:grid-cols-3 gap-10">
          {["Выбор автомобиля", "Онлайн-бронь", "Получение авто"].map((step, i) => (
            <div key={i} className="p-8 border border-black rounded-xl hover:bg-gray-900 hover:text-white transition-all">
              <div className="text-5xl font-thin mb-4">{i + 1}</div>
              <h3 className="text-xl font-medium mb-2">{step}</h3>
              <p className="text-gray-600">
                {i === 0 && "Выберите подходящую модель из нашего автопарка."}
                {i === 1 && "Забронируйте онлайн за несколько секунд."}
                {i === 2 && "Получите ключи — и наслаждайтесь поездкой."}
              </p>
            </div>
          ))}
        </div>
      </section>

     

      {/* Отзывы */}
      <section className="py-24 bg-white text-center">
        <h2 className="text-4xl font-semibold mb-16">Отзывы клиентов</h2>
        <div className="flex flex-wrap justify-center gap-10 px-6 max-w-6xl mx-auto">
          {[1, 2, 3].map((i) => (
            <div key={i} className="max-w-sm border border-gray-900 rounded-xl p-8 hover:bg-gray-900 hover:text-white transition-all">
              <p className="italic mb-4">
                “Арендовал BMW X5 на выходные. Машина в идеале, обслуживание — топ!”
              </p>
              <p className="font-semibold">— Клиент №{i}</p>
            </div>
          ))}
        </div>
      </section>

      {/* CTA */}
      <section className="py-32 bg-gray-900 text-center text-white">
        <h2 className="text-4xl font-semibold mb-6">Готов к поездке?</h2>
        <p className="text-lg mb-8 text-gray-300">
          Забронируйте автомобиль прямо сейчас и испытайте свободу на скорости.
        </p>
        <button className="border border-white px-8 py-3 rounded-full text-white hover:bg-white hover:text-black transition-all">
          Забронировать сейчас
        </button>
      </section>

    </div>
  );
}
