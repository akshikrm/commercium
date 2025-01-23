import cloudinary from "@utils/file"

type Upload = {
    asset_id: number
    url: string
    secure_url: string
    public_id: string
}

export const uploadSingleFile = async (file: File): Promise<Upload> => {
    const reqData = new FormData()
    reqData.append("file", file)
    reqData.append("upload_preset", "ml_default")
    try {
        const { data } = await cloudinary.post("image/upload", reqData)
        return data
    } catch (err) {
        return Promise.reject({ message: "failed to get products" })
    }
}
