import type { Message, MessageRequest } from "./entity";

export type MessageListResponse = {
    data: Message[];
};

export interface MessageRepository {
    listAdmin(): Promise<MessageListResponse>;
    listUnpublished(): Promise<MessageListResponse>;
    create(input: MessageRequest): Promise<void>;
    update(input: Message): Promise<void>;
}
