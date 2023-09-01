import { QueryClient, QueryClientProvider } from 'react-query';
import './App.css'
import {Demo} from "./components/Demo.tsx";
import {ReactQueryDevtools} from "react-query/devtools";
import axios from 'axios';
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import "./index.css";
import {API_KEY} from "./util/Constants.ts";



const queryClient = new QueryClient()
axios.interceptors.request.use(
  config => {
    config.headers['Authorization'] = `Bearer ${API_KEY}`;
    return config;
  },
  error => {
    return Promise.reject(error);
  }
)

const router = createBrowserRouter([
  {
    path: "/dashboard/home",
    element: <Demo/>,
  },
  {
    path: "/dashboard",
    element: <Demo/>,
  },
  {
    path: "/dashboard/services",
    element: <h1>Services</h1>,
  },
]);

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}

export default App
