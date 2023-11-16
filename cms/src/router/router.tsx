import {createBrowserRouter} from "react-router-dom";
import {Home} from "../pages/home.tsx";
import { Messages } from "../pages/messages.tsx";
import { Notifications } from "../pages/noticafications.tsx";
import { Payments } from "../pages/payments.tsx";
import { Users } from "../pages/users.tsx";

export const router = createBrowserRouter([
    {
        path: "/messages",
        element: <Messages/>,
    },
    {
        path: "/dashboard",
        element: <Home/>,
    },
    {
        path: "/notifications",
        element: <Notifications/>,
    },
    {
        path: "/payments",
        element: <Payments/>,
    },
    {
        path: "/users",
        element: <Users/>,
    },
])