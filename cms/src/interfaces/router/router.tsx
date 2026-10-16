import {createBrowserRouter} from "react-router-dom";
import {Home} from "../presentation/pages/home.tsx";
import { Messages } from "../presentation/pages/messages/messages.tsx";
import { Notifications } from "../presentation/pages/notifications/noticafications.tsx";
import { Payments } from "../presentation/pages/payments/payments.tsx";
import { Users } from "../presentation/pages/users/users.tsx";
import { HOME, MESSAGES, NOFICATIONS, PAYMENTS, USERS } from "@/application/config/constants";

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