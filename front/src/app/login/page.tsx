"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);
    const router = useRouter();

    async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setError("");
    setLoading(true);

    try {
      const res = await fetch("http://localhost:8080/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });

      if (!res.ok) {
        const errData = await res.json();
        throw new Error(errData.error || "Ошибка входа");
      }

      const data = await res.json();

      document.cookie = `token=${data.token}; path=/; max-age=${7 * 24 * 60 * 60}; SameSite=Lax`;

      if (data.role === 1) {
        router.push("/admin");
      } else {
        router.push("/catalog");
      }
    } catch (err: any) {
      setError(err.message);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="flex flex-col items-center ">
      <main className="flex flex-col gap-[32px] row-start-2 items-center sm:items-start">
        <form onSubmit={handleSubmit} className="flex flex-col gap-[16px]">
          <label htmlFor="email">Эл. почта</label>
          <input
            type="email"
            name="email"
            className="border p-1 rounded-md"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
          />

          <label htmlFor="password">Пароль</label>
          <input
            type="password"
            name="password"
            className="border p-1 rounded-md"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          {error && <p className="text-red-500">{error}</p>}

          <button
            type="submit"
            className="bg-blue-500 text-white p-2 rounded-md mt-2"
            disabled={loading}
          >
            {loading ? "Входим..." : "Login"}
          </button>
        </form>
      </main>
    </div>
  );
}
