
import { Nav } from "../../components/nav"
import { useQuery } from "react-query"
import axios from "axios"
import { TransactionTable } from "./table/transaction-table"

export const Payments = () => {
    const { data, isLoading } = useQuery({
        queryKey: ['transactions'],
        queryFn: async () => {
          const { data } = await axios.get('/api/v1/transactions')
          return data
        },
    })

    if (isLoading) return <div  className="w-full h-screen" >Loading...</div>

    return (
        <main className="w-full h-screen">
            <Nav onSigninClick={() => {}}/>
            <TransactionTable data={data} />
        </main>
    )
}


  







