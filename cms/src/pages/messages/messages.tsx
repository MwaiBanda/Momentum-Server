import { useQuery } from "react-query"
import { Nav } from "../../components/nav"
import { Card } from 'flowbite-react';
import axios from "axios"
import { Message } from "../../models/message"
import { useState } from "react";
import AddMessageModal from "./modals/add-message";
import { auth } from "../../util/firebase";
import SigninModal from "../../components/signinmodal";
import EditMessageModal from "./modals/edit-message";
import { Rings } from 'react-loader-spinner'

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
    const [openEditModal, setOpenEditMessageModal] = useState(false);
    const [editableMessage, setEditableMessage] = useState<Message | null>(null);

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
        <Nav onSigninClick={() => setOpenAuthModal(true)}/>
        <div className="w-full flex flex-col justify-normal items-end">
        <button className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10" onClick={() => {
           if (auth.currentUser) {
                setOpenMessageModal(true)
           } else {
                setOpenAuthModal(true)
           }
        }}>
            Add Message
        </button>
        </div>
        <div className="lg:grid lg:grid-cols-4 lg:gap-4  md:grid md:grid-cols-2 md:gap-2  w-full h-screen flex flex-col justify-normal items-center">
            {data.data.map((message: Message) => (
                <div key={message.id}>
                    <MessageCard message={message} onEditMessage={
                        (message) => {
                            if (auth.currentUser) {
                                setEditableMessage(message)
                                setOpenEditMessageModal(true)
                            } else {
                                setOpenAuthModal(true)
                            }
                        }
                    }/>
                </div>
            ))}
        </div>
        <SigninModal openModal={openAuthModal} setOpenModal={setOpenAuthModal}/>
        <AddMessageModal openModal={openModal} setOpenModal={setOpenMessageModal}/>
        <EditMessageModal openModal={openEditModal} setOpenModal={setOpenEditMessageModal} message={editableMessage}/>
        </main>
    )
}


function MessageCard({ message, onEditMessage }: { message: Message, onEditMessage: (message: Message) => void }) {
  return (
    <Card className="max-w-lg mt-6" renderImage={() => <img className="max-h-[210px]" crossOrigin="anonymous"  src={message.thumbnail} title="source: imgur.com" />} horizontal={false}>
        <div className="min-w-96 flex justify-between items-center">
        <div>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                {message.preacher}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400">
                {message.date}
            </p>
            <p className="font-normal text-gray-700 dark:text-gray-400 line-clamp-1">
                {message.series}
            </p>
            <p className="font-normal text-gray-700 dark:text-gray-400 line-clamp-1">
                {message.title}
            </p>
        </div>
        <button 
        onClick={() => {
            onEditMessage(message)
        }}
        className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2 px-4 rounded max-h-10">
            Modify
        </button>
        </div>
        
    </Card>
  );
}



