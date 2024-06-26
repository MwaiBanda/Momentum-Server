
import { MomentumNavigation } from "../../components/navigation"
import { useQuery } from "react-query"
import axios from "axios"
import { TransactionTable } from "./table/transaction-table"
import { Rings } from "react-loader-spinner"

export const Payments = () => {
    const { data, isLoading } = useQuery({
        queryKey: ['transactions'],
        queryFn: async () => {
          const { data } = await axios.get('/api/v1/transactions')
          return data
        },
    })

    if (isLoading) return <div className="w-full h-screen flex flex-col justify-center items-center" >
        <Rings
            visible={true}
            height="150"
            width="150"
            color="#ec750c"
            ariaLabel="rings-loading"
            wrapperStyle={{}}
            wrapperClass=""
        />
    </div>
    return (
        <main className="w-full h-screen">
            <MomentumNavigation content={
                <TransactionTable data={data} />
            }/>
        </main>
    )
}


  







