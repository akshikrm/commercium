import fileServer from "@utils/file"
import { AxiosResponse } from "axios"

type Upload = {
    id: number
    path: string
}

export const uploadSingleFile = async (file: File): Promise<Upload> => {
    const reqData = new FormData()
    reqData.append("file", file)
    try {
        const { data } = await fileServer.post("upload", reqData)
        return data.data
    } catch (err) {
        const { data } = err as AxiosResponse
        console.error(data)
        return Promise.reject({ message: "failed to get products" })
    }
}
