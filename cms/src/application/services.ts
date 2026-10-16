import type { AuthUseCase } from "./service/auth/usecase";
import type { MessageUseCase } from "./service/message/usecase";
import type { NotificationUseCase } from "./service/notification/usecase";
import type { TransactionUseCase } from "./service/transaction/usecase";

export type ApplicationServices = {
    auth: AuthUseCase;
    message: MessageUseCase;
    notification: NotificationUseCase;
    transaction: TransactionUseCase;
};
