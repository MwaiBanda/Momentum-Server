import { createContext, useContext } from "react";
import type { ApplicationServices } from "@/application/services";

export const ApplicationContext = createContext<ApplicationServices | null>(null);

export function useApplication(): ApplicationServices {
    const application = useContext(ApplicationContext);
    if (!application) {
        throw new Error("Application services are not available");
    }
    return application;
}
