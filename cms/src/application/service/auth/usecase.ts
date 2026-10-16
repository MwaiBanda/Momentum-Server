import type { AuthStateListener } from "@/domain/auth/repository";
import type { AuthUser } from "@/domain/auth/entity";

export interface AuthUseCase {
    currentUser(): AuthUser | null;
    subscribe(listener: AuthStateListener): () => void;
    signIn(email: string, password: string): Promise<AuthUser>;
    signOut(): Promise<void>;
}
