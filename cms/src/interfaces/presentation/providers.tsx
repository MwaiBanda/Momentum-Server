import type { ReactNode } from "react";
import { QueryClientProvider } from "react-query";
import { ReactQueryDevtools } from "react-query/devtools";
import { queryClient } from "@/infrastructure/query/client";
import type { ApplicationServices } from "@/application/services";
import { ApplicationContext } from "./application-context";

export function AppProviders({ children, application }: { children: ReactNode; application: ApplicationServices }) {
    return (
        <QueryClientProvider client={queryClient}>
            <ApplicationContext.Provider value={application}>
                {children}
            </ApplicationContext.Provider>
            <ReactQueryDevtools initialIsOpen={false} />
        </QueryClientProvider>
    );
}
