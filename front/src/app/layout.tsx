"use client";

import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import { QueryProvider } from "./providers";
import { LayoutContent } from "@/components/layoutComponent";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="ru">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-gradient-to-r from-gray-100 via-gray-50 to-gray-100`}
      >
        <QueryProvider>
          <LayoutContent>{children}</LayoutContent>          
        </QueryProvider>
      </body>
    </html>
  );
}
