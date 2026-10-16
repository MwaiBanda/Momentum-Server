import { Link, isRouteErrorResponse, useRouteError } from "react-router-dom";
import { HOME } from "@/application/config/constants";

function RouterState({ title, message }: { title: string; message: string }) {
    return (
        <main className="flex min-h-screen flex-col items-center justify-center gap-4 p-6 text-center">
            <h1 className="text-3xl font-bold text-gray-900">{title}</h1>
            <p className="max-w-md text-gray-600">{message}</p>
            <Link
                className="rounded-lg bg-momentum-orange px-4 py-2 font-semibold text-white hover:bg-momentum-orange"
                to={HOME}
            >
                Return to dashboard
            </Link>
        </main>
    );
}

export function RouterErrorPage() {
    const error = useRouteError();

    if (isRouteErrorResponse(error)) {
        return (
            <RouterState
                title={`${error.status} ${error.statusText}`}
                message="The dashboard could not display this page."
            />
        );
    }

    return (
        <RouterState
            title="Something went wrong"
            message="The dashboard encountered an unexpected error. Please return to the dashboard and try again."
        />
    );
}

export function NotFoundPage() {
    return (
        <RouterState
            title="Page not found"
            message="The page you requested does not exist."
        />
    );
}
