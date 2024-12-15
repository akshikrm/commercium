import { ReactNode } from "react"

type Props<T> = {
    list?: T[]
    render: (item: T, i?: number) => ReactNode
}
const RenderList = <T,>({ list = [], render }: Props<T>) => {
    if (list.length > 0) {
        return list.map(render)
    }
    return null
}

export default RenderList
