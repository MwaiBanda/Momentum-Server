import {createBrowserRouter} from "react-router-dom";
import {Home} from "../pages/home.tsx";
import { Messages } from "../pages/messages.tsx";
import { Notifications } from "../pages/noticafications.tsx";
import { Payments } from "../pages/payments.tsx";
import { Users } from "../pages/users.tsx";
import { HOME, MESSAGES, NOFICATIONS, PAYMENTS, USERS } from "../util/constants.ts";

export const router = createBrowserRouter([
    {
        path: MESSAGES,
        element: <Messages/>,
    },
    {
        path: HOME,
        element: <Home/>,
    },
    {
        path: NOFICATIONS,
        element: <Notifications/>,
    },
    {
        path: PAYMENTS,
        element: <Payments/>,
    },
    {
        path: USERS,
        element: <Users/>,
    },
])