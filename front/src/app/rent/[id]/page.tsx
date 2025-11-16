"use client";

import { useEffect, useState } from "react";
import { useParams,useRouter } from "next/navigation";
import apiFetch from "../../../utils/api";


type Images = { ID: number; CarID: number; URL: string };
type Brand = { Name: string };
type Type = { Name: string };

type Car = {
  ID: number;
  Brand: Brand;
  Model: string;
  Year: number;
  Type: Type;
  Price: number;
  Images: Images[];
};

export default function RentPage() {
  const params = useParams();
  const carId = Number(params?.id); // получаем id из URL
  const router = useRouter();
  const [car, setCar] = useState<Car | null>(null);
  const [loadingCar, setLoadingCar] = useState(true);

  // Даты
  const [startDate, setStartDate] = useState<string | null>(null);
  const [endDate, setEndDate] = useState<string | null>(null);
  const [selectedYear, setSelectedYear] = useState<number>(2025);
  const [selectedMonth, setSelectedMonth] = useState<number>(new Date().getMonth());
  const [showMonthSelector, setShowMonthSelector] = useState(false);
  const [loading, setLoading] = useState(false);
  const [currentIndex, setCurrentIndex] = useState(0);

  const [bookedDates, setBookedDates] = useState<string[]>([]);
  const monthNames = [
    "Январь","Февраль","Март","Апрель","Май","Июнь",
    "Июль","Август","Сентябрь","Октябрь","Ноябрь","Декабрь"
  ];

  // --- Загрузка машины ---
  useEffect(() => {
    if (!carId) return;
    (async () => {
      try {
        // Машина
        const data = await apiFetch(`/catalog/cars/${carId}`, { method: "GET" });
        setCar(data);

        // Забронированные даты
        const booked = await apiFetch(`/rents/bookings/${carId}`, { method: "GET" });
        // Преобразуем в строки YYYY-MM-DD
        const dates: string[] = booked.map((b: any) => {
          const start = new Date(b.start_date);
          const end = new Date(b.end_date);
          const arr: string[] = [];
          for (let d = start; d <= end; d.setDate(d.getDate() + 1)) {
            arr.push(d.toISOString().split("T")[0]);
          }
          return arr;
        }).flat();
        setBookedDates(dates);

      } catch (err) {
        console.error("Ошибка загрузки машины или бронирований:", err);
      } finally {
        setLoadingCar(false);
      }
    })();
  }, [carId]);

  if (loadingCar) {
    return <div className="min-h-screen flex items-center justify-center bg-gray-900 text-white">Загрузка машины...</div>;
  }

  if (!car) {
    return <div className="min-h-screen flex items-center justify-center bg-gray-900 text-red-500">Машина не найдена</div>;
  }

  // --- Календарь ---
  const daysInMonth = new Date(selectedYear, selectedMonth + 1, 0).getDate();
  const dates = Array.from({ length: daysInMonth }, (_, i) => {
    const d = new Date(selectedYear, selectedMonth, i + 1);
    return d.toISOString().split("T")[0];
  });

  const today = new Date();
  today.setHours(0, 0, 0, 0);

  const isPast = (date: string) => new Date(date) < today;
  const isInRange = (date: string) => startDate && endDate ? new Date(date) >= new Date(startDate) && new Date(date) <= new Date(endDate) : false;
  const isRangeValid = (start: string, end: string) => !dates.filter(d => new Date(d) >= new Date(start) && new Date(d) <= new Date(end)).some(d => bookedDates.includes(d));

  const handleSelectDate = (date: string) => {
    if (isPast(date) || bookedDates.includes(date)) return;
    if (!startDate || (startDate && endDate)) {
      setStartDate(date);
      setEndDate(null);
    } else if (startDate && !endDate) {
      if (new Date(date) < new Date(startDate)) {
        setStartDate(date);
      } else {
        if (isRangeValid(startDate, date)) setEndDate(date);
        else alert("Выбранный диапазон пересекается с забронированными датами!");
      }
    }
  };

  const pricePerDay = car.Price;
  let total = 0;
  if (startDate && endDate) {
    const start = new Date(startDate);
    const end = new Date(endDate);
    const days = Math.ceil((end.getTime() - start.getTime()) / (1000*60*60*24)) + 1;
    total = days * pricePerDay;
  }

  const handleSubmit = async () => {
    if (!startDate || !endDate) return;
    setLoading(true);
    try {
      await apiFetch("/rents/save", {
        method: "POST",
        body: JSON.stringify({
          car_id: car.ID,
          start_date: new Date(startDate).toISOString(),
          end_date: new Date(endDate).toISOString(),
        }),
      });
      router.push("/profile");
      alert(`✅ Успешно забронировано!\nПериод: ${startDate} - ${endDate}\nСтоимость: ${total} р`);
      setStartDate(null);
      setEndDate(null);
    } catch (err) {
      console.error(err);
      alert("❌ Ошибка при бронировании");
    } finally {
      setLoading(false);
    }
  };

  const images = car.Images.length > 0 ? car.Images : [{ ID: 0, CarID: car.ID, URL: "/placeholder-car.jpg" }];
  const prevSlide = () => setCurrentIndex(prev => prev === 0 ? images.length-1 : prev-1);
  const nextSlide = () => setCurrentIndex(prev => prev === images.length-1 ? 0 : prev+1);

  return (
   <div className="min-h-screen bg-gray-900 text-white py-16 px-6">
      <div className="max-w-6xl mx-auto space-y-10">
        {/* Фото + Инфо + Календарь */}
        <div className="flex flex-col lg:flex-row gap-10">

          {/* Фото и слайдер */}
          <div className="flex-1 bg-gray-800 rounded-3xl overflow-hidden shadow-[0_0_40px_rgba(255,255,255,0.05)] border border-gray-700">
            <div className="relative w-full h-[380px]">
              <img
                src={`http://localhost:8080/api/catalog${images[currentIndex].URL}`}
                alt={`car ${currentIndex}`}
                className="w-full h-full object-cover object-center transition-transform duration-700 hover:scale-105"
              />
              <button onClick={prevSlide} className="absolute top-1/2 left-4 -translate-y-1/2 bg-black/50 px-3 py-2 rounded-full hover:bg-black/70">‹</button>
              <button onClick={nextSlide} className="absolute top-1/2 right-4 -translate-y-1/2 bg-black/50 px-3 py-2 rounded-full hover:bg-black/70">›</button>
              <div className="absolute bottom-4 left-1/2 -translate-x-1/2 flex space-x-2">
                {images.map((_, idx) => (
                  <span key={idx} className={`w-3 h-3 rounded-full ${idx===currentIndex?"bg-white":"bg-gray-500"}`} />
                ))}
              </div>
            </div>
            <div className="p-8">
              <h2 className="text-3xl font-semibold mb-3">{car.Brand.Name} {car.Model}</h2>
              <div className="text-gray-400 space-y-1">
                <p>Год выпуска: {car.Year}</p>
                <p>Тип: {car.Type.Name}</p>
              </div>
              <p className="text-green-400 font-bold text-2xl mt-6">{pricePerDay} р / день</p>
            </div>
          </div>

          {/* Календарь */}
          <div className="flex-1 bg-gray-800 rounded-3xl border border-gray-700 shadow-[0_0_40px_rgba(255,255,255,0.05)] p-8">
            <div className="flex justify-between items-center mb-6">
              <h2 className="text-2xl font-semibold">Даты аренды</h2>
              <div className="flex items-center gap-3">
                <input type="number" value={selectedYear} min={2025} max={2030} onChange={e=>setSelectedYear(Number(e.target.value))} className="p-2 w-24 rounded-xl bg-gray-900 border border-gray-700 text-white text-center"/>
                <button onClick={()=>setShowMonthSelector(!showMonthSelector)} className="px-4 py-2 rounded-xl bg-gray-900 border border-gray-700 hover:bg-gray-700 transition">{monthNames[selectedMonth]}</button>
              </div>
            </div>
            {showMonthSelector && <div className="grid grid-cols-3 gap-2 mb-4">{monthNames.map((name, idx)=>(<button key={idx} onClick={()=>{setSelectedMonth(idx); setShowMonthSelector(false)}} className={`p-2 rounded-lg text-sm font-semibold transition ${idx===selectedMonth?"bg-blue-600":"bg-gray-700 hover:bg-gray-600"}`}>{name}</button>))}</div>}
            <div className="grid grid-cols-7 gap-2">{dates.map(date=>{
              const isBooked = bookedDates.includes(date);
              const isSelected = date===startDate || date===endDate || isInRange(date);
              return (
                <button key={date} onClick={()=>handleSelectDate(date)} disabled={isPast(date)||isBooked} className={`p-2 rounded-lg text-sm transition font-medium ${isPast(date)?"bg-gray-800 text-gray-600 cursor-not-allowed":isBooked?"bg-red-500 text-white cursor-not-allowed":isSelected?"bg-blue-600 text-white":"bg-gray-700 hover:bg-gray-600"}`}>
                  {new Date(date).getDate()}
                </button>
              )
            })}</div>
          </div>
        </div>

        {/* Итог */}
        <div className="bg-gray-800 border border-gray-700 rounded-3xl p-10 shadow-[0_0_40px_rgba(255,255,255,0.05)]">
          <div className="flex flex-col md:flex-row md:items-center md:justify-between gap-6">
            <div>
              <h2 className="text-2xl font-semibold mb-2">Детали бронирования</h2>
              <p className="text-gray-300">Начало: <span className="text-white font-medium">{startDate||"—"}</span></p>
              <p className="text-gray-300">Конец: <span className="text-white font-medium">{endDate||"—"}</span></p>
            </div>
            <div className="text-right">
              <p className="text-green-400 text-3xl font-bold">{total>0?`${total} р`:"0 р"}</p>
              <p className="text-gray-400 text-sm mt-1">{startDate&&endDate?"Итого за период":"Выберите даты"}</p>
            </div>
          </div>
          <button onClick={handleSubmit} disabled={!startDate||!endDate||loading} className="mt-10 w-full bg-white text-black font-semibold text-lg py-4 rounded-xl hover:bg-gray-200 transition disabled:opacity-50 disabled:cursor-not-allowed">
            {loading?"Бронирование...":"Забронировать"}
          </button>
        </div>
      </div>
    </div>
  );
}
