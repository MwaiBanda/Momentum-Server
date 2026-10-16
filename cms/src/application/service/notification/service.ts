import type { NotificationRequest } from "@/domain/notification/entity";
import type { NotificationRepository } from "@/domain/notification/repository";
import type { NotificationUseCase } from "./usecase";

export class Service implements NotificationUseCase {
    public constructor(private readonly repository: NotificationRepository) {}

    public send(input: NotificationRequest): Promise<void> {
        return this.repository.send(input);
    }
}
