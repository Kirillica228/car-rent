import Checkbox from "./components/checkbox";
import PriceRangeSlider from "./components/rangeprice"

export default function Shop(){
    return (
        <div className="flex justify-between items-start w-[100%] mt-[-30px] mb-[-30px]">
            <div className="bg-white p-7 w-[30%] h-[100%] flex flex-col items-start gap-[20px]">
                <div className="flex flex-col gap-[10px]">
                    <h1 className="font-medium text-[14px] text-[#596780]">Тип</h1>
                    <div className="flex items-start gap-[5px] flex-col">
                        <Checkbox name="Спорт" number={2}></Checkbox>
                        <Checkbox name="ГУУУУУУУУУД"></Checkbox>
                    </div>
                </div>
                <div className="flex flex-col gap-[10px]">
                    <h1 className="font-medium text-[14px] text-[#596780]">Вместимость</h1>
                    <div className="flex items-start gap-[5px] flex-col">
                        <Checkbox name="2 чел"></Checkbox>
                        <Checkbox name="4 чел"></Checkbox>
                    </div>
                </div>
                <div>
                    <h1>Цена</h1>
                    <PriceRangeSlider></PriceRangeSlider>
                </div>
            </div>
            <div className="flex flex-col items-center w-[70%] mt-[30px]">
                <div className="flex items-center relative justify-center gap-[40px]">
                    <div className="bg-white p-7 rounded-[10px] flex flex-col items-start gap-[10px]">
                        <div className="flex items-center gap-[10px]">
                            <img src="/mark.png" alt="" />
                            <h1>Pick</h1>
                        </div>
                        <div className="flex gap-[15px]">
                            <div>
                                <h1>
                                    Дата
                                </h1>
                                <button>Выбирите дату</button>
                            </div>
                            <div className="bg-[#e4e4e4] w-[2px]"></div>
                            <div>
                                <h1>
                                    Время
                                </h1>
                                <button>Выбирите время</button>
                            </div>
                        </div>
                    </div>
                    <button className="absolute">
                        <img src="/Swap.png" alt="" className="bg-[#3563E9] p-5 rounded-[10px]"/>
                    </button>
                    <div className="bg-white p-7 rounded-[10px] flex flex-col items-start gap-[10px]">
                        <div className="flex items-center gap-[10px]">
                            <img src="/mark.png" alt="" />
                            <h1>Pick</h1>
                        </div>
                        <div className="flex gap-[15px]">
                            <div>
                                <h1>
                                    Дата
                                </h1>
                                <button>Выбирите дату</button>
                            </div>
                            <div className="bg-[#e4e4e4] w-[2px]"></div>
                            <div>
                                <h1>
                                    Время
                                </h1>
                                <button>Выбирите время</button>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="flex items-start w-[100%] p-9">
                    <div className="bg-white p-4 w-[300px] flex flex-col gap-[50px] rounded-[10px] items-center">
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
                        <div className="w-[100%] gap-[15px] flex flex-col">
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
                                <div className="flex gap-[5px] items-end">
                                    <h1 className="font-semibold text-[20px]">1000</h1>
                                    <p>₽/день</p>
                                </div>
                                <button className="bg-[#3563E9] px-1 py-1 text-white rounded-[5px]">Забронировать</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}