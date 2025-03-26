"use client";

import { useState } from "react";

export default function PriceRangeSlider() {
  const [value, setValue] = useState(50); // Начальное значение

  const handleChange = (event) => {
    setValue(Number(event.target.value));
  };

  return (
    <div className="flex flex-col items-start gap-2 w-full max-w-md">
      <input
        type="range"
        min="0"
        max="100"
        step="1"
        value={value}
        onChange={handleChange}
        className="w-full h-2 bg-gray-300 rounded-lg appearance-none cursor-pointer accent-blue-500"
        style={{
            background: `linear-gradient(to right, #3b82f6 ${value}%, #d1d5db ${value}%)`,
          }}
      />
      <span className="text-gray-600 text-lg font-semibold">
        До {value.toFixed(1)} ₽
      </span>
    </div>
  );
}
