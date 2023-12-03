import { useQuery } from "react-query"
import { Nav } from "../components/nav"
import { Card } from 'flowbite-react';
import axios from "axios"
import { Message } from "../models/message"

function MessageCard({ message }: { message: Message }) {
  return (
    <Card className="max-w-lg mt-6" imgSrc="assets/logo.png" horizontal>
        <div className="w-96">
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

    if (isLoading) return <div>Loading...</div>

    return (
        <main className="w-full h-screen">
        <Nav/>
        <div className="w-full h-screen flex flex-col justify-normal items-center">
            {data.data.map((message: Message) => (
                <div key={message.id}>
                    <MessageCard message={message}/>
                </div>
            ))}
        </div>
        </main>
    )
}




