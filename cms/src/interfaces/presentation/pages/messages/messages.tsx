import { useQuery } from "@tanstack/react-query";
import { TextInput } from "flowbite-react";
import { useEffect, useState } from "react";
import { Rings } from "react-loader-spinner";
import { useApplication } from "@/interfaces/presentation/application-context";
import type { Message } from "@/domain/message/entity";
import type { MessageListResponse } from "@/domain/message/repository";
import { MomentumNavigation } from "../../components/navigation";
import AddMessageModal from "./modals/add-message";
import EditMessageModal from "./modals/edit-message";
import { filterMessages } from "./components/message-filter";
import { MessageGrid } from "./components/message-grid";

const searchTags = ["preacher", "title", "series", "date"];

export const Messages = () => {
    const application = useApplication();
    const [showAuthModal, setShowAuthModal] = useState(false);
    const [searchText, setSearchText] = useState("");
    const [searchTagIndex, setSearchTagIndex] = useState(0);
    const [openModal, setOpenMessageModal] = useState(false);
    const [openEditModal, setOpenEditMessageModal] = useState(false);
    const [editableMessage, setEditableMessage] = useState<Message | null>(null);

    const publishedQuery = useQuery<MessageListResponse>({
        queryKey: ["message-admin"],
        queryFn: () => application.message.listAdmin(),
    });
    const unpublishedQuery = useQuery<MessageListResponse>({
        queryKey: ["message-unpublished"],
        queryFn: () => application.message.listUnpublished(),
    });

    useEffect(() => {
        const interval = window.setInterval(() => {
            setSearchTagIndex((currentIndex) => (currentIndex + 1) % searchTags.length);
        }, 2000);

        return () => window.clearInterval(interval);
    }, []);

    const isLoading = publishedQuery.isPending || unpublishedQuery.isPending;
    const hasError = publishedQuery.isError || unpublishedQuery.isError;
    const filteredPublishedMessages = filterMessages(publishedQuery.data?.data ?? [], searchText);
    const filteredUnpublishedMessages = filterMessages(unpublishedQuery.data?.data ?? [], searchText);

    const retryQueries = () => {
        void Promise.all([publishedQuery.refetch(), unpublishedQuery.refetch()]);
    };

    const openEditor = (message?: Message) => {
        if (!application.auth.currentUser()) {
            setShowAuthModal(true);
            return;
        }

        if (message) {
            setEditableMessage(message);
            setOpenEditMessageModal(true);
        } else {
            setOpenMessageModal(true);
        }
    };

    if (isLoading) {
        return (
            <div className="flex h-screen w-full flex-col items-center justify-center" role="status" aria-label="Loading messages">
                <Rings visible height="150" width="150" color="#ec750c" ariaLabel="loading messages" />
            </div>
        );
    }

    if (hasError) {
        return (
            <div className="flex h-screen w-full flex-col items-center justify-center gap-4 p-6 text-center" role="alert">
                <h1 className="text-xl font-semibold text-gray-900">Messages are temporarily unavailable</h1>
                <p className="text-gray-600">We could not load the message library. Please try again.</p>
                <button
                    type="button"
                    onClick={retryQueries}
                    className="rounded-lg bg-momentum-orange px-4 py-2 font-semibold text-white hover:bg-momentum-orange"
                >
                    Try again
                </button>
            </div>
        );
    }

    return (
        <main className="h-screen w-full">
            <MomentumNavigation
                content={
                    <>
                        <div className="flex w-full items-end justify-center">
                            <label htmlFor="message-search" className="sr-only">Search messages</label>
                            <TextInput
                                id="message-search"
                                className="ml-4 w-[240px] rounded-none lg:w-[500px]"
                                type="search"
                                placeholder={`Search by "${searchTags[searchTagIndex]}"`}
                                value={searchText}
                                onChange={(event) => setSearchText(event.target.value)}
                            />
                            <button
                                type="button"
                                className="z-10 ml-[-5px] max-h-[42px] rounded-none rounded-r-lg bg-momentum-orange px-4 py-2.5 font-bold text-white hover:bg-momentum-orange"
                                onClick={() => openEditor()}
                            >
                                Add Message
                            </button>
                        </div>

                        <section aria-labelledby="private-messages-heading">
                            <h2 id="private-messages-heading" className="mt-2 text-lg font-bold">Private</h2>
                            <p className="text-sm text-gray-500">View and edit staged content prior to publishing.</p>
                            <MessageGrid
                                messages={filteredUnpublishedMessages}
                                onEditMessage={openEditor}
                                emptyMessage="No staged messages match your search."
                            />
                        </section>

                        <section aria-labelledby="public-messages-heading">
                            <h2 id="public-messages-heading" className="mt-5 text-lg font-bold">Public</h2>
                            <p className="text-sm text-gray-500">View and edit currently published content.</p>
                            <MessageGrid
                                messages={filteredPublishedMessages}
                                onEditMessage={openEditor}
                                emptyMessage="No published messages match your search."
                            />
                        </section>
                    </>
                }
                showAuth={showAuthModal}
                onShowAuthChange={setShowAuthModal}
            />
            <AddMessageModal openModal={openModal} setOpenModal={setOpenMessageModal} />
            <EditMessageModal
                key={editableMessage?.id ?? "edit-message"}
                openModal={openEditModal}
                setOpenModal={setOpenEditMessageModal}
                message={editableMessage}
            />
        </main>
    );
};
