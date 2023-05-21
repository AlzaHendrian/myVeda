import { Link } from 'react-router-dom'
import blouse from '../assets/blouse.jpg'

const Card = () => {
    return (
        <>
            <Link to={'/detail-card'}>
                <div className='m-auto'>
                    <img src={blouse} className='rounded-lg w-[200px]'/>
                    <p>Womens NS Dizzylissy Shirt</p>
                    <div className='flex justify-between mt-3 w-48'>
                        <p>Price</p>
                        <p>$5.00</p>
                    </div>
                </div>
            </Link>
        </>
    )
}

export default Card