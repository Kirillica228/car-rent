import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Link from "next/link";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});


export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
      </body>
      <header className="flex items-center justify-around p-5">
        <h1>Car Rent</h1>
        <nav className="flex items-center justify-between w-[80%]">
          <ul className="flex items-center justify-between w-full">
            <li>
              <Link href={"/"}>Home</Link>
            </li>
            <li>
              <Link href={"/login"}>Login</Link>
            </li>
            <li>
              <Link href={"/register"}>Register</Link>
            </li>
          </ul>
        </nav>
      </header>
      <main className="flex flex-col items-center justify-center h-screen">
        {children}
      </main>
    </html>
  );
}
