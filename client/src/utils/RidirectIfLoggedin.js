import { Navigate } from "react-router-dom";

const RedirectIfLoggedIn = ({children})=>{
    if(localStorage.getItem("userRole")){
        return <Navigate to="/home" />
    }
    return children;

}
export default RedirectIfLoggedIn;