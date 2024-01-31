import { useMutation } from "react-query";
import { Nav } from "../../components/nav"
import { Button, Card, Label, TextInput, Textarea } from 'flowbite-react';
import axios from "axios";
import { useState } from "react";
import { auth } from "../../util/firebase";

interface Notification {
    title: string;
    body: string;
    topic: string;
};

function NotificationCard({ setOpenAuthModal }: { setOpenAuthModal: () => void}){
    const [title, setTitle] = useState("")
    const [body, setBody] = useState("")
    const [topic, setTopic] = useState("MomentumUsers")

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
                    <TextInput id="title" type="text" required onChange={(e) => {
                        setTitle(e.target.value)
                    }} />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label htmlFor="body" value="Body" />
                    </div>
                    <Textarea id="body" rows={5} required onChange={(e) => {
                        setBody(e.target.value)
                    }} />
                </div>

                <div>
                    <fieldset className="flex max-w-md flex-col gap-4" onChange={(e) => {setTopic((e.target as HTMLInputElement).value)}}>
                        <legend className="mb-4">Choose your target audience</legend>
                        <div className="flex items-center gap-2">
                            <Radio id="all-users" name="audience" value="MomentumUsers" defaultChecked />
                            <Label htmlFor="all-users">All Users</Label>
                        </div>
                        <div className="flex items-center gap-2">
                            <Radio id="android" name="audience" value="MomentumAndroid" />
                            <Label htmlFor="android">Android</Label>
                        </div>
                        <div className="flex items-center gap-2">
                            <Radio id="ios" name="audience" value="MomentumIOS" />
                            <Label htmlFor="ios">iOS</Label>
                        </div>
                    </fieldset>
                </div>

                <Button className="bg-momentum-orange hover:bg-momentum-orange  enabled:hover:bg-momentum-orange" type="submit" onClick={
                    (e) => {
                        e.preventDefault();
                        if (auth.currentUser) {
                            if (title === "" || body === "") {
                                alert("Please fill in all fields")
                                return
                            }
                            mutation.mutate({
                                title: title,
                                body: body,
                                topic: topic
                            })
                       } else {
                            setOpenAuthModal()
                       }
                        
                    }
                }>Send Notification</Button>
            </form>
        </Card>
    );
}

export const Notifications = () => {
    const [openAuthModal, setOpenAuthModal] = useState(false);

    return (
        <main className="w-full h-screen">
            <Nav onSigninClick={() => { setOpenAuthModal(true) }}/>
            <div className="w-full h-screen flex justify-center items-center">
                <NotificationCard setOpenAuthModal={() => { setOpenAuthModal(true) }}/>
            </div>
            <SigninModal openModal={openAuthModal} setOpenModal={setOpenAuthModal}/>
        </main>
    )
}
import { Radio } from 'flowbite-react';
import SigninModal from "@/components/signinmodal";

