import type { AxiosInstance } from "axios";
import type { Transaction } from "@/domain/transaction/entity";
import type { TransactionRepository } from "@/domain/transaction/repository";

export class HttpRepository implements TransactionRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async list(): Promise<Transaction[]> {
        const response = await this.client.get<Transaction[]>("/api/v1/transactions");
        return response.data;
    }
}
