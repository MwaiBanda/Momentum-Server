import { TextArea } from "@/interfaces/presentation/components/textarea";
import { TextField } from "@/interfaces/presentation/components/textfield";
import { Message, Passage } from "@/domain/message/entity";
import { useApplication } from "@/interfaces/presentation/application-context";
import { Button, Modal, ModalBody, ModalFooter, ModalHeader, ToggleSwitch } from "flowbite-react";
import { useState } from "react";
import { useMutation } from "@tanstack/react-query";

export default function EditMessageModal({ message, openModal, setOpenModal }: { message: Message | null, openModal: boolean, setOpenModal: (open: boolean) => void }) {
    const application = useApplication();
    const [title, setTitle] = useState(() => message?.title ?? "");
    const [series, setSeries] = useState(() => message?.series ?? "");
    const [preacher, setPreacher] = useState(() => message?.preacher ?? "");
    const [image, setImage] = useState(() => message?.thumbnail ?? "");
    const [date, setDate] = useState(() => message?.date ?? "");
    const [published, setPublished] = useState(() => message?.published ?? false);
    const [passages, setPassages] = useState<Passage[]>(() => [...message?.passages ?? []]);
    const [passagesToUpdate, setPassagesToUpdate] = useState<Passage[]>([]);

    const mutation = useMutation({
        mutationFn: (data: Message) => application.message.update(data),
        onSuccess: () => {
            alert("Message posted successfully")
        }
    })

    const updatePassage = (index: number, updatedPassage: Passage) => {
        setPassages((currentPassages) => currentPassages.map((passage, currentIndex) => (
            currentIndex === index ? updatedPassage : passage
        )));
        setPassagesToUpdate((currentPassages) => {
            const matchesPassage = (passage: Passage) => updatedPassage.id !== undefined
                ? passage.id === updatedPassage.id
                : passage.order === updatedPassage.order;

            return currentPassages.some(matchesPassage)
                ? currentPassages.map((passage) => matchesPassage(passage) ? updatedPassage : passage)
                : [...currentPassages, updatedPassage];
        });
    };

    return (
        <Modal show={openModal} size="md" popup onClose={() =>{ 
            setPassages([...message?.passages ?? []])
            setOpenModal(false)
        }}>
            <ModalHeader />
            <ModalBody>
            <div className="space-y-6">
                    <h3 className="text-xl font-medium text-gray-900 dark:text-white">Edit Message</h3>
                    <div>
                    <ToggleSwitch color="green" checked={published} label="Published" onChange={setPublished} />

                    </div>
                    <TextField
                        value={title}
                        title="Title"
                        placeholder="Enter the title of the message"
                        id="title"
                        onChange={(event) => {
                            setTitle(event.target.value)
                        }}
                    />
                    <TextField
                        value={series}
                        title="Series"
                        placeholder="Enter the series of the message"
                        id="series"
                        onChange={(event) => {
                            setSeries(event.target.value)
                        }}
                    />
                    <TextField
                        value={preacher}
                        title="Preacher"
                        placeholder="Enter the name of the preacher"
                        id="preacher"
                        onChange={(event) => {
                            setPreacher(event.target.value)
                        }}
                    />
                    <TextField
                        value={date}
                        title="Date"
                        placeholder="Enter the display date of the message"
                        id="date"
                        onChange={(event) => {
                            setDate(event.target.value)
                        }}
                    />
                    <TextField
                        value={image}
                        title="Image"
                        placeholder="Enter the image cover url of the message"
                        id="image"
                        onChange={(event) => {
                            setImage(event.target.value)
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
                                        const temp: Passage = {
                                            id: passage.id,
                                            header: event.target.value,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        updatePassage(index, temp)
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
                                        const temp: Passage = {
                                            id: passage.id,
                                            verse: event.target.value,
                                            message: passage.message,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        updatePassage(index, temp)
                                    }}
                                />
                                <TextArea
                                    value={passage.message}
                                    title={`Message`}
                                    placeholder="Enter the message of the passage"
                                    id={`edit-message${index + 1}`}
                                    onChange={(event) => {
                                        const temp: Passage = {
                                            id: passage.id,
                                            verse: passage.verse,
                                            message: event.target.value,
                                            type: passage.type,
                                            order: passage.order
                                        }
                                        updatePassage(index, temp)
                                    }}
                                />
                            </>
                        }
                    })}
                </div>
            </ModalBody>
            <ModalFooter>
                <Button className="bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                    mutation.mutate({
                        id: message?.id ?? "",
                        title: title,
                        series: series,
                        preacher: preacher,
                        thumbnail: image,
                        date: date,
                        passages: passagesToUpdate,
                        published: published
                    })
                    setOpenModal(false)
                }}>Update</Button>
            </ModalFooter>
        </Modal>
    )
}