import { createContext } from "react"

export const authenticationContext = createContext<{
    user: Profile
    refresh: () => void
}>({
    user: { role: null, first_name: "", last_name: "" },
    refresh: () => null
})
