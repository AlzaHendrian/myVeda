import { useContext } from "react"
import { Link } from "react-router-dom"
import { ModalContext } from "../context/ModalContext"
import Register from "./Register"

const Search = () => {
    const [modalState, modalDispatch] = useContext(ModalContext)
   return (
   <>
        <div className="flex justify-between items-center py-10">
            <input type="text" placeholder="input your product!" className="border rounded-lg px-12 py-2 ms-5 w-[20%]"/>
            
            <ul className="flex justify-between w-[40%]">
                <li><Link to={'/'}>Home</Link></li>
                <li><Link to={'/jacket'}>Jacket</Link></li>
                <li><Link to={'/pants'}>Pants</Link></li>
                <li><Link to={'/all-product'}>All Product</Link></li>
            </ul>
            <div className="flex me-8 items-center">
                <button className="bg-black py-2 px-8 text-white rounded-md font-semibold" onClick={() => modalDispatch({type: 'REGISTER_MODAL'})}>Sign Up</button>
            </div>
        </div>
        {modalState.isModalRegist && (
            <Register/>
        )}
    </>
   )
}

export default Search