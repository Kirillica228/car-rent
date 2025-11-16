"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import apiFetch from "../../../utils/api";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table"
import { useCars } from "@/hooks/useCars";
import { useQueryClient } from '@tanstack/react-query';

export default function AdminCarPage() {
  const [selected, setSelected] = useState<number[]>([]);
  const router = useRouter();
  const { data: cars, isLoading } = useCars();
  const queryClient = useQueryClient();

  const toggleSelect = (id: number) => {
    setSelected((prev) =>
      prev.includes(id) ? prev.filter((x) => x !== id) : [...prev, id]
    );
  };

  const toggleSelectAll = (checked: boolean) => {
    if (checked) {
      setSelected((cars || []).map((c) => c.id));
    } else {
      setSelected([]);
    }
  };

  const handleDelete = async () => {
    if (selected.length === 0) return;
    if (!confirm(`Удалить ${selected.length} машин(ы)?`)) return;

    try {
      await apiFetch("/admin/cars/delete", {
        method: "DELETE",
        body: JSON.stringify({ id: selected}),
      });
      setSelected([]);
      queryClient.invalidateQueries({ queryKey: ['cars'] });
    } catch (err) {
      console.error("Ошибка удаления:", err);
    }
  };

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-4xl font-extrabold text-black">Машины</h1>
        <div className="flex gap-4">
          {selected.length > 0 && (
            <button
              onClick={handleDelete}
              className="bg-gradient-to-r from-red-600 to-pink-600 hover:from-red-700 hover:to-pink-700 text-white font-bold px-6 py-3 rounded-xl shadow-lg transition transform hover:scale-105"
            >
              🗑️ Удалить ({selected.length})
            </button>
          )}
          <button
            onClick={() => router.push("/admin/car/add")}
            className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-bold px-6 py-3 rounded-xl shadow-lg transition transform hover:scale-105"
          >
            ➕ Добавить
          </button>
        </div>
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead><input type="checkbox" checked={selected.length === (cars || []).length} onChange={(e) => toggleSelectAll(e.target.checked)}/></TableHead>
            <TableHead>id</TableHead>
            <TableHead>brand</TableHead>
            <TableHead>type</TableHead>
            <TableHead>model</TableHead>
            <TableHead>year</TableHead>
            <TableHead>license_plate</TableHead>
            <TableHead>status</TableHead>
            <TableHead>price</TableHead>
            <TableHead>isVisible</TableHead>
            <TableHead>color</TableHead>
            <TableHead>description</TableHead>
            <TableHead>transmission</TableHead>
          </TableRow>
        </TableHeader>
        
        <TableBody>
          {(cars || []).map((car, i) => (
          <TableRow
            key={car.id}
            onClick={() => router.push(`/admin/car/${car.id}`)}
          >
            <TableCell >
              <input 
              type="checkbox" 
              checked={selected.includes(car.id)} 
              onClick={(e) => {
              e.stopPropagation(); 
              toggleSelect(car.id);
              }}  />
            </TableCell>
            <TableCell >{car.id}</TableCell>
            <TableCell>{car.brand.name}</TableCell>
            <TableCell>{car.type.name}</TableCell>
            <TableCell>{car.model}</TableCell>
            <TableCell>{car.year}</TableCell>
            <TableCell>{car.licensePlate}</TableCell>
            <TableCell>{car.status}</TableCell>
            <TableCell>{car.price}</TableCell>
            <TableCell>{car.isVisible ? "true" : "false"}</TableCell>
            <TableCell>{car.color}</TableCell>
            <TableCell>{car.description}</TableCell>
            <TableCell>{car.transmission}</TableCell>
          </TableRow>
          ))}
        </TableBody>
      </Table>
      
    </div>
  );
}
