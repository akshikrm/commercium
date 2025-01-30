import { scale } from "@cloudinary/url-gen/actions/resize"
import { AdvancedImage } from "@cloudinary/react"
import { Cloudinary, CloudinaryImage } from "@cloudinary/url-gen"
import { Box, Button, IconButton, Stack, TextField } from "@mui/material"
import { FunctionComponent, useEffect, useId, useState } from "react"
import { Controller, useFormContext } from "react-hook-form"
import RenderList from "@components/render-list"
import useUploadImage from "@hooks/use-upload-image"
import Render from "@components/render"
import { byRadius } from "@cloudinary/url-gen/actions/roundCorners"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"

type Props = {
    label: string
    name: string
}

const genImageFromPublicID = (publicId: string): CloudinaryImage => {
    const cld = new Cloudinary({ cloud: { cloudName: "commercium" } })
    return cld
        .image(publicId)
        .resize(scale().width(100).height(100))
        .roundCorners(byRadius(15))
}

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
    const { control, getValues, setValue } = useFormContext()
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

    return (
        <>
            <Stack
                sx={{
                    marginBottom: 2,
                    overflowX: "scroll",
                    height: 100
                }}
                direction='row'
            >
                <RenderList
                    list={images}
                    render={image => {
                        return <ImageMiniPreview image={image.image} />
                    }}
                />
            </Stack>
            <Box>
                <Button
                    variant='outlined'
                    fullWidth
                    onClick={() => {
                        const upload: HTMLInputElement | null =
                            document.querySelector("input[type='file']")
                        if (upload) {
                            upload.click()
                        }
                    }}
                >
                    click to upload image
                </Button>
            </Box>
            <Controller
                name={name}
                control={control}
                render={({ fieldState }) => {
                    const { error, invalid } = fieldState
                    return (
                        <>
                            <TextField
                                type='file'
                                sx={{ display: "none" }}
                                onChange={async e => {
                                    const { files } =
                                        e.target as HTMLInputElement
                                    if (files?.length) {
                                        const itemID = `${id}${images.length}`
                                        const payload = {
                                            id: itemID,
                                            file: files[0]
                                        }

                                        setImages([
                                            ...images,
                                            { id: itemID, image: "" }
                                        ])
                                        mutation.mutate(payload)
                                    }
                                }}
                                label={label}
                                error={invalid}
                                helperText={error?.message}
                            />
                        </>
                    )
                }}
            />
        </>
    )
}

const ImageMiniPreview = ({ image }: { image: CloudinaryImage | string }) => {
    const [onHover, setOnHover] = useState(false)
    console.log(image)
    return (
        <Box
            sx={{
                height: 100,
                width: 100,
                position: "relative"
            }}
            onMouseEnter={() => {
                setOnHover(true)
            }}
            onMouseLeave={() => {
                setOnHover(false)
            }}
        >
            <Render
                when={image !== ""}
                show={
                    <>
                        <Render
                            when={onHover}
                            show={
                                <Box
                                    sx={{
                                        position: "absolute",
                                        top: 0,
                                        bottom: 0,
                                        left: 0,
                                        right: 0,
                                        borderRadius: "15px",
                                        backgroundColor: "#c1c1c13b",
                                        cursor: "pointer",
                                        display: "flex",
                                        justifyContent: "center",
                                        alignItems: "center"
                                    }}
                                >
                                    <IconButton size='small'>
                                        <RenderIcon icon={icons.delete} />
                                    </IconButton>
                                </Box>
                            }
                        />
                        <AdvancedImage cldImg={image as CloudinaryImage} />
                    </>
                }
                otherwise={
                    <Box
                        sx={{
                            position: "absolute",
                            top: 0,
                            bottom: 0,
                            left: 0,
                            right: 0,
                            borderRadius: "15px",
                            backgroundColor: "#c1c1c111",
                            cursor: "pointer",
                            display: "flex",
                            justifyContent: "center",
                            alignItems: "center"
                        }}
                    >
                        <IconButton size='small'>
                            <RenderIcon icon={icons.animated.loading} />
                        </IconButton>
                    </Box>
                }
            />
        </Box>
    )
}

export default RHFImageUpload
