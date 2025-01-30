type Preview = {
    publicID: string
    status: "error" | "idle" | "pending" | "success"
}

type FileUploadPayload = {
    id: string
    file: File
}

type ImagePreview = {
    id: string
    image: CloudinaryImage
}
