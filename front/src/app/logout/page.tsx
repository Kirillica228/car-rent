"use client";

import Cookies from "js-cookie";
import { useEffect } from "react";
import { useRouter } from "next/navigation";

export default function LogoutPage() {
  const router = useRouter();

  useEffect(() => {
    // Удаляем cookie
    Cookies.remove("token");

    // Редиректим на главную
    router.push("/");
  }, [router]);

  return null; // страница пустая, т.к. сразу редирект
}
