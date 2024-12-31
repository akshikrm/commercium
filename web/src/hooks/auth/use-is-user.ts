import useAuth from "@hooks/auth/use-auth"
import { useMemo } from "react"

const useIsUser = (): boolean => {
    const { user } = useAuth()
    return useMemo(() => {
        return user.role === "user"
    }, [user.role])
}

export default useIsUser
