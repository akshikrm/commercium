import { useEffect, useState } from "react"
import { useFormContext } from "react-hook-form"
import useUploadImage from "@hooks/use-upload-image"

const useRHFImageUpload = (name: string) => {
    const [previews, setPreviews] = useState<Preview[]>([])
    const { setValue, watch } = useFormContext()
    const mutation = useUploadImage()

    const image: string = watch("image")

    useEffect(() => {
        if (image.length > 0) {
            const imageSplit = image.split("/")
            const lastPart = imageSplit[imageSplit.length - 1]
            if (lastPart) {
                const [productID] = lastPart.split(".")
                setPreviews([
                    {
                        publicID: productID,
                        status: "success"
                    }
                ])
            }
        }
    }, [image])

    const { status, data } = mutation
    const { public_id, secure_url } = data || {}

    useEffect(() => {
        if (status === "success") {
            setValue(name, secure_url)
            setPreviews(prev => {
                const temp = [...prev]
                const selectedIndex = temp.findIndex(
                    ({ publicID }) => publicID === ""
                )
                temp.splice(selectedIndex, 1, {
                    publicID: public_id || "",
                    status: "success"
                })
                return temp
            })
        }

        if (status === "pending") {
            setPreviews([...previews, { publicID: "", status: "pending" }])
        }
    }, [status, public_id, secure_url])

    return { mutation, previews }
}

export default useRHFImageUpload
