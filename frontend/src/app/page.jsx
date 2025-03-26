import Image from "next/image";

export default function Home() {
  return (
    <div className="">
        <div className="flex gap-[50px]">
          <div className="relative">
            <img src="/fon1.png" alt="" />

              <div className="absolute top-0 w-[300px] flex flex-col items-start gap-[0.5em] p-5">
                <h1 className="text-white font-medium text-[20px]">Лучшая платформа для проката автомобилей</h1>
                <p className="text-white font-light text-[15px]">Простота и надежность в аренде автомобиля. Разумеется, по низкой цене.</p>
                <button className="font-medium text-white text-[15px] bg-[#3563E9] p-2 rounded-[10px]">Забронировать</button>
              </div>
              <div className="absolute bottom-0 right-0 p-5">
              <img src="/image 7.png" alt="" className=""/>
              </div>

          </div>
          <div className="relative">
            <img src="/fon2.png" alt="" />

              <div className="absolute top-0 w-[350px] flex flex-col items-start gap-[0.5em] p-5">
                <h1 className="text-white font-medium text-[20px]">Простой способ арендовать автомобиль по низкой цене</h1>
                <p className="text-white font-light text-[15px]">Предоставление недорогих услуг проката автомобилей, а также безопасных и комфортных условий.</p>
                <button className="font-medium text-white text-[15px] bg-[#54A6FF] p-2 rounded-[10px]">Забронировать</button>
              </div>
              <div className="absolute bottom-0 right-0 p-5">
              <img src="/image 8.png" alt="" className=""/>
              </div>
          </div>

        </div>
        <div className="flex flex-col items-start gap-[20px] mt-[30px]">
          <h1>Популярные машины</h1>
          <div className="bg-white p-4 w-[300px] flex flex-col gap-[30px] rounded-[10px] items-center">
            <div className="w-[100%]">
              <div className="flex items-center justify-between">
                <h1>Koenigsegg</h1>
                <button>
                  <img src="/heart.png" alt="" />
                </button>
              </div>
                <p>Sport</p>
            </div>
            <div className="relative">
              <img src="/car.png" alt=""/>
              <img src="/shadow.png" alt="" className="absolute top-[30%]"/>
            </div>
            <div className="flex items-center justify-between w-[100%]">
              <div className="flex items-center gap-[5px]">
                <img src="/benz.png" alt="" />
                <h1>90L</h1>
              </div>
              <div className="flex items-center gap-[5px]">
                <img src="/man.png" alt="" />
                <h1>Manual</h1>
              </div>
              <div className="flex items-center gap-[5px]">
                <img src="/people.png" alt="" />
                <h1>2 чел.</h1>
              </div>
            </div>
            <div className="flex items-center justify-between w-[100%]">
              <h1>1000/день</h1>
              <button className="bg-[#3563E9] px-1 py-1 text-white rounded-[5px]">Забронировать</button>
            </div>

          </div>
          
        </div>
        <div className="flex flex-col items-start gap-[20px] mt-[30px]">
          <h1>Рекомендуемые машины</h1>
          <div className="bg-white p-4 w-[300px] flex flex-col gap-[30px] rounded-[10px]">
            <div className="">
              <div className="flex items-center justify-between">
                <h1>Koenigsegg</h1>
                <button>
                  <img src="/heart.png" alt="" />
                </button>
              </div>
                <p>Sport</p>
            </div>
            <div className="relative">
              <img src="/car.png" alt=""/>
              <img src="/shadow.png" alt="" className="absolute top-[30%]"/>
            </div>
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-[5px]">
                <img src="/benz.png" alt="" />
                <h1>90L</h1>
              </div>
              <div className="flex items-center gap-[5px]">
                <img src="/man.png" alt="" />
                <h1>Manual</h1>
              </div>
              <div className="flex items-center gap-[5px]">
                <img src="/people.png" alt="" />
                <h1>2 чел.</h1>
              </div>
            </div>
            <div className="flex items-center justify-between">
              <h1>1000/день</h1>
              <button className="bg-[#3563E9] px-1 py-1 text-white rounded-[5px]">Забронировать</button>
            </div>

          </div>
        </div>
    </div>
    
  );
}
