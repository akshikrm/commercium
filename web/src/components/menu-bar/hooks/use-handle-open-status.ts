import { useLocation } from "react-router"
import { useEffect, useState } from "react"

const useHandleOpenStatus = (children: Paths[] = []): [boolean, () => void] => {
    const [open, setOpen] = useState(false)
    const { pathname } = useLocation()

    useEffect(() => {
        if (children?.length && !open) {
            for (let child of children) {
                const isOpen = child.path === pathname
                setOpen(isOpen)
                if (isOpen) {
                    return
                }
            }
        }
    }, [pathname, children])

    const toggleOpen = () => setOpen(!open)

    return [open, toggleOpen]
}

export default useHandleOpenStatus
