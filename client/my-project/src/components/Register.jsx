import { useContext, useState } from "react";
import { ModalContext } from "../context/ModalContext";
import { API } from "../config/Api";

const Register = () => {
    const [message, setMessage] = useState(null)
    const [_, modalDispatch] = useContext(ModalContext)
    const [form, setForm] = useState({
        email: '',
        password: '',
        fullname: '',
        phone: '',
        address: '',
    })

    const {email, password, fullname, phone, address} = form

    const handleOnChange = (e) => {
        setForm({
            ...form,
            [e.target.name]: e.target.value,
        })
    }

    const handleSubmit = async (e) => {
        try {
            e.preventDefault();
            const response = await API.post("/register", form)
            console.log("register success :", response)
            alert("successfully register!")
            setForm({
                email: "",
                password: "",
                fullname: "",
                phone: "",
                address: "",
            })
        }catch(err) {
            alert("failed registered!")
        }
    }

  return (
    <div className="absolute w-[100%] mt-[-3rem] flex justify-center rounded-md">
        <form onSubmit={handleSubmit} className="flex flex-col w-[35%] p-6 bg-gray-100 rounded-lg shadow-md fixed top-28">
            <div className="flex justify-between items-center">
                <h2 className="text-2xl font-bold">Register</h2>
                <h2 className="font-bold text-2xl cursor-pointer" onClick={() => modalDispatch({type: 'CLOSE_MODAL'})}>X</h2>
            </div>
            <div className="flex justify-between items-center mt-6 mb-2">
                <div>
                    <label htmlFor="email" className="block text-gray-700 font-semibold mb-2">Email</label>
                    <input name="email" onChange={handleOnChange} value={email} type="email" id="email" className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full" required />
                </div>
                <div>
                    <label htmlFor="password" className="block text-gray-700 font-semibold mb-2">Password</label>
                    <input name="password" onChange={handleOnChange} value={password} type="password" id="password" className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full" required
                    />
                </div>
            </div>
            <div className="flex justify-between items-center mb-2">
                <div>
                    <label htmlFor="fullname" className="block text-gray-700 font-semibold mb-2">Full Name</label>
                    <input name="fullname" onChange={handleOnChange} value={fullname} type="text" id="fullname" className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full" required
                    />
                </div>
                <div>
                    <label htmlFor="phone" className="block text-gray-700 font-semibold mb-2">Phone</label>
                    <input name="phone" onChange={handleOnChange} value={phone} type="number" id="phone" className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full" required
                    />
                </div>
            </div>
            <div className="mb-6">
                <label htmlFor="address" className="block text-gray-700 font-semibold mb-2">Address</label>
                <textarea  name="address" onChange={handleOnChange} value={address} type="text" id="address" className="border-gray-300 focus:border-indigo-500 focus:ring focus:ring-indigo-200 rounded-md p-2 w-full resize-none" required
                ></textarea>
            </div>
            <button className="bg-indigo-500 text-white rounded-md py-2 px-4 hover:bg-indigo-600 transition duration-200" type="submit">Register</button>
            <div className="flex justify-center mt-5">
                <h2>don't have an account ? <button onClick={() => modalDispatch({type: "LOGIN_MODAL"})}>Sign In</button></h2>
            </div>
        </form>
    </div>
  );
}

export default Register