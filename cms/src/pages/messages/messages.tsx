import { useQuery } from "react-query"
import { MomentumNavigation } from "../../components/navigation"
import { Card, TextInput } from 'flowbite-react';
import axios from "axios"
import { Message } from "../../models/message"
import { useEffect, useState } from "react";
import AddMessageModal from "./modals/add-message";
import { auth } from "../../util/firebase";
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

    const [showAuthModal, setShowAuthModal] = useState(false);
    const [searchText, setSearchText] = useState("");
    const [searchTags, setSearchTags] = useState([
        "preacher", "title", "series", "date"
    ]);
    const [openModal, setOpenMessageModal] = useState(false);
    const [openEditModal, setOpenEditMessageModal] = useState(false);
    const [editableMessage, setEditableMessage] = useState<Message | null>(null);
    
    useEffect(() => {
        setTimeout(() => {
            setSearchTags([searchTags[searchTags.length - 1], ...searchTags.slice(0, searchTags.length - 1)])
        }, 2000)
    }, [searchTags])

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
        <>
            <main className="w-full h-screen">
                <MomentumNavigation content={
                    <>
                    <div className="w-full flex justify-center items-end">
                    <TextInput
                         className="lg:w-[500px] w-[240px] ml-4 rounded-none" 
                         type="text" 
                         placeholder={`Search by "${searchTags[0]}"`}
                         onChange={(e) => {
                            setSearchText(e.target.value)
                         }}
                     />        
                    <button className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold py-2.5 px-4 rounded-none max-h-[42px] ml-[-5px] z-10 rounded-r-lg" onClick={() => {
                        if (auth.currentUser) {
                                setOpenMessageModal(true)
                        } else {
                                setShowAuthModal(true)
                        }
                        }}>
                            Add Message
                    </button>
                </div>
                <div className="grid grid-cols-2 gap-2 lg:grid-cols-5 lg:gap-4  md: md:grid-cols-3  w-full">
                    {data.data?.filter((m: Message) => {
                        return m.title.toLowerCase().includes(searchText.toLowerCase()) ||
                        m.preacher.toLowerCase().includes(searchText.toLowerCase()) ||
                        m.date.toLowerCase().includes(searchText.toLowerCase()) ||
                        m.series.toLowerCase().includes(searchText.toLowerCase()) ||
                        searchText.length === 0
                    })?.map((message: Message) => (
                        <div key={message.id}>
                            <MessageCard message={message} onEditMessage={
                                (message) => {
                                    if (auth.currentUser) {
                                        setEditableMessage(message)
                                        setOpenEditMessageModal(true)
                                    } else {
                                        setShowAuthModal(true)
                                    }
                                }
                            }/>
                        </div>
                    ))}
                </div>
                </>
                } showAuth={showAuthModal} />
                <AddMessageModal openModal={openModal} setOpenModal={setOpenMessageModal}/>
                <EditMessageModal openModal={openEditModal} setOpenModal={setOpenEditMessageModal} message={editableMessage}/>
            </main>
        </>
    )
}


function MessageCard({ message, onEditMessage }: { message: Message, onEditMessage: (message: Message) => void }) {
  return (
    <Card className="max-w-lg mt-6" renderImage={() => <img className="max-h-[140px] md:max-h-[160px] lg:max-h-[180px]" crossOrigin="anonymous"  src={message.thumbnail} title="source: imgur.com" />} horizontal={false}>
        <div className="min-w-96 flex justify-between items-top">
        <div>
            <h5 className="text-2xl font-bold tracking-tight text-gray-900 dark:text-white line-clamp-1">
                {message.preacher}
            </h5>
            <p className="font-normal text-gray-700 dark:text-gray-400 line-clamp-1">
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
        className="bg-momentum-orange hover:bg-momentum-orange text-white font-bold p-2 rounded-full max-h-11 ">
            <svg className="w-6 h-6 ml-1 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" >
                <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m14.304 4.844 2.852 2.852M7 7H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4.5m2.409-9.91a2.017 2.017 0 0 1 0 2.853l-6.844 6.844L8 14l.713-3.565 6.844-6.844a2.015 2.015 0 0 1 2.852 0Z"/>
            </svg>
        </button>
        </div>
        
    </Card>
  );
}



