import type { Transaction } from "@/domain/transaction/entity";

export interface TransactionUseCase {
    list(): Promise<Transaction[]>;
}
