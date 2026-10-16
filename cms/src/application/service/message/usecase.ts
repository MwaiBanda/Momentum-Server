import type { Message, MessageRequest } from "@/domain/message/entity";
import type { MessageListResponse } from "@/domain/message/repository";

export interface MessageUseCase {
    listAdmin(): Promise<MessageListResponse>;
    listUnpublished(): Promise<MessageListResponse>;
    create(input: MessageRequest): Promise<void>;
    update(input: Message): Promise<void>;
}
