import { useMutation } from "@tanstack/react-query"
import { files } from "@api"

const useUploadImage = () => {
    const mutation = useMutation({
        mutationFn: async (payload: FileList) => await files.single(payload[0])
    })

    return mutation
}

export default useUploadImage
