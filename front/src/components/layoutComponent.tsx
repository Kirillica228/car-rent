"use client";

import Link from "next/link";
import { useRouter } from "next/navigation";
import { useLogout, useMe } from "@/hooks/useAuth";

export function LayoutContent({ children }: { children: React.ReactNode }) {
  const router = useRouter();
  const { data, isLoading} = useMe();
  const logoutMutation = useLogout();

  const handleLogout = () => {
    logoutMutation.mutate(undefined, {
      onSuccess: () => {
        // После успешного logout перенаправляем на логин
        window.location.replace("/");
      },
      onError: (err) => {
        console.error('Logout failed', err);
      },
    });
  };
  return (
    <>
      {/* HEADER */}
      <header className="flex flex-col md:flex-row items-center justify-between p-6 bg-gray-900 text-white shadow-lg relative overflow-hidden">
        {/* Фоновые машинки */}
        <div className="absolute top-0 left-0 w-full h-full pointer-events-none">
          <div className="absolute top-0 left-10 w-32 h-20 bg-red-500 rotate-12 rounded-lg animate-car1"></div>
          <div className="absolute top-10 right-10 w-24 h-16 bg-blue-500 -rotate-6 rounded-lg animate-car2"></div>
        </div>

        <Link href="/" className="z-10 font-bold text-3xl md:text-4xl tracking-tight">
          CarRent
        </Link>

        <nav className="z-10 mt-4 md:mt-0">
          <ul className="flex items-center gap-6 text-lg">
            <li>
              <Link href="/" className="hover:text-yellow-400 transition-colors">О нас</Link>
            </li>
            <li>
              <Link href="/catalog" className="hover:text-yellow-400 transition-colors">Каталог</Link>
            </li>
            <li>
              <Link href="/about" className="hover:text-yellow-400 transition-colors">Где нас найти</Link>
            </li>

            {isLoading ? null : data ? (
              <>
                {data.role === "admin" && (
                  <li>
                    <Link href="/admin" className="hover:text-yellow-400 transition-colors">
                      Панель администратора
                    </Link>
                  </li>
                )}
                <li>
                  <Link href="/profile" className="hover:text-yellow-400 transition-colors">
                    Мои бронирования
                  </Link>
                </li>
                <li>
                  <button
                    onClick={() => {
                      
                      handleLogout();
                    }}
                    className="hover:text-yellow-400 transition-colors"
                  >
                    Выйти
                  </button>
                </li>
              </>
            ) : (
              <>
                <li>
                  <Link href="/login" className="hover:text-yellow-400 transition-colors">
                    Войти
                  </Link>
                </li>
                <li>
                  <Link href="/register" className="hover:text-yellow-400 transition-colors">
                    Регистрация
                  </Link>
                </li>
              </>
            )}
          </ul>
        </nav>
      </header>
      {children}

      <footer className="py-10 text-center text-gray-500 bg-white border-t border-gray-200 text-sm">
        © {new Date().getFullYear()} CarRent. Стиль. Скорость. Свобода.
      </footer>

      <style jsx>{`
        @keyframes moveCar1 {
          0% { transform: translateX(-50%) rotate(12deg); }
          50% { transform: translateX(50vw) rotate(12deg); }
          100% { transform: translateX(-50%) rotate(12deg); }
        }
        @keyframes moveCar2 {
          0% { transform: translateX(50%) rotate(-6deg); }
          50% { transform: translateX(-50vw) rotate(-6deg); }
          100% { transform: translateX(50%) rotate(-6deg); }
        }
        .animate-car1 { animation: moveCar1 10s linear infinite; }
        .animate-car2 { animation: moveCar2 12s linear infinite; }
      `}</style>
    </>
  );
}
