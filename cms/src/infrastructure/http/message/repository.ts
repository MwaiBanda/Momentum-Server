import type { AxiosInstance } from "axios";
import type { Message, MessageRequest } from "@/domain/message/entity";
import type { MessageListResponse, MessageRepository } from "@/domain/message/repository";
import type { components } from "@/infrastructure/http/generated/openapi";

type ApiCreateInput = components["schemas"]["message.CreateInput"];
type ApiMessageListResponse = components["schemas"]["message.MessageListResponse"];
type ApiMessageResponse = components["schemas"]["message.MessageResponse"];
type ApiUpdateInput = components["schemas"]["message.UpdateInput"];

export class HttpRepository implements MessageRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async listAdmin(): Promise<MessageListResponse> {
        const response = await this.client.get<ApiMessageListResponse>("/api/v1/messages/admin");
        return mapMessageListResponse(response.data);
    }

    public async listUnpublished(): Promise<MessageListResponse> {
        const response = await this.client.get<ApiMessageListResponse>("/api/v1/messages/unpublished");
        return mapMessageListResponse(response.data);
    }

    public async create(input: MessageRequest): Promise<void> {
        await this.client.post<ApiMessageResponse>("/api/v1/messages", toCreateInput(input));
    }

    public async update(input: Message): Promise<void> {
        await this.client.put<ApiMessageResponse>("/api/v1/messages", toUpdateInput(input));
    }
}

function mapMessageListResponse(response: ApiMessageListResponse): MessageListResponse {
    return {
        data: (response.data ?? []).map(mapMessage),
    };
}

function mapMessage(response: ApiMessageResponse): Message {
    return {
        id: response.id ?? "",
        preacher: response.preacher ?? "",
        series: response.series ?? "",
        thumbnail: response.thumbnail ?? "",
        title: response.title ?? "",
        date: response.date ?? "",
        published: response.published ?? false,
        passages: (response.passages ?? []).map((passage) => ({
            id: passage.id,
            header: passage.header,
            verse: passage.verse,
            message: passage.message,
            type: passage.type ?? "",
            order: passage.order ?? 0,
        })),
    };
}

function toCreateInput(input: MessageRequest): ApiCreateInput {
    return {
        preacher: input.preacher,
        series: input.series,
        thumbnail: input.thumbnail,
        title: input.title,
        date: input.date,
        passages: input.passages.map((passage) => ({
            header: passage.header,
            verse: passage.verse,
            message: passage.message,
            type: passage.type,
            order: passage.order,
        })),
    };
}

function toUpdateInput(input: Message): ApiUpdateInput {
    return {
        id: input.id,
        preacher: input.preacher,
        series: input.series,
        thumbnail: input.thumbnail,
        title: input.title,
        date: input.date,
        published: input.published,
        passages: input.passages.map((passage) => ({
            id: passage.id,
            header: passage.header,
            verse: passage.verse,
            message: passage.message,
        })),
    };
}
