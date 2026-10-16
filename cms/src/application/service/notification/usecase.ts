import type { NotificationRequest } from "@/domain/notification/entity";

export interface NotificationUseCase {
    send(input: NotificationRequest): Promise<void>;
}
