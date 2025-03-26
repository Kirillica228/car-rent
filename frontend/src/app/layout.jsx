import { Montserrat } from "next/font/google"; 
import "./globals.css";
import Link from "next/link";

const montserrat = Montserrat({ 
  variable: "--font-montserrat", 
  subsets: ["latin", "cyrillic"],
  weight: ["300","400","500","600","700"],
});

export default function RootLayout({ children }) {
  return (
    <html lang="en" className="bg-[#F6F7F9] m-0 p-0">
      <body
        className={`${montserrat.variable}  antialiased flex flex-col items-center`} 
      >
        <header className="bg-white flex items-center justify-between p-4 w-[100%] mb-[30px]">
          <div className="flex items-center gap-[10em]">
            <Link href="/" className="font-medium text-[32px] text-[#3563E9]">MORENT</Link>
            <div className="relative w-[370px] flex items-center">
              <img src="/search-normal.png" alt="" className="absolute left-3"/>
              <input type="text" className="w-[100%] border-1 pl-[64px] pr-[52px] py-[6px] border-[#C3D4E9] rounded-[70px] outline-none" placeholder="Название машины,категории"/>
              <img src="/filter.png" className="absolute right-4" alt="" />
            </div>
          </div>
          <div className="flex gap-[10px]">
            <a className="rounded-[50px] border-[#C3D4E9] border-1 p-1 flex items-center px-1.5">
              <img src="/heart.png" alt=""/>
            </a>
            <a className="rounded-[50px] border-[#C3D4E9] border-1 p-1">
              <img src="/user.png" alt="" />
            </a>
          </div>
          

        </header>
        {children}
        <footer className="bg-white w-[100%] p-7 flex flex-col items-center justify-between mt-[30px]">
          <div className="flex items-start justify-between w-[100%] mt-[50px] p-7">
            <div>
              <Link href="/" className="font-medium text-[32px] text-[#3563E9]">MORENT</Link>
              <p className="font-light text-[18px] text-[#131313]">Наше видение заключается в обеспечении удобства.</p>
            </div>
            <div className="flex gap-[100px]">
              <div className="flex flex-col items-start gap-[20px]">
                <h1 className="font-meduim text-[20px] text-black">О нас</h1>
                <div className="flex flex-col items-start gap-[15px]">
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                </div>
              </div>
              <div className="flex flex-col items-start gap-[20px]">
                <h1 className="font-meduim text-[20px] text-black">О нас</h1>
                <div className="flex flex-col items-start gap-[15px]">
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                </div>
              </div>
              <div className="flex flex-col items-start gap-[20px]">
                <h1 className="font-meduim text-[20px] text-black">О нас</h1>
                <div className="flex flex-col items-start gap-[15px]">
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                  <a href="" className="font-light text-[20px] text-black">бал бала аа</a>
                </div>
              </div>
            </div>
          </div>
          <div className="w-[100%] bg-[#e4e4e4] h-[3px]"></div>
          <div className="flex items-center justify-between w-[100%] p-3">
            <h1 className="font-semibold text-[17px]">©2022 MORENT. All rights reserved</h1>
            <div className="flex items-center gap-[50px]">
              <h1 className="font-semibold text-[16px]">Privacy & Policy</h1>
              <h1 className="font-semibold text-[16px]">Terms & Condition</h1>
            </div>
          </div>
        </footer>
      </body>
    </html>
  );
}
