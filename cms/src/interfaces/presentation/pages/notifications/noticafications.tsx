import { useMutation } from "@tanstack/react-query";
import { MomentumNavigation } from "../../components/navigation"
import { Button, Card, Label, TextInput, Textarea } from 'flowbite-react';
import { useApplication } from "@/interfaces/presentation/application-context";
import { useState } from "react";
import { Radio } from 'flowbite-react';

import type { NotificationRequest } from "@/domain/notification/entity";
import { schemas } from "@/infrastructure/http/generated/schemas";

export const Notifications = () => {
    const [openAuthModal, setOpenAuthModal] = useState(false);

    return (
        <main className="w-full h-screen">
            <MomentumNavigation content={
                <div className="w-full h-screen flex justify-center items-center">
                <NotificationCard setOpenAuthModal={() => { setOpenAuthModal(true) }}/>
                </div>
            } showAuth={openAuthModal} onShowAuthChange={setOpenAuthModal}/>
        </main>
    )
}

function NotificationCard({ setOpenAuthModal }: { setOpenAuthModal: () => void}){
    const application = useApplication();
    const [title, setTitle] = useState("")
    const [body, setBody] = useState("")
    const [topic, setTopic] = useState("MomentumUsers")

    const mutation = useMutation({
        mutationFn: (data: NotificationRequest) => application.notification.send(data),
        onSuccess: () => {
            alert("Notification sent successfully")
        },
        onError: () => {
            alert("Notification could not be sent. Please try again.")
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
                        <Label htmlFor="title">Title</Label>
                    </div>
                    <TextInput id="title" type="text" required onChange={(e) => {
                        setTitle(e.target.value)
                    }} />
                </div>
                <div>
                    <div className="mb-2 block">
                        <Label htmlFor="body">Body</Label>
                    </div>
                    <Textarea id="body" rows={5} required onChange={(e) => {
                        setBody(e.target.value)
                    }} />
                </div>

                <div>
                    <fieldset className="flex max-w-md flex-col gap-4">
                        <legend className="mb-4">Choose your target audience</legend>
                        <div className="flex items-center gap-2">
                            <Radio id="all-users" name="audience" value="MomentumUsers" defaultChecked onChange={(e) => setTopic(e.target.value)} />
                            <Label htmlFor="all-users">All Users</Label>
                        </div>
                        <div className="flex items-center gap-2">
                            <Radio id="android" name="audience" value="MomentumAndroid" onChange={(e) => setTopic(e.target.value)} />
                            <Label htmlFor="android">Android</Label>
                        </div>
                        <div className="flex items-center gap-2">
                            <Radio id="ios" name="audience" value="MomentumIOS" onChange={(e) => setTopic(e.target.value)} />
                            <Label htmlFor="ios">iOS</Label>
                        </div>
                    </fieldset>
                </div>

                <Button className="bg-momentum-orange hover:bg-momentum-orange  enabled:hover:bg-momentum-orange" type="submit" onClick={
                    (e) => {
                        e.preventDefault();
                        if (application.auth.currentUser()) {
                             const result = schemas.notification_Input.safeParse({ title, body, topic });
                             if (!result.success) {
                                 alert(result.error.issues[0]?.message ?? "Please check the notification details.")
                                 return
                             }
                             mutation.mutate(result.data)
                       } else {
                            setOpenAuthModal()
                       }
                        
                    }
                }>Send Notification</Button>
            </form>
        </Card>
    );
}
