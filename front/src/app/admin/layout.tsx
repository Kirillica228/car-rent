import Link from "next/link";

export default function AdminLayout({ children }: { children: React.ReactNode }){
    return (

        <div className="flex min-h-screen">
            {/* Сайдбар */}
            <aside className="w-64 bg-gray-900 text-white flex flex-col p-6 shadow-lg">
            <h2 className="text-2xl font-bold mb-8">Admin Panel</h2>
            <nav className="flex flex-col gap-4">
                <Link
                href="/admin/car"
                className="px-4 py-2 rounded hover:bg-gray-700 transition flex items-center gap-2"
                >
                🚗 Car
                </Link>
                <Link
                href="/admin/brand"
                className="px-4 py-2 rounded hover:bg-gray-700 transition flex items-center gap-2"
                >
                🏷️ Brand
                </Link>
                <Link
                href="/admin/type"
                className="px-4 py-2 rounded hover:bg-gray-700 transition flex items-center gap-2"
                >
                🛠️ Type
                </Link>
            </nav>
            </aside>

            {/* Основной контент */}
            <main className="flex-1 flex flex-col items-center justify-start py-10 px-6 min-h-screen">
            {children}
            </main>
        </div>


    );
}