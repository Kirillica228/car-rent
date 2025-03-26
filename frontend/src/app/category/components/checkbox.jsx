"use client"
import { useState } from "react";

export default function Checkbox({name,number }) {
  const [checked, setChecked] = useState(false);

  return (
    <div className="flex gap-[5px] items-center">
      <div className="flex justify-center items-start">
      <button
        onClick={() => setChecked(!checked)}
        className={`w-5 h-5 rounded-[5px] flex justify-center items-center transition-all duration-200 focus:outline-none
          ${checked ? "bg-blue-500" : "border-2 border-gray-400"}
        `}
      >
        {checked && (
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-4 w-4 text-white"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              fillRule="evenodd"
              d="M16.707 5.293a1 1 0 010 1.414l-7 7a1 1 0 01-1.414 0l-3-3a1 1 0 011.414-1.414L9 11.586l6.293-6.293a1 1 0 011.414 0z"
              clipRule="evenodd"
            />
          </svg>
        )}
      </button>
    </div>
      <label htmlFor="" className="font-semibold text-[#596780] text-[15px] flex">
        {name}
        <p>{number
          ? `(${number})`
          : ""
          }
        </p>
      </label>
    </div>
    
  );
}