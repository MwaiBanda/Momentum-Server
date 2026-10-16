import type { AuthUser } from "@/domain/auth/entity";
import type { AuthRepository, AuthStateListener } from "@/domain/auth/repository";
import type { AuthUseCase } from "./usecase";

export class Service implements AuthUseCase {
    public constructor(private readonly repository: AuthRepository) {}

    public currentUser(): AuthUser | null {
        return this.repository.currentUser();
    }

    public subscribe(listener: AuthStateListener): () => void {
        return this.repository.subscribe(listener);
    }

    public signIn(email: string, password: string): Promise<AuthUser> {
        return this.repository.signIn(email, password);
    }

    public signOut(): Promise<void> {
        return this.repository.signOut();
    }
}
