// lib/api.ts
import Cookies from "js-cookie";

const API_URL = "http://localhost:8080/api"; // или вынеси в env

export default async function apiFetch(
  endpoint: string,
  options: RequestInit = {},
) {
  const token = Cookies.get("token");

  const headers: Record<string, string> = {
    ...(options.headers as Record<string, string>),
  };

  const isFormData =
    typeof FormData !== "undefined" &&
    options.body &&
    typeof options.body === "object" &&
    options.body.constructor?.name === "FormData";

  if (!isFormData) {
    headers["Content-Type"] = "application/json";
  }

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const res = await fetch(`${API_URL}${endpoint}`, {
    ...options,
    headers,
    // credentials: "include", // чтобы куки тоже летели
  });

  if (!res.ok) {
    let message = "Ошибка запроса";
    try {
      const data = await res.json();
      message = data.error || message;
    } catch {}
    throw new Error(message);
  }

  const text = await res.text();
  return text ? JSON.parse(text) : null;
}
