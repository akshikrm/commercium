import { useEffect, useState } from "react"
import { useFormContext } from "react-hook-form"
import useUploadImage from "@hooks/use-upload-image"

const useRHFImageUpload = (name: string) => {
    const [previews, setPreviews] = useState<Preview[]>([])
    const { setValue } = useFormContext()
    const mutation = useUploadImage()

    const { status, data } = mutation
    const { public_id } = data || {}

    useEffect(() => {
        if (status === "success") {
            setValue(name, public_id)
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
    }, [status, public_id])

    return { mutation, previews }
}

export default useRHFImageUpload
