import axios from "axios";
import { Modal, Button, Label, TextInput, Textarea } from "flowbite-react"
import { Ref, RefObject, createRef, useRef, useState } from "react"
import { useMutation } from "react-query";
import { MessageRequest } from "../models/message";
enum PassageType {
    Header,
    Message
}

type PassageHolder = {
    type: PassageType;
    text: Ref<HTMLInputElement>;
    area: Ref<HTMLTextAreaElement>
}


export default function AddMessageModal({ openModal, setOpenModal }: { openModal: boolean, setOpenModal: (open: boolean) => void }) {
    const [refs, setRefs] = useState<PassageHolder[]>([]);
    const titleInputRef = useRef<HTMLInputElement>(null);
    const seriesInputRef = useRef<HTMLInputElement>(null);
    const preacherInputRef = useRef<HTMLInputElement>(null);
    const imageInputRef = useRef<HTMLInputElement>(null);
    const dateInputRef = useRef<HTMLInputElement>(null);

    const mutation = useMutation((data: MessageRequest) => {
        return axios.post('/api/v1/messages', data)
    }, {
        onSuccess: () => {
            alert("Message posted successfully")
        }
    })

    return (
        <Modal show={openModal} size="md" popup onClose={() => setOpenModal(false)} initialFocus={titleInputRef}>
            <Modal.Header />
            <Modal.Body>
                <div className="space-y-6">
                    <h3 className="text-xl font-medium text-gray-900 dark:text-white">Add a Message</h3>
                    <TextField
                        reference={titleInputRef}
                        title="Title"
                        placeholder="Enter the title of the message"
                        id="title"
                    />
                    <TextField
                        reference={seriesInputRef}
                        title="Series"
                        placeholder="Enter the series of the message"
                        id="series"
                    />
                    <TextField
                        reference={preacherInputRef}
                        title="Preacher"
                        placeholder="Enter the name of the preacher"
                        id="preacher"
                    />
                    <TextField
                        reference={dateInputRef}
                        title="Date"
                        placeholder="Enter the display date of the message"
                        id="date"
                    />
                    <TextField
                        reference={imageInputRef}
                        title="Image"
                        placeholder="Enter the image cover url of the message"
                        id="image"
                    />
                    {refs.map((passage, index) => {
                        {
                            switch (passage.type) {
                                case PassageType.Header:
                                    return (
                                        <TextArea
                                            reference={passage.area!!}
                                            title={`Add header`}
                                            placeholder="Enter the header of the passage"
                                            id={`header${index + 1}`}
                                        />
                                    )
                                case PassageType.Message:
                                    return (
                                        <div>
                                            <TextField
                                                reference={passage.text!!}
                                                title={`Add verse`}
                                                placeholder="Enter the verse of the passage"
                                                id={`verse${index + 1}`}
                                            />
                                            <TextArea
                                                reference={passage.area!!}
                                                title={`Add message`}
                                                placeholder="Enter the message of the passage"
                                                id={`message${index + 1}`}
                                            />
                                        </div>

                                    )
                            }
                        }

                    }
                    )}
                    <div className="w-full flex justify-between">
                        <Button onClick={() => {
                            mutation.mutate({
                                title: titleInputRef.current?.value ?? "",
                                series: seriesInputRef.current?.value ?? "",
                                preacher: preacherInputRef.current?.value ?? "",
                                thumbnail: imageInputRef.current?.value ?? "",
                                date: dateInputRef.current?.value ?? "",
                                passages: refs.map((passage, index) => {
                                    switch (passage.type) {
                                        case PassageType.Header:
                                            return {
                                                header: (passage.area as RefObject<HTMLTextAreaElement>).current?.value ?? "",
                                                verse: null,
                                                message: null,
                                                type: "header",
                                                order: index
                                            }
                                        case PassageType.Message:
                                            return {
                                                header: null,
                                                verse: (passage.text as RefObject<HTMLInputElement>).current?.value ?? "",
                                                message: (passage.area as RefObject<HTMLTextAreaElement>).current?.value ?? "",
                                                type: "message",
                                                order: index
                                            }
                                    }
                                })
                            })
                            setOpenModal(false)
                        }}>Submit</Button>
                        <div className="flex">
                            <Button onClick={() => {
                                setRefs([...refs, { type: PassageType.Header, text: null, area: createRef() }])
                                console.log(refs)
                            }}>Add header</Button>
                            <Button className="ml-1" onClick={() => {
                                setRefs([...refs, { type: PassageType.Message, text: createRef(), area: createRef() }])
                                console.log(refs)
                            }}>Add message</Button>
                        </div>

                    </div>

                </div>
            </Modal.Body>
        </Modal>
    )
}

function TextField({ reference, title, placeholder, id }: { reference: Ref<HTMLInputElement>, title: string, placeholder: string, id: string }) {
    return (
        <div>
            <div className="mb-2 block">
                <Label htmlFor={id} value={title} />
            </div>
            <TextInput id={id} ref={reference} placeholder={placeholder} required />
        </div>
    )
}

function TextArea({ reference, title, placeholder, id }: { reference: Ref<HTMLTextAreaElement>, title: string, placeholder: string, id: string }) {
    return (
        <div>
            <div className="mb-2 block">
                <Label htmlFor={id} value={title} />
            </div>
            <Textarea id={id} ref={reference} placeholder={placeholder} required />
        </div>
    )
}
