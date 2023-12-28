import { useMutation } from "react-query";
import { Nav } from "../../components/nav"
import { Button, Card, Label, TextInput, Textarea } from 'flowbite-react';
import axios from "axios";
import { useState } from "react";

interface Notification {
    title: string;
    body: string;
    topic: string;
};

function NotificationCard() {
    const [title, setTitle] = useState("")
    const [body, setBody] = useState("")
    const [topic, setTopic] = useState("")

    const mutation = useMutation((data: Notification) => {
        return axios.post('/api/v1/notifications', data)
    }, {
        onSuccess: () => {
            alert("Notification sent successfully")
        }
    })

    return (
        <Card className="min-w-[23%]">
            <form className="flex flex-col gap-4">
                <h5 className="mb-2 text-3xl font-bold text-gray-900 dark:text-white">Reach all users anytime</h5>
                <p className="mb-5 text-base text-gray-500 dark:text-gray-400 sm:text-lg">
                    Notify users on iOS & Android on upcoming events and lastest updates.
                </p>
                <div>
                    <div className="mb-2 block">
                        <Label htmlFor="title" value="Title" />
                    </div>
                    <TextInput id="title" type="text" onChange={(e) => {
                        setTitle(e.target.value)
                    }} />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label htmlFor="body" value="Body" />
                    </div>
                    <Textarea id="body" rows={5} onChange={(e) => {
                        setBody(e.target.value)
                    }} />
                </div>

                <div>
                    <div className="mb-2 block">
                        <Label htmlFor="topic" value="Topic" />
                    </div>
                    <TextInput id="topic" type="text" onChange={(e) => {
                        setTopic(e.target.value)
                    }} />
                </div>

                <Button className="bg-momentum-orange hover:bg-momentum-orange  enabled:hover:bg-momentum-orange" type="submit" onClick={
                    (e) => {
                        e.preventDefault();
                        mutation.mutate({
                            title: title,
                            body: body,
                            topic: topic
                        })
                    }
                }>Send Notification</Button>
            </form>
        </Card>
    );
}

export const Notifications = () => {
    return (
        <main className="w-full h-screen">
            <Nav onSigninClick={() => {
                
            }}/>
            <div className="w-full h-screen flex justify-center items-center">
                <NotificationCard />
            </div>
        </main>
    )
}