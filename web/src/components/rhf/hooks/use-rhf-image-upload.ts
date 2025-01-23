import { useEffect, useState } from "react"
import { useFormContext } from "react-hook-form"
import useUploadImage from "@hooks/use-upload-image"

const useRHFImageUpload = (name: string) => {
	const [previews, setPreviews] = useState<Preview[]>([])
	const { setValue, watch, getValues } = useFormContext()
	const mutation = useUploadImage()
	const images: string[] = watch("image")

	useEffect(() => {
		const test: Preview[] = images?.map(image => {
			const imageSplit = image.split("/")
			const lastPart = imageSplit[imageSplit.length - 1]
			const [productID] = lastPart.split(".")
			return {
				publicID: productID,
				status: "success"
			}
		})
		setPreviews(test)
	}, [images])

	const { status, data } = mutation
	const { public_id, secure_url } = data || {}

	useEffect(() => {
		if (status === "success") {
			setValue(name, [...getValues("image"), secure_url])
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
