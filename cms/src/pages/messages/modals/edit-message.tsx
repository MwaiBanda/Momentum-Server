import { TextArea } from "@/components/textarea";
import { TextField } from "@/components/textfield";
import { Message, Passage } from "@/models/message";
import { Button, Modal, ToggleSwitch } from "flowbite-react";
import { useEffect, useState } from "react";

export default function EditMessageModal({ message, openModal, setOpenModal }: { message: Message | null, openModal: boolean, setOpenModal: (open: boolean) => void }) {
    const [title, setTitle] = useState("");
    const [series, setSeries] = useState("");
    const [preacher, setPreacher] = useState("");
    const [image, setImage] = useState("");
    const [date, setDate] = useState("");
    const [published, setPublished] = useState(false);
    const [canUpdateMessage, setCanUpdateMessage] = useState(false);
    const [passages, setPassages] = useState<Passage[]>([]);
    const [passagesToUpdate, setPassagesToUpdate] = useState<Passage[]>([]);

    useEffect(() => {
        console.log(message)
        if (message) {
            setTitle(message.title)
            setPublished(message.published)
            setSeries(message.series)
            setPreacher(message.preacher)
            setImage(message.thumbnail)
            setDate(message.date)
            setPassages([...message?.passages ?? []])
        }
    } ,[message])

    useEffect(() => {
        console.log(passagesToUpdate)
    }, [passagesToUpdate])

    return (
        <Modal show={openModal} size="md" popup onClose={() =>{ 
            setPassages([...message?.passages ?? []])
            setOpenModal(false)
        }}>
            <Modal.Header />
            <Modal.Body>
            <div className="space-y-6">
                    <h3 className="text-xl font-medium text-gray-900 dark:text-white">Edit Message</h3>
                    <div>
                    <ToggleSwitch checked={published} label="Published" onChange={setPublished} />

                    </div>
                    <TextField
                        value={title}
                        title="Title"
                        placeholder="Enter the title of the message"
                        id="title"
                        onChange={(event) => {
                            setTitle(event.target.value)
                            setCanUpdateMessage(true)
                        }}
                    />
                    <TextField
                        value={series}
                        title="Series"
                        placeholder="Enter the series of the message"
                        id="series"
                        onChange={(event) => {
                            setSeries(event.target.value)
                            setCanUpdateMessage(true)
                        }}
                    />
                    <TextField
                        value={preacher}
                        title="Preacher"
                        placeholder="Enter the name of the preacher"
                        id="preacher"
                        onChange={(event) => {
                            setPreacher(event.target.value)
                            setCanUpdateMessage(true)
                        }}
                    />
                    <TextField
                        value={date}
                        title="Date"
                        placeholder="Enter the display date of the message"
                        id="date"
                        onChange={(event) => {
                            setDate(event.target.value)
                            setCanUpdateMessage(true)
                        }}
                    />
                    <TextField
                        value={image}
                        title="Image"
                        placeholder="Enter the image cover url of the message"
                        id="image"
                        onChange={(event) => {
                            setImage(event.target.value)
                            setCanUpdateMessage(true)
                        }}
                    />
                    {passages.map((passage, index) => {
                        if ((passage?.header ?? "").length > 0){
                            return <>
                                <TextArea
                                    value={passage.header}
                                    title={`Header`}
                                    placeholder="Enter the header of the passage"
                                    id={`edit-header${index + 1}`}
                                    onChange={(event) => {
                                        let temp: Passage = {
                                            id: passage.id,
                                            header: event.target.value,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        passages[index] = temp
                                        if (passagesToUpdate.find(passage => passage.id === temp.id) === undefined) {
                                            setPassagesToUpdate([...passagesToUpdate, temp])
                                        } else {    
                                            setPassagesToUpdate([...passagesToUpdate.filter(passage => passage.id !== temp.id), temp])
                                        }
                                        setPassages([...passages])
                                    }}
                                />
                            </>
                        } else {
                            return <>
                                <TextField
                                    value={passage.verse}
                                    title={`Verse`}
                                    placeholder="Enter the verse of the passage"
                                    id={`edit-verse${index + 1}`}
                                    onChange={(event) => {
                                        let temp: Passage = {
                                            id: passage.id,
                                            verse: event.target.value,
                                            message: passage.message,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        passages[index] = temp
                                        if (passagesToUpdate.filter(passage => passage.id === temp.id).length === 0) {
                                            setPassagesToUpdate([...passagesToUpdate, temp])
                                        } else {    
                                            setPassagesToUpdate([...passagesToUpdate.filter(passage => passage.id !== temp.id), temp])
                                        }
                                        setPassages([...passages])
                                    }}
                                />
                                <TextArea
                                    value={passage.message}
                                    title={`Message`}
                                    placeholder="Enter the message of the passage"
                                    id={`edit-message${index + 1}`}
                                    onChange={(event) => {
                                        let temp: Passage = {
                                            id: passage.id,
                                            verse: passage.verse,
                                            message: event.target.value,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        passages[index] = temp
                                        if (passagesToUpdate.filter(passage => passage.id === temp.id).length === 0) {
                                            setPassagesToUpdate([...passagesToUpdate, temp])
                                        } else {    
                                            setPassagesToUpdate([...passagesToUpdate.filter(passage => passage.id !== temp.id), temp])
                                        }
                                        setPassages([...passages])
                                    }}
                                />
                            </>
                        }
                    })}
                </div>
            </Modal.Body>
            <Modal.Footer>
                <Button onClick={() => {
                    if (canUpdateMessage) {
                        
                    }   
                    if (passagesToUpdate.length > 0) {
                        
                    }
                    setOpenModal(false)
                }}>Update</Button>
            </Modal.Footer>
        </Modal>
    )
}