import { userRoles as ur } from "../data/userRole";
import RequestManagement from "../pages/Admin/RequestManagement/RequestManagement";
import AccountManager from "../pages/Admin/AccountManagement";
import BoothManagement from "../pages/Admin/BoothManagement/BoothManagement"

export const admin_routes = [
    {
        path: "admin",
        ele: <RequestManagement />,
        availability: [ur.admin, ur.company],
    }, 
    {
        path: "account",
        ele: <AccountManager />,
        availability: [ur.admin],
    },
    {
        path: "admin/booth",
        ele: <BoothManagement />,
        availability: [ur.admin, ur.company],
    }
]