import type { NotificationRequest } from "./entity";

export interface NotificationRepository {
    send(input: NotificationRequest): Promise<void>;
}
