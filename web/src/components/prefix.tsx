import { ReactNode } from "react"

export const Currency = ({ children }: { children: ReactNode }) => {
    return <>&#8377;{children || 0}</>
}
