import type { Message } from "@/domain/message/entity";

const searchableFields = (message: Message): string[] => [
    message.title,
    message.preacher,
    message.date,
    message.series,
];

export function filterMessages(messages: Message[], searchText: string): Message[] {
    const normalizedSearch = searchText.trim().toLowerCase();

    if (!normalizedSearch) {
        return messages;
    }

    return messages.filter((message) =>
        searchableFields(message).some((field) =>
            field.toLowerCase().includes(normalizedSearch),
        ),
    );
}
