import type { Message, MessageRequest } from "@/domain/message/entity";
import type { MessageRepository, MessageListResponse } from "@/domain/message/repository";
import type { MessageUseCase } from "./usecase";

export class Service implements MessageUseCase {
    public constructor(private readonly repository: MessageRepository) {}

    public listAdmin(): Promise<MessageListResponse> {
        return this.repository.listAdmin();
    }

    public listUnpublished(): Promise<MessageListResponse> {
        return this.repository.listUnpublished();
    }

    public create(input: MessageRequest): Promise<void> {
        return this.repository.create(input);
    }

    public update(input: Message): Promise<void> {
        return this.repository.update(input);
    }
}
