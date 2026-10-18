import { TextArea } from "@/interfaces/presentation/components/textarea";
import { TextField } from "@/interfaces/presentation/components/textfield";
import type { Message, Passage } from "@/domain/message/entity";
import { useApplication } from "@/interfaces/presentation/application-context";
import { Alert, Button, Modal, ModalBody, ModalFooter, ModalHeader, ToggleSwitch } from "flowbite-react";
import { Fragment, useState } from "react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { schemas } from "@/infrastructure/http/generated/schemas";

export default function EditMessageModal({ message, openModal, setOpenModal }: { message: Message | null, openModal: boolean, setOpenModal: (open: boolean) => void }) {
    const application = useApplication();
    const queryClient = useQueryClient();
    const [title, setTitle] = useState(() => message?.title ?? "");
    const [series, setSeries] = useState(() => message?.series ?? "");
    const [preacher, setPreacher] = useState(() => message?.preacher ?? "");
    const [image, setImage] = useState(() => message?.thumbnail ?? "");
    const [date, setDate] = useState(() => message?.date ?? "");
    const [published, setPublished] = useState(() => message?.published ?? false);
    const [passages, setPassages] = useState<Passage[]>(() => [...message?.passages ?? []]);
    const [passagesToUpdate, setPassagesToUpdate] = useState<Passage[]>([]);
    const [submitMessage, setSubmitMessage] = useState<{ type: "success" | "error"; text: string } | null>(null);

    const mutation = useMutation({
        mutationFn: (data: Message) => application.message.update(data),
        onSuccess: async () => {
            await Promise.all([
                queryClient.invalidateQueries({ queryKey: ["message-admin"] }),
                queryClient.invalidateQueries({ queryKey: ["message-unpublished"] }),
            ]);
            setSubmitMessage({ type: "success", text: "Message updated successfully." });
        },
        onError: () => {
            setSubmitMessage({ type: "error", text: "Message could not be updated. Please try again." });
        },
    });

    const resetForm = () => {
        setTitle(message?.title ?? "");
        setSeries(message?.series ?? "");
        setPreacher(message?.preacher ?? "");
        setImage(message?.thumbnail ?? "");
        setDate(message?.date ?? "");
        setPublished(message?.published ?? false);
        setPassages([...message?.passages ?? []]);
        setPassagesToUpdate([]);
        setSubmitMessage(null);
    };

    const closeModal = () => {
        if (mutation.isPending) {
            return;
        }
        resetForm();
        setOpenModal(false);
    };

    const submitUpdate = () => {
        if (!message?.id) {
            setSubmitMessage({ type: "error", text: "This message cannot be updated because it has no ID." });
            return;
        }

        const result = schemas.message_UpdateInput.safeParse({
            id: message.id,
            title,
            series,
            preacher,
            thumbnail: image,
            date,
            passages: passagesToUpdate,
            published,
        });

        if (!result.success) {
            setSubmitMessage({ type: "error", text: result.error.issues[0]?.message ?? "Please check the message details." });
            return;
        }

        setSubmitMessage(null);
        mutation.mutate({ ...result.data, passages: passagesToUpdate });
    };

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
        <Modal show={openModal} size="md" popup onClose={closeModal}>
            <ModalHeader />
            <ModalBody>
            <div className="space-y-6">
                    <h3 className="text-xl font-medium text-gray-900 dark:text-white">Edit Message</h3>
                    {submitMessage && (
                        <Alert color={submitMessage.type === "success" ? "success" : "failure"} role={submitMessage.type === "error" ? "alert" : "status"}>
                            {submitMessage.text}
                        </Alert>
                    )}
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
                            return <Fragment key={passage.id ?? `passage-${index}`}>
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
                            </Fragment>
                        } else {
                            return <Fragment key={passage.id ?? `passage-${index}`}>
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
                            </Fragment>
                        }
                    })}
                </div>
            </ModalBody>
            <ModalFooter>
                <Button color="gray" disabled={mutation.isPending} onClick={closeModal}>Cancel</Button>
                <Button
                    className="bg-momentum-orange hover:bg-momentum-orange enabled:hover:bg-momentum-orange"
                    disabled={mutation.isPending}
                    onClick={submitUpdate}
                >
                    {mutation.isPending ? "Saving..." : "Update"}
                </Button>
            </ModalFooter>
        </Modal>
    )
}