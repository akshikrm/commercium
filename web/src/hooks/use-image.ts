import { useEffect, useId, useState } from "react"
import { useFormContext } from "react-hook-form"
import useUploadImage from "@hooks/use-upload-image"
import { genImageFromPublicID } from "@utils/gen-image"

const useImage = () => {
    const { getValues, setValue } = useFormContext()
    const mutation = useUploadImage()
    const publicIdList: string[] = getValues("image")
    const id = useId()

    const [images, setImages] = useState<ImagePreview[]>([])
    useEffect(() => {
        setImages(
            publicIdList.map((publicId, i) => {
                return {
                    id: `${id}${i}`,
                    image: publicId ? genImageFromPublicID(publicId) : ""
                }
            })
        )
    }, [publicIdList])

    const { status, data, variables } = mutation

    useEffect(() => {
        if (status === "success") {
            setValue("image", [...getValues("image"), data.public_id])
            setImages(prev => {
                const temp = [...prev]
                const imageIndex = temp.findIndex(
                    ({ id }) => id === variables.id
                )
                const updatedImage = temp[imageIndex]
                updatedImage.image = genImageFromPublicID(data.public_id)
                temp.splice(imageIndex, 1, updatedImage)
                return temp
            })
            if (!getValues("primary_image")) {
                setValue("primary_image", data.secure_url)
            }
        }
    }, [status, data, getValues, variables])

    const handleUpload = (files: FileList) => {
        const itemID = `${id}${images.length}`
        const payload = {
            id: itemID,
            file: files[0]
        }

        setImages([...images, { id: itemID, image: "" }])
        mutation.mutate(payload)
    }

    return { images, handleUpload }
}

export default useImage
