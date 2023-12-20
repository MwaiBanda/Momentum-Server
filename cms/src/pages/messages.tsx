import { useQuery } from "react-query"
import { Nav } from "../components/nav"
import { Card } from 'flowbite-react';
import axios from "axios"
import { Message } from "../models/message"
import { useState } from "react";
import AddMessageModal from "../components/addmessagemodal";
import { auth } from "../util/firebase";
import SignupModal from "../components/signupmodal";

function MessageCard({ message }: { message: Message }) {
  return (
    <Card className="max-w-lg mt-6" renderImage={() => <img crossOrigin="anonymous" src={message.thumbnail} title="source: imgur.com" />} horizontal={false}>
        <div className="min-w-96 flex justify-between items-center">
        <div>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                {message.preacher}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
                {message.date}
            </p>
            <p className="font-normal text-gray-700 dark:text-gray-400">
                {message.series}
            </p>
            <p className="font-normal text-gray-700 dark:text-gray-400">
                {message.title}
            </p>
        </div>
        <button className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10">
            Modify
        </button>
        </div>
        
    </Card>
  );
}



export const Messages = () => {
    const { data, isLoading } = useQuery({
        queryKey: ['message'],
        queryFn: async () => {
          const { data } = await axios.get('/api/v1/messages/admin')
          return data
        },
    })

    const [openAuthModal, setOpenAuthModal] = useState(false);
    const [openModal, setOpenMessageModal] = useState(false);
 
    if (isLoading) return <div>Loading...</div>

    return (
        <main className="w-full h-screen">
        <Nav/>
        <div className="w-full flex flex-col justify-normal items-end">
        <button className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10" onClick={() => {
           if (auth.currentUser?.uid === undefined) {
              setOpenAuthModal(true)
           } else {
              setOpenMessageModal(true)
           }
        }}>
            Add Message
        </button>
        </div>
        <div className="w-full h-screen flex flex-col justify-normal items-center">
            {data.data.map((message: Message) => (
                <div key={message.id}>
                    <MessageCard message={message}/>
                </div>
            ))}
        </div>
        <SignupModal openModal={openAuthModal} setOpenModal={setOpenAuthModal}/>
        <AddMessageModal openModal={openModal} setOpenModal={setOpenMessageModal}/>
        </main>
    )
}


