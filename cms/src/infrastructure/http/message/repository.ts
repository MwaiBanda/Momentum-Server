import type { AxiosInstance } from "axios";
import type { Message, MessageRequest } from "@/domain/message/entity";
import type { MessageListResponse, MessageRepository } from "@/domain/message/repository";

export class HttpRepository implements MessageRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async listAdmin(): Promise<MessageListResponse> {
        const response = await this.client.get<MessageListResponse>("/api/v1/messages/admin");
        return response.data;
    }

    public async listUnpublished(): Promise<MessageListResponse> {
        const response = await this.client.get<MessageListResponse>("/api/v1/messages/unpublished");
        return response.data;
    }

    public async create(input: MessageRequest): Promise<void> {
        await this.client.post("/api/v1/messages", input);
    }

    public async update(input: Message): Promise<void> {
        await this.client.put("/api/v1/messages", input);
    }
}
