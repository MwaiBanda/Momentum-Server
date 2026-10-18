import type { AxiosInstance } from "axios";
import type { Transaction } from "@/domain/transaction/entity";
import type { TransactionRepository } from "@/domain/transaction/repository";
import type { components } from "@/infrastructure/http/generated/openapi";
import { schemas } from "@/infrastructure/http/generated/schemas";

type ApiTransactionResponse = components["schemas"]["transaction.Response"];

export class HttpRepository implements TransactionRepository {
    public constructor(private readonly client: AxiosInstance) {}

    public async list(): Promise<Transaction[]> {
        const response = await this.client.get<ApiTransactionResponse[]>("/api/v1/transactions");
        return response.data.map((transaction) => mapTransaction(schemas.transaction_Response.parse(transaction)));
    }
}

function mapTransaction(response: ApiTransactionResponse): Transaction {
    return {
        id: response.id,
        amount: response.amount,
        date: response.date,
        createdOn: response.createdOn,
        description: response.description,
        user: {
            id: response.user.id,
            email: response.user.email,
            fullname: response.user.fullname,
            phone: response.user.phone,
        },
    };
}
