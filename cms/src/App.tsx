import { QueryClient, QueryClientProvider } from 'react-query';
import './App.css'
import {Demo} from "./components/Demo.tsx";
import {ReactQueryDevtools} from "react-query/devtools";
import axios from 'axios';
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
function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Demo/>
      <ReactQueryDevtools initialIsOpen={false} />
    </QueryClientProvider>
  )
}

export default App
