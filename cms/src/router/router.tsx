import {createBrowserRouter} from "react-router-dom";
import {Demo} from "../components/demo.tsx";

export const router = createBrowserRouter([
    {
        path: "/dashboard/home",
        element: <Demo/>,
    },
    {
        path: "/dashboard",
        element: <Demo/>,
    },
    {
        path: "/services",
        element: <h1>Services</h1>,
    },
])