"use client";

import { useState, useEffect } from "react";
import Image from "next/image";
import { useDropzone } from "react-dropzone";

type ExistingImage = { type: "existing"; src: string };
type NewImage = { type: "new"; file: File; src: string };
type ImageItem =
  | { type: "existing"; src: string }            // уже загруженные фото
  | { type: "new"; file: File; src: string };   

type ImageUploaderProps = {
  initialImages?: string[]; // ссылки на существующие фото с бека
  onChange?: (newFiles: File[], removedExisting: string[]) => void; // callback для формы
};

export default function ImageUploader({ initialImages = [], onChange }: ImageUploaderProps) {
  const [existingImages, setExistingImages] = useState<string[]>(initialImages);
  const [newFiles, setNewFiles] = useState<File[]>([]);
  const [removedExisting, setRemovedExisting] = useState<string[]>([]);

  const { getRootProps, getInputProps } = useDropzone({
    accept: { "image/*": [] },
    onDrop: (acceptedFiles) => {
      setNewFiles((prev) => [...prev, ...acceptedFiles]);
    },
  });

  const allImages: ImageItem[] = [
  ...existingImages.map((url) => ({ type: "existing" as const, src: url })),
  ...newFiles.map((file) => ({
    type: "new" as const,
    file,
    src: URL.createObjectURL(file),
  })),
];

  // уведомляем родителя о изменениях
  useEffect(() => {
    if (onChange) {
      onChange(newFiles, removedExisting);
    }
  }, [newFiles, removedExisting, onChange]);

  const handleRemove = (item: ImageItem) => {
    if (item.type === "existing") {
      setExistingImages((prev) => prev.filter((url) => url !== item.src));
      setRemovedExisting((prev) => [...prev, item.src]);
    } else {
      setNewFiles((prev) => prev.filter((f) => f !== item.file));
    }
  };

  return (
    <div>
      <button onClick={() => console.log(allImages)}>ddddddddddd</button>
      {/* Дропзона для новых фото */}
      <div
        {...getRootProps()}
        className="p-6 border-2 border-dashed border-gray-400 rounded-xl text-center cursor-pointer hover:border-gray-600 transition"
      >
        <input {...getInputProps()} />
        <p>Перетащите фото или нажмите, чтобы выбрать</p>
      </div>

      {/* Превью всех фото */}
      {allImages.length > 0 && (
        <div className="mt-4 grid grid-cols-2 md:grid-cols-3 gap-4">
          {allImages.map((item, idx) => (
            <div key={idx} className="relative group">
              {item.type === "existing" && (
                  <img
                      src={`http://localhost:8080/api/catalog${item.src}`}
                      alt={`image-${idx}`}
                      className="rounded-xl object-cover w-full h-32"
                    />
                )}
              {item.type === "new" && (
                  <img
                      src={item.src}
                      alt={`image-${idx}`}
                      className="rounded-xl object-cover w-full h-32"
                    />
                )}
              <button
                type="button"
                onClick={() => handleRemove(item)}
                className="absolute top-1 right-1 bg-red-500 text-white rounded-full px-2 py-1 text-sm opacity-80 hover:opacity-100 transition"
              >
                ❌
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
