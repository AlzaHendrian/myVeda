import blouse from '../assets/blouse.jpg'
import Card from '../components/Card'
import FavoriteButton from './FavoriteButton'

const DetailCard = () => {
    return (
        <>
            <div className=' bg-green-600 mt-1 py-4'>
                <div className='flex justify-between w-[95%] mx-auto'>
                    <div>
                        <div className='flex justify-center'>
                            <img src={blouse} className='w-[420px] rounded-lg'/>
                        </div>
                        <div className='flex justify-center mt-6'>
                            <div className='grid grid-cols-3 gap-8'>
                                <img src={blouse} className='w-[120px] border-solid border-4 border-black'/>
                                <img src={blouse} className='w-[120px] border-solid border-4 border-black'/>
                                <img src={blouse} className='w-[120px] border-solid border-4 border-black'/>
                            </div>
                        </div>
                    </div>
                    <div className='mx-auto w-[50%]'>
                        <h2 className='font-bold text-4xl text-black mb-6' style={{letterSpacing: '4px'}}>Womens NS Dizzylissy Shirt</h2>
                        <div className='flex justify-between items-center'>
                            <div>
                                <i className='text-2xl text-black mt-6' style={{letterSpacing: '4px'}}>Size : X-large</i>
                                <h2 className='font-bold text-2xl text-black'>$5.99</h2>
                            </div>
                            <div className='text-4xl'>
                                <FavoriteButton/>
                            </div>
                        </div>
                        <div className='grid grid-cols-2 gap-3 mt-6'>
                            <h2 className='bg-black text-white py-4 flex justify-center'>Add to Card</h2>
                            <h2 className='bg-white text-black py-4 flex justify-center'>Buy Now!</h2>
                        </div>
                        <div className='mt-10 text-xl font-semibold'> DETAILS
                            <div className='ms-3 text-lg font-normal mt-4'>
                                <p className='mb-3'>
                                    Womens NS Dizzylissy Shirt
                                </p>
                                <ul className='list-disc ms-6'>
                                    <li className='mb-2'>
                                        Style:
                                    </li>
                                    <li className='mb-2'>
                                        Material:
                                    </li>
                                    <li className='mb-2'>
                                        Color:
                                    </li>
                                    <li className='mb-2'>
                                        Condition:
                                    </li>
                                </ul>
                            </div>
                            
                        </div>
                    </div>
                </div>
            </div>

            <div className='bg-green-600 mt-1 py-4'> 
                <p className="text-white font-semibold text-2xl ms-20 mt-8">YOU MAY ALSO LIKE</p>
                <div className="flex justify-center mt-4">
                    <div className="grid grid-cols-5 gap-8">
                        <Card/>
                        <Card/>
                        <Card/>
                        <Card/>
                        <Card/>
                    </div>
                </div>
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
        </>
    )
}

export default DetailCard