import Booth from "../pages/Booth/Booth";
import { userRoles as ur } from "../data/userRole";

export const booth_routes = [
    {
        path: "booth",
        ele: <Booth />,
        availability: [ur.admin, ur.company],
    }
]