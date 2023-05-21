import Search from "./Search"
import Card from "./Card"

const CardPants = () => {
    return (
        <>
            <div className="w-[100] h-auto mt-1 bg-green-600">
                <div className="w-[100%]">
                    <Search/>
                </div>
                <div className="m-auto w-[20%]">
                    <p className="flex justify-center text-white font-bold text-4xl">PANTS</p>
                    <hr className="py-1 rounded-lg bg-black"/>
                </div>
                <div className="flex justify-center mt-12">
                    <div className="grid grid-cols-5 gap-8">
                        <Card/>
                        <Card/>
                        <Card/>
                        <Card/>
                        <Card/>
                    </div>
                </div>
            </div>
        </>
    )
}

export default CardPants