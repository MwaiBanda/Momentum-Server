
import { MomentumNavigation } from "../../components/navigation"
import { useQuery } from "@tanstack/react-query"
import { useApplication } from "@/interfaces/presentation/application-context";
import { TransactionTable } from "./table/transaction-table"
import { Rings } from "react-loader-spinner"

export const Payments = () => {
    const application = useApplication();
    const { data, isLoading } = useQuery({
        queryKey: ['transactions'],
        queryFn: async () => {
          return application.transaction.list()
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
                <TransactionTable data={data ?? []} />
            }/>
        </main>
    )
}


  



