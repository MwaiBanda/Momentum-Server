import type { Transaction } from "./entity";

export interface TransactionRepository {
    list(): Promise<Transaction[]>;
}
