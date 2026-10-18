import axios from "axios";
import { schemas } from "@/infrastructure/http/generated/schemas";

export class ApiError extends Error {
    public constructor(
        public readonly code: number,
        message: string,
    ) {
        super(message);
        this.name = "ApiError";
    }
}

export function normalizeApiError(error: unknown): unknown {
    if (!axios.isAxiosError(error) || !error.response) {
        return error;
    }

    const parsed = schemas.http_ErrorResponse.safeParse(error.response.data);
    if (!parsed.success) {
        return error;
    }

    return new ApiError(parsed.data.code, parsed.data.message);
}
