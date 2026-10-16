import {RouterProvider} from "react-router-dom";
import { router } from "@/interfaces/router/router";
import { AppProviders } from "./presentation/providers";
import type { ApplicationServices } from "@/application/services";

function App({ application }: { application: ApplicationServices }) {
  return (
    <AppProviders application={application}>
      <RouterProvider router={router} />
    </AppProviders>
  )
}

export default App
