"use client";
import { useState } from "react";

const images = [
    "https://sun9-46.userapi.com/impg/e9wrXZvFBPu02CIbe6NfpzAAOVGtfQPIAt3DGQ/HbDul2iQJ9k.jpg?size=807x538&quality=96&sign=fc79c11d05bc947d1536acb5b4589309&c_uniq_tag=cxERSl6mfz0DnkZd4jB_8q-bx4cJtr_2-K9ga2r74AE&type=album",
    "https://sun9-46.userapi.com/impg/e9wrXZvFBPu02CIbe6NfpzAAOVGtfQPIAt3DGQ/HbDul2iQJ9k.jpg?size=807x538&quality=96&sign=fc79c11d05bc947d1536acb5b4589309&c_uniq_tag=cxERSl6mfz0DnkZd4jB_8q-bx4cJtr_2-K9ga2r74AE&type=album",
    "https://sun9-46.userapi.com/impg/e9wrXZvFBPu02CIbe6NfpzAAOVGtfQPIAt3DGQ/HbDul2iQJ9k.jpg?size=807x538&quality=96&sign=fc79c11d05bc947d1536acb5b4589309&c_uniq_tag=cxERSl6mfz0DnkZd4jB_8q-bx4cJtr_2-K9ga2r74AE&type=album"
];

export default function ImageSlider() {
  const [current, setCurrent] = useState(0);



  return (
    <div className="relative w-full mx-auto">
      {/* Слайдер */}
      <div className="relative overflow-hidden w-[100%]">
        <div
          className="flex transition-transform duration-500 ease-in-out "
          style={{ transform: `translateX(-${current * 100}%)` }}
        >
          {images.map((src, index) => (
            <div key={index} className="min-w-full flex flex-col items-center justify-center p-3 bg-black">
              <img
                src={src}
                alt={`Slide ${index + 1}`}
                className="w-full h-auto rounded-lg object-cover"
              />
            </div>
          ))}
        </div>
      </div>

     

      {/* Мини-просмотр следующих фото */}
      <div className="flex justify-center mt-4 space-x-2">
        {images.map((src, index) => (
          <div
            key={index}
            className={`w-[300px] h-[100px] rounded-md overflow-hidden border ${
              index === current ? "border-blue-500 p-1" : "border-gray-300"
            } cursor-pointer`}
            onClick={() => setCurrent(index)}
          >
            <img
              src={src}
              className="w-full h-full object-cover"
            />
          </div>
        ))}
      </div>
    </div>
  );
}
