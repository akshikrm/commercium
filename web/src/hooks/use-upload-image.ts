import { useMutation } from "@tanstack/react-query"
import { files } from "@api"

const useUploadImage = () => {
    const mutation = useMutation({
        mutationFn: async (payload: FileUploadPayload) =>
            await files.single(payload.file)
    })

    return mutation
}

export default useUploadImage
