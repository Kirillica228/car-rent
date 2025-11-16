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

type Brand = {
  ID: number;
  name: string;
  is_visible: boolean;
};

export default function AdminBrandPage() {
  const [brands, setBrands] = useState<Brand[]>([]);
  const [loading, setLoading] = useState(false);
  const [selected, setSelected] = useState<number[]>([]);
  const router = useRouter();

  async function fetchBrands() {      
    setLoading(true);
    try {
      const data = await apiFetch("/admin/brands/list", { method: "GET" });
      setBrands(Array.isArray(data) ? data : []);
    } catch (err) {
      console.error(err);
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchBrands();
  }, []);

  const toggleSelect = (id: number) => {
    setSelected((prev) =>
      prev.includes(id) ? prev.filter((x) => x !== id) : [...prev, id]
    );
  };

  const toggleSelectAll = (checked: boolean) => {
    if (checked) {
      setSelected(brands.map((b) => b.ID));
    } else {
      setSelected([]);
    }
  };

  const handleDelete = async () => {
    if (selected.length === 0) return;
    if (!confirm(`Удалить ${selected.length} бренд(ов)?`)) return;

    try {
      await apiFetch("/admin/brands/delete", {
        method: "DELETE",
        body: JSON.stringify(selected),
      });
      setSelected([]);
      fetchBrands();
    } catch (err) {
      console.error("Ошибка удаления:", err);
    }
  };

  return (
    <div className="w-full">
      <div className="flex justify-between items-center mb-8">
       <h1 className="text-4xl font-extrabold text-black"> Бренды</h1>
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
            onClick={() => router.push("/admin/brand/add")}
            className="bg-gradient-to-r from-blue-600 to-purple-600 hover:from-blue-700 hover:to-purple-700 text-white font-bold px-6 py-3 rounded-xl shadow-lg transition transform hover:scale-105"
          >
            ➕ Добавить
          </button>
        </div>
      </div>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead><input type="checkbox" checked={selected.length === brands.length} onChange={(e) => toggleSelectAll(e.target.checked)}/></TableHead>
            <TableHead>Id</TableHead>
            <TableHead>Name</TableHead>
            <TableHead>IsVisible</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {brands.map((brand, i) => (
          <TableRow
            key={brand.ID}
            onClick={() => router.push(`/admin/brand/${brand.ID}`)}
          >
            <TableCell >
              <input 
              type="checkbox" 
              checked={selected.includes(brand.ID)} 
              onClick={(e) => {
              e.stopPropagation(); 
              toggleSelect(brand.ID);
              }}  />
            </TableCell>
            <TableCell className="font-medium">{brand.ID}</TableCell>
            <TableCell>{brand.name}</TableCell>
            <TableCell>{brand.is_visible ? "true" : "false"}</TableCell>
          </TableRow>
          ))}
        </TableBody>
      </Table>
      
      
    </div>
  );
}
