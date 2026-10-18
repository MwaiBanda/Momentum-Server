import type { AxiosInstance } from "axios";
import type { NotificationRequest } from "@/domain/notification/entity";
import type { NotificationRepository } from "@/domain/notification/repository";
import type { components } from "@/infrastructure/http/generated/openapi";
import { schemas } from "@/infrastructure/http/generated/schemas";

type ApiNotificationResponse = components["schemas"]["notification.Response"];

export class HttpRepository implements NotificationRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async send(input: NotificationRequest): Promise<void> {
        const response = await this.client.post<ApiNotificationResponse>("/api/v1/notifications", {
            title: input.title,
            body: input.body,
            topic: input.topic,
        });
        schemas.notification_Response.parse(response.data);
    }
}
