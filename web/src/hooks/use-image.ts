import { useEffect, useId, useState } from "react"
import { useFormContext } from "react-hook-form"
import useUploadImage from "@hooks/use-upload-image"

const useImage = () => {
    const { getValues, setValue, watch } = useFormContext()
    const mutation = useUploadImage()
    const publicIdList: string[] = watch("image")
    const id = useId()

    const [images, setImages] = useState<ImagePreview[]>([])
    useEffect(() => {
        setImages(
            publicIdList.map((publicId, i) => {
                return {
                    id: `${id}${i}`,
                    publicID: publicId
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
                updatedImage.publicID = data.public_id
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
        setImages([...images, { id: itemID, publicID: "" }])
        mutation.mutate(payload)
    }

    const handleDelete = (publicID: string) => {
        const publicIDs: string[] = [...getValues("image")]
        const findIndex = publicIDs.findIndex(id => id === publicID)
        publicIDs.splice(findIndex, 1)
        setValue("image", publicIDs)
    }

    return { images, handleUpload, handleDelete }
}

export default useImage
