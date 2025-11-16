import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export async function middleware(req: NextRequest) {
  const isAdminPath = req.nextUrl.pathname.startsWith("/admin");

  // Если это не админский путь — пропускаем
  if (!isAdminPath) return NextResponse.next();

  const token = req.cookies.get("access_token")?.value;
  console.log(token)
  // 1️⃣ Если токена нет — редирект на /login
  if (!token) {
    return NextResponse.redirect(new URL("/login", req.url));
  }

  // 2️⃣ Проверяем токен
  try {
    const res = await fetch("http://localhost:8080/api/authorize/me", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    // if (res.status === 401) return NextResponse.redirect(new URL("/login", req.url));
    if (res.status === 403) return NextResponse.redirect(new URL("/", req.url));
    if (!res.ok) return NextResponse.redirect(new URL("/error", req.url));

    const user = await res.json();

    // Проверяем роль
    if (!user.role || user.role.toLowerCase() !== "admin") {
      return NextResponse.redirect(new URL("/", req.url));
    }

    return NextResponse.next();
  } catch (err) {
    console.error("Auth check failed:", err);
    return NextResponse.redirect(new URL("/error", req.url));
  }
}

export const config = {
  matcher: ["/admin", "/admin/:path*"], // проверяем только админские маршруты
};
