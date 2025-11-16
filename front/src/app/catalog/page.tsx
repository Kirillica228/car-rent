'use client';

import { useState, useMemo } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import Cookies from 'js-cookie';
import Link from 'next/link';

import ICar from '@/types/ICar';
import { useBrands } from '@/hooks/useBrands';
import { useTypes } from '@/hooks/useTypes';
import { useCars } from '@/hooks/useCars';

export default function Catalog() {
  const router = useRouter();
  const searchParams = useSearchParams();

  const token = Cookies.get('token');

  // ---------------- Фильтры и даты ----------------
  const [selectedBrands, setSelectedBrands] = useState<number[]>(
    searchParams.get('brand_id')?.split(',').map(Number) || []
  );
  const [selectedTypes, setSelectedTypes] = useState<number[]>(
    searchParams.get('type_id')?.split(',').map(Number) || []
  );
  const [sort, setSort] = useState(searchParams.get('sort') || 'price-asc');
  const [startDate, setStartDate] = useState<string>('');
  const [endDate, setEndDate] = useState<string>('');

  // ---------------- Данные через хуки ----------------
  const { data: brands } = useBrands();
  const { data: types } = useTypes();

  const query = useMemo(() => {
    const params = new URLSearchParams();
    selectedBrands.forEach(b => params.append('brand_id', b.toString()));
    selectedTypes.forEach(t => params.append('type_id', t.toString()));
    if (sort) params.set('sort', sort);
    return params.toString();
  }, [selectedBrands, selectedTypes, sort]);

  const { data: carsData, isLoading: carsLoading } = useCars(); 
  // const { data: busyCars } = useBusyCars(startDate, endDate);

  // const cars = useMemo(() => {
  //   if (!carsData) return [];
  //   const busyIds = busyCars || [];
  //   return carsData.filter(car => !busyIds.includes(car.ID));
  // }, [carsData, busyCars]);

  // ---------------- Фильтры ----------------
  const toggleBrand = (id: number) =>
    setSelectedBrands(prev => (prev.includes(id) ? prev.filter(b => b !== id) : [...prev, id]));

  const toggleType = (id: number) =>
    setSelectedTypes(prev => (prev.includes(id) ? prev.filter(t => t !== id) : [...prev, id]));

  // ---------------- Применение фильтров ----------------
  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const params = new URLSearchParams();
    selectedBrands.forEach(b => params.append('brand_id', b.toString()));
    selectedTypes.forEach(t => params.append('type_id', t.toString()));
    params.set('sort', sort);
    if (startDate) params.set('start_date', startDate);
    if (endDate) params.set('end_date', endDate);
    router.push(`/catalog?${params.toString()}`);
  };

  // ---------------- Рендер ----------------
  return (
    <div className="min-h-screen bg-gray-900 text-white">
      {/* Фильтры */}
      <section className="border-b border-gray-700 bg-gray-900/90 sticky top-0 z-10">
        <form onSubmit={handleSubmit}>
          <div className="max-w-7xl mx-auto px-6 py-6 flex flex-col md:flex-row md:items-center md:justify-between gap-6">
            <div className="flex flex-wrap gap-8">
              {/* Бренды */}
              <div>
                <h3 className="text-sm font-semibold text-gray-300 mb-2 uppercase tracking-wider">Бренды</h3>
                <div className="flex flex-wrap gap-3">
                  {brands?.map(b => (
                    <button
                      key={b.id}
                      type="button"
                      onClick={() => toggleBrand(b.id)}
                      className={`px-3 py-1 rounded-full border text-sm transition-all ${
                        selectedBrands.includes(b.id)
                          ? 'bg-white text-black border-white'
                          : 'border-gray-600 text-gray-300 hover:border-white hover:text-white'
                      }`}
                    >
                      {b.name}
                    </button>
                  ))}
                </div>
              </div>

              {/* Типы */}
              <div>
                <h3 className="text-sm font-semibold text-gray-300 mb-2 uppercase tracking-wider">Типы</h3>
                <div className="flex flex-wrap gap-3">
                  {types?.map(t => (
                    <button
                      key={t.id}
                      type="button"
                      onClick={() => toggleType(t.id)}
                      className={`px-3 py-1 rounded-full border text-sm transition-all ${
                        selectedTypes.includes(t.id)
                          ? 'bg-white text-black border-white'
                          : 'border-gray-600 text-gray-300 hover:border-white hover:text-white'
                      }`}
                    >
                      {t.name}
                    </button>
                  ))}
                </div>
              </div>

              {/* Даты */}
              <div className="flex flex-col gap-2">
                <label className="text-sm text-gray-400">С</label>
                <input
                  type="date"
                  value={startDate}
                  onChange={e => setStartDate(e.target.value)}
                  className="border border-gray-600 rounded-md bg-gray-800 text-gray-100 px-3 py-2 text-sm"
                />
              </div>
              <div className="flex flex-col gap-2">
                <label className="text-sm text-gray-400">По</label>
                <input
                  type="date"
                  value={endDate}
                  onChange={e => setEndDate(e.target.value)}
                  className="border border-gray-600 rounded-md bg-gray-800 text-gray-100 px-3 py-2 text-sm"
                />
              </div>
            </div>

            {/* Сортировка */}
            <div className="flex items-center gap-3">
              <select
                value={sort}
                onChange={e => setSort(e.target.value)}
                className="border border-gray-600 rounded-md bg-gray-800 text-gray-100 px-3 py-2 text-sm focus:ring-2 focus:ring-white focus:outline-none"
              >
                <option value="price-asc">Цене ↑</option>
                <option value="price-desc">Цене ↓</option>
                <option value="year-desc">Году (новые)</option>
                <option value="year-asc">Году (старые)</option>
              </select>
              <button
                type="submit"
                className="bg-white text-black px-4 py-2 rounded-md hover:bg-gray-200 transition-all"
              >
                Применить
              </button>
            </div>
          </div>
        </form>
      </section>

      {/* Контент */}
      <main className="max-w-7xl mx-auto px-6 py-10">
        {carsLoading ? (
          <p className="text-gray-400 text-center py-20">Загрузка...</p>
        ) : carsData?.length === 0 ? (
          <p className="text-gray-400 text-center py-20">Машины по выбранным фильтрам не найдены.</p>
        ) : (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {carsData?.map(car => {
              const imageUrl =
                car.images && car.images.length > 0
                  ? `http://localhost:8080/api/catalog${car.images[0].URL}`
                  : '/placeholder-car.jpg';

              return (
                <Link key={car.id} href={`/catalog/${car.id}`} className="group block">
                  <div className="relative bg-gray-800 border border-gray-700 rounded-2xl overflow-hidden hover:border-gray-500 transition-all duration-300">
                    <div className="relative h-56 overflow-hidden">
                      <img
                        src={imageUrl}
                        alt={car.model}
                        className="w-full h-full object-cover object-center transition-transform duration-500 group-hover:scale-110 group-hover:brightness-90"
                      />
                      <div className="absolute inset-0 bg-gradient-to-t from-black/80 via-black/20 to-transparent opacity-70 group-hover:opacity-80 transition-opacity"></div>
                      <div className="absolute bottom-4 left-4 right-4">
                        <h3 className="text-2xl font-semibold text-white">
                          {car.brand.name} {car.model}
                        </h3>
                      </div>
                    </div>
                    <div className="p-5 flex flex-col gap-2">
                      <p className="text-gray-400 text-sm">
                        {car.type.name} • {car.year}
                      </p>
                      <div className="flex justify-between items-center mt-1">
                        <p className="text-white font-semibold text-lg">{car.price} р / день</p>
                        {token && (
                          <span className="text-sm font-medium text-black bg-white px-4 py-2 rounded-md hover:bg-gray-200 transition-all">
                            Забронировать
                          </span>
                        )}
                      </div>
                    </div>
                  </div>
                </Link>
              );
            })}
          </div>
        )}
      </main>
    </div>
  );
}
