import {GrFavorite} from "react-icons/gr"
import {FiSearch} from "react-icons/fi"
import {MdAccountCircle} from "react-icons/md"
import {MdOutlineAccountCircle} from "react-icons/md"
import { useContext, useState } from "react"
import { ModalContext } from "../context/ModalContext"
import {AiOutlineShopping} from "react-icons/ai"
import Login from "./Login"
import { Link } from "react-router-dom"
import { UserContext } from "../context/UserContext"

const Navbar = () => {
    const [modalState, modalDispatch] = useContext(ModalContext)
    const [userState, _] = useContext(UserContext)
    const [showLogin, setShowLogin] = useState(false)

    console.log("modalState login : ", modalState)

    return (
        <>
            <div className="flex justify-between bg-green-600 items-center">
                <p className="font-semibold w-36 ms-4">Contact me</p>
                <Link to={'/'}>
                    <p className="text-2xl font-semibold py-4 w-[100%] flex justify-center ms-14">Mvtware.id</p>
                </Link>
                <div className="text-xl">
                    <ul className="flex justify-between w-64 me-4 items-center">
                        <li>
                            <a href="#"><GrFavorite/></a>
                        </li>
                        <li>
                            <FiSearch/>
                        </li>
                        <li>
                            <div className="flex justify-between me-8 items-center">
                                <p className="me-2"><AiOutlineShopping/></p>
                            </div>
                        </li>
                        {/* if user not login */}
                        {!userState.isLogin && (
                         <li>
                            <button onClick={() => modalDispatch({type: 'LOGIN_MODAL'})}>Sign In</button>
                         </li> 
                        )}

                        {/* if user login */}
                        {userState.isLogin && (
                        <li>
                            <MdAccountCircle/>
                        </li> 
                        )}
                    </ul>  
                </div>
            </div>
            {modalState.isModalLogin && (
                <Login/>
            )}
        </>
    )
}

export default Navbar