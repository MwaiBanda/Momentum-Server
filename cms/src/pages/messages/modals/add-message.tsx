import axios from "axios";
import { Modal, Button } from "flowbite-react"
import { Ref, RefObject, createRef, useRef, useState } from "react"
import { useMutation } from "react-query";
import { MessageRequest } from "../../../models/message";
import { TextField } from "@/components/textfield";
import { TextArea } from "@/components/textarea";

enum PassageType {
    Header,
    Message
}

type PassageHolder = {
    type: PassageType;
    text: React.RefObject<HTMLInputElement>;
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
                    <h3 className="text-xl font-medium text-gray-900 dark:text-white">Add Message</h3>
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
                        <Button className="bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
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
                                                type: "header",
                                                order: index
                                            }
                                        case PassageType.Message:
                                            return {
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
                            <Button className="bg-momentum-orange hover:bg-momentum-orange active:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                                setRefs([...refs, { type: PassageType.Header, text: createRef(), area: createRef() }])
                                console.log(refs)
                            }}>Header</Button>
                            <div className="flex">
                            <Button className="ml-1 bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                                setRefs([...refs, { type: PassageType.Message, text: createRef(), area: createRef() }])
                                console.log(refs)
                            }}>Message</Button>
                           {refs.length > 0 ? <Button className="ml-1 rounded-full bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange" onClick={() => {
                                setRefs([...refs.slice(0, refs.length - 1)])
                                console.log(refs)
                            }}><svg className="w-4 h-4 rounded-full text-white-800 dark:text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h13a5 5 0 0 1 0 10H7M3 9l4-4M3 9l4 4"/>
                          </svg></Button> : <></>}
                            </div>
                            
                        </div>

                    </div>

                </div>
            </Modal.Body>
        </Modal>
    )
}




