import type { Transaction } from "@/domain/transaction/entity";
import type { TransactionRepository } from "@/domain/transaction/repository";
import type { TransactionUseCase } from "./usecase";

export class Service implements TransactionUseCase {
    public constructor(private readonly repository: TransactionRepository) {}

    public list(): Promise<Transaction[]> {
        return this.repository.list();
    }
}
