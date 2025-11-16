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

type CarType = {
  ID: number;
  name: string;
  is_visible: boolean;
};

export default function AdminTypePage() {
  const [types, setTypes] = useState<CarType[]>([]);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<number[]>([]);
  const router = useRouter();

  async function fetchTypes() {
    setLoading(true);
    try {
      const data = await apiFetch("/admin/types/list", { method: "GET" });
      setTypes(Array.isArray(data) ? data : []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchTypes();
  }, []);

  const toggleSelect = (id: number) => {
    setSelected((prev) =>
      prev.includes(id) ? prev.filter((x) => x !== id) : [...prev, id]
    );
  };

  const toggleSelectAll = (checked: boolean) => {
    if (checked) {
      setSelected(types.map((t) => t.ID));
    } else {
      setSelected([]);
    }
  };

  const handleDelete = async () => {
    if (selected.length === 0) return;
    if (!confirm(`Удалить выбранные ${selected.length} типов?`)) return;

    try {
      await apiFetch("/admin/types/delete", {
        method: "DELETE",
        body: JSON.stringify(selected),
      });
      setSelected([]);
      fetchTypes();
    } catch (err) {
      console.error("Ошибка удаления:", err);
    }
  };

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-4xl font-extrabold text-black">Типы</h1>
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
            onClick={() => router.push("/admin/type/add")}
            className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-bold px-6 py-3 rounded-xl shadow-lg transition transform hover:scale-105"
          >
            ➕ Добавить
          </button>
        </div>
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead><input type="checkbox" checked={selected.length === types.length} onChange={(e) => toggleSelectAll(e.target.checked)}/></TableHead>
            <TableHead>Id</TableHead>
            <TableHead>Name</TableHead>
            <TableHead>IsVisible</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {types.map((type, i) => (
            
          <TableRow key={type.ID}
                  onClick={() => router.push(`/admin/type/${type.ID}`)}
                  className={`transition ${
                    i % 2 === 0 ? "bg-white/0" : "bg-white/5"
                  } hover:bg-white/20 cursor-pointer`}>
            <TableCell>
              <input 
              type="checkbox" 
              checked={selected.includes(type.ID)} 
              onClick={(e) => {
              e.stopPropagation(); 
              toggleSelect(type.ID);
              }}  />
            </TableCell>
            <TableCell>{type.ID}</TableCell>
            <TableCell>{type.name}</TableCell>
            <TableCell>{type.is_visible ? "true" : "false"}</TableCell>
            <TableCell></TableCell>
          </TableRow>
          ))}
        
        </TableBody>
      </Table>
    </div>
  );
}
