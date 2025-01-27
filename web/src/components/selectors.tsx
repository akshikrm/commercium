import useIsUser from "@hooks/auth/use-is-user"
import Render from "./render"
import { ReactNode } from "react"

type Props = {
    children: ReactNode
}

export const ShowForAdmin = ({ children }: Props) => {
    const isUser = useIsUser()
    return <Render when={!isUser} show={children} />
}
