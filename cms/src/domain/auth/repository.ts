import type { AuthUser } from "./entity";

export type AuthStateListener = (user: AuthUser | null) => void;

export interface AuthRepository {
    currentUser(): AuthUser | null;
    subscribe(listener: AuthStateListener): () => void;
    signIn(email: string, password: string): Promise<AuthUser>;
    signOut(): Promise<void>;
}
