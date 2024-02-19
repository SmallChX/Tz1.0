import Profile from "../pages/Profile/Profile";
import { userRoles as ur } from "../data/userRole";
import HomePage from "../pages/HomePage/HomePage"

export const general_routes = [
   
    {
        path: "/profile",
        ele: <Profile />,
        availability: [ur.company, ur.admin, ur.student],
    },
    {
        path: "/home",
        ele: <HomePage />,
        availability: [ur.company, ur.admin, ur.student],
    }, 
]
