import ImageSlider from "./components/slider";

export default function CartPage(){

    return (
        <div className="p-7">
            <div className="flex flex-col w-[100%] gap-[25px] justify-center">
                <div className="flex w-[100%] gap-[50px] justify-center">
                    <div className="w-[45%]">
                        <ImageSlider></ImageSlider>
                    </div>         
                    <div className="bg-white p-7 rounded-[10px] flex flex-col justify-between items-start w-[55%]">
                        <div className="flex items-start justify-between w-[100%]">
                            <div className="flex flex-col items-start">
                                <h1 className="font-bold text-[20px]">Nissan GT-R</h1>
                                <div className="flex items-center gap-[10px]">
                                    <div>
                                        <h1 className="">тута звезды</h1>
                                    </div>
                                    <h1 className="font-medium text-[14px] text-[#596780]">400 отзывов</h1>
                                </div>
                            </div>
                            <img src="/heart.png" alt="" />
                        </div>
                        <h1 className="font-light text-[17px] text-[#596780] whitespace-">
                            балаыллывлалыфлалвыфлалфвылалыфлалыфвлаылфвлаыфлалыфвалфылвалыф
                        </h1>
                        <div>
                            <div className="flex items-center gap-[10px]">
                                <h1 className="font-light text-[#90A3BF] text-[16px]">Тип</h1>
                                <h1 className="font-medium text-[#90A3BF] text-[16px]">Спорткар</h1>
                            </div>
                        </div>
                        <div className="flex items-center justify-between w-[100%]">
                            <span className="flex items-end">
                                <h1 className="font-medium text-[20px]">800₽</h1>
                                <p className="font-light text-[16px] text-[#90A3BF]">/день</p>
                            </span>
                            <button className="bg-[#3563E9] p-4 text-white text-[16px] rounded-[10px]">
                                Забронировать
                            </button>
                        </div>
                    </div>
                </div>
                <div className="w-[100%] bg-white p-7 rounded-[10px] flex flex-col items-start gap-[30px]">
                    <div className="flex items-center gap-[10px]">
                        <h1 className="font-medium text-[#1A202C] text-[18px]">Отзывы</h1>
                        <span className="bg-[#3563E9] px-2 py-1 rounded-[5px] text-white text-[13px]">
                            13
                        </span>
                    </div>
                    <div className="flex flex-col items-center w-[100%] ">
                        <div className="flex items-start gap-[20px] w-[100%]">
                            <img src="" alt="" className="w-[3em] h-[3em] bg-black rounded-4xl"/>
                            <div className="flex flex-col items-start gap-[20px] w-[100%]">
                                <div className="flex items-center justify-between w-[100%]">
                                    <h1 className="font-medium text-[20px]">ЧЕргык румек аглы</h1>
                                    <div className="flex flex-col items-start gap-[10px]">
                                        <h1>21.21.2121</h1>
                                        <h1>тута звезды</h1>
                                    </div>
                                </div>
                                <h1 className="font-light text-[#596780] text-[16px] break-all">dddddddddааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааааdddddddddddddddddddзывввввввввввввввввввввввввввввввацфуацауцйайцацуацуацуйауйцайцуацйуайцуайцуайцуайцуайцуауйцаццццццццццццццццццццццццццццццццццццццццц</h1>

                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    );
}