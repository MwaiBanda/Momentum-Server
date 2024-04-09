
import { Nav } from "../../components/nav"
import { useQuery } from "react-query"
import axios from "axios"
import { TransactionTable } from "./table/transaction-table"
import { useState } from "react"
import SigninModal from "@/components/signinmodal"
import { Rings } from "react-loader-spinner"

export const Payments = () => {
    const { data, isLoading } = useQuery({
        queryKey: ['transactions'],
        queryFn: async () => {
          const { data } = await axios.get('/api/v1/transactions')
          return data
        },
    })
    const [openAuthModal, setOpenAuthModal] = useState(false);

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
            <Nav onSigninClick={() => {setOpenAuthModal(true)}}/>
            <TransactionTable data={data} />
            <SigninModal openModal={openAuthModal} setOpenModal={setOpenAuthModal}/>
        </main>
    )
}


  







