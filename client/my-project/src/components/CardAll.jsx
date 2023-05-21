import Search from "./Search"
import Card from "./Card"

const CardAll = () => {
    return (
        <>
            <div className="w-[100] h-auto mt-1 bg-green-600">
                <div className="w-[100%]">
                    <Search/>
                </div>
                <div className="m-auto w-[20%]">
                    <p className="flex justify-center text-white font-bold text-4xl">ALL PRODUCT</p>
                    <hr className="py-1 rounded-lg bg-black"/>
                </div>
                <div> 
                    <p className="text-white font-semibold text-2xl ms-20 mt-12">Jacket</p>
                    <div className="flex justify-center mt-4">
                        <div className="grid grid-cols-5 gap-8">
                            <Card/>
                            <Card/>
                            <Card/>
                            <Card/>
                            <Card/>
                        </div>
                    </div>
                </div>
                <div> 
                    <p className="text-white font-semibold text-2xl ms-20 mt-8">Pants</p>
                    <div className="flex justify-center mt-4">
                        <div className="grid grid-cols-5 gap-8">
                            <Card/>
                            <Card/>
                            <Card/>
                            <Card/>
                            <Card/>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default CardAll