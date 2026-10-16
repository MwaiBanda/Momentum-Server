import type { AxiosInstance } from "axios";
import type { NotificationRequest } from "@/domain/notification/entity";
import type { NotificationRepository } from "@/domain/notification/repository";

export class HttpRepository implements NotificationRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async send(input: NotificationRequest): Promise<void> {
        await this.client.post("/api/v1/notifications", input);
    }
}
