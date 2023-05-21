import { createContext, useReducer } from "react";

export const ModalContext = createContext();

const initialState = {
    isModalLogin: false,
    isModalRegist: false,
};

const reducer = (state, action) => {
    const {type, _} =action;

    switch (type) {
        case 'LOGIN_MODAL' :
            return {
                isModalLogin: true,
                isModalRegist: false,
            };
        case 'REGISTER_MODAL' :
            return {
                isModalRegist: true,
                isModalLogin: false,
            };
        case 'CLOSE_MODAL' :
            return {
                isModalLogin: false,
                isModalRegist: false,
            };
        default:
            throw new Error();
    }
}

export const ModalContextProvider = ({children}) => {
    const [state, dispatch] = useReducer(reducer, initialState);

    return (
        <ModalContext.Provider value={[state, dispatch]}>
            {children}
        </ModalContext.Provider>
    ) 
}

