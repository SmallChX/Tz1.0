import React from "react";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import RequireAuth from "../utils/RequireAuth";
import RedirectIfLoggedIn from "../utils/RidirectIfLoggedin";

// unprotectedRoutes
import { auth_routes } from "./UnprotectedRoutes";
// protectedRoutes
// import { admin_routes } from "./AdminRoutes";
import { booth_routes } from "./BoothRoutes";
import { general_routes } from "./GeneralRoutes";
import { admin_routes } from "./AdminRoutes";
import NavbarDisplay from "../components/Nav";

const AppRoutes = () => {
    const protectedRoutes = [
        ...admin_routes,
        ...booth_routes,
        ...general_routes
    ];

    const unprotectedRoutes = [...auth_routes];

    return (
        <BrowserRouter>
            <NavbarDisplay></NavbarDisplay>
            <Routes>
                {
                    unprotectedRoutes.map((e) => {
                        return (
                            <Route
                                key={e.path}
                                exact
                                path={e.path}
                                element={
                                    <RedirectIfLoggedIn>
                                        {e.ele}
                                    </RedirectIfLoggedIn>
                                }
                                // element={e.ele}
                            />
                        );
                    })
                }

                {
                    protectedRoutes.map((e) => {
                        return (
                            <Route
                                key={e.path}
                                exact
                                path={e.path}
                                element={
                                    <RequireAuth userroles={e?.availability}>
                                        {e.ele}
                                    </RequireAuth>
                                }
                                // element={e.ele}
                            />
                        );
                    })
                }
            </Routes>
        </BrowserRouter>
    );
};
export default AppRoutes;