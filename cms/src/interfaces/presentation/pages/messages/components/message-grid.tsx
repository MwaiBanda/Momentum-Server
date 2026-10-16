import { Card } from "flowbite-react";
import { useState } from "react";
import type { Message } from "@/domain/message/entity";

type MessageGridProps = {
    messages: Message[];
    onEditMessage: (message: Message) => void;
    emptyMessage: string;
};

export function MessageGrid({ messages, onEditMessage, emptyMessage }: MessageGridProps) {
    if (messages.length === 0) {
        return <p className="mt-6 text-sm text-gray-500" role="status">{emptyMessage}</p>;
    }

    return (
        <div className="grid w-full grid-cols-2 gap-2 md:grid-cols-3 lg:grid-cols-5 lg:gap-4">
            {messages.map((message) => (
                <MessageCard key={message.id} message={message} onEditMessage={onEditMessage} />
            ))}
        </div>
    );
}

function MessageCard({ message, onEditMessage }: { message: Message; onEditMessage: (message: Message) => void }) {
    const [imageFailed, setImageFailed] = useState(false);

    return (
        <Card
            className="mt-6 max-w-lg"
            renderImage={() =>
                imageFailed || !message.thumbnail ? (
                    <div
                        className="flex h-[140px] items-center justify-center bg-gray-100 text-sm text-gray-500 md:h-[160px] lg:h-[180px]"
                        role="img"
                        aria-label={`No thumbnail available for ${message.title}`}
                    >
                        No image available
                    </div>
                ) : (
                    <img
                        className="max-h-[140px] object-cover md:max-h-[160px] lg:max-h-[180px]"
                        src={message.thumbnail}
                        alt={`${message.title} thumbnail`}
                        loading="lazy"
                        onError={() => setImageFailed(true)}
                    />
                )
            }
            horizontal={false}
        >
            <div className="flex w-full min-w-0 items-start justify-between gap-2">
                <div className="min-w-0 flex-1">
                    <h5 className="line-clamp-1 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">
                        {message.preacher}
                    </h5>
                    <p className="line-clamp-1 font-normal text-gray-700 dark:text-gray-400">
                        {message.date}
                    </p>
                    <p className="line-clamp-1 font-normal text-gray-700 dark:text-gray-400">
                        {message.series}
                    </p>
                    <p className="line-clamp-1 font-normal text-gray-700 dark:text-gray-400">
                        {message.title}
                    </p>
                </div>
                <button
                    type="button"
                    aria-label={`Edit ${message.title}`}
                    onClick={() => onEditMessage(message)}
                    className="max-h-11 shrink-0 rounded-full bg-momentum-orange p-2 font-bold text-white hover:bg-momentum-orange"
                >
                    <svg className="ml-1 h-6 w-6 text-white" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none">
                        <path stroke="currentColor" strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="m14.304 4.844 2.852 2.852M7 7H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4.5m2.409-9.91a2.017 2.017 0 0 1 0 2.853l-6.844 6.844L8 14l.713-3.565 6.844-6.844a2.015 2.015 0 0 1 2.852 0Z" />
                    </svg>
                </button>
            </div>
        </Card>
    );
}
