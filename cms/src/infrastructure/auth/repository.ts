import { signInWithEmailAndPassword, signOut } from "firebase/auth";
import { auth } from "./firebase";
import type { AuthUser } from "@/domain/auth/entity";
import type { AuthRepository, AuthStateListener } from "@/domain/auth/repository";

function mapUser(user: typeof auth.currentUser): AuthUser | null {
    return user ? { email: user.email } : null;
}

export class FirebaseRepository implements AuthRepository {
    public currentUser(): AuthUser | null {
        return mapUser(auth.currentUser);
    }

    public subscribe(listener: AuthStateListener): () => void {
        return auth.onAuthStateChanged((user) => listener(mapUser(user)));
    }

    public async signIn(email: string, password: string): Promise<AuthUser> {
        const result = await signInWithEmailAndPassword(auth, email, password);
        const user = mapUser(result.user);
        if (!user) {
            throw new Error("Firebase returned no authenticated user");
        }
        return user;
    }

    public async signOut(): Promise<void> {
        await signOut(auth);
    }
}
