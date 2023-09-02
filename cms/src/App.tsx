import {QueryClientProvider} from 'react-query';
import './App.css'
import {ReactQueryDevtools} from "react-query/devtools";
import {RouterProvider} from "react-router-dom";
import "./index.css";
import {router} from "./router/router.tsx";
import "./util/interceptors.ts"
import {queryClient} from "./util/constants.ts";

function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <RouterProvider router={router} />
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}

export default App
