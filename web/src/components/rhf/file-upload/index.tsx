import { Box, Button, IconButton, Stack, TextField } from "@mui/material"
import { CloudinaryImage } from "@cloudinary/url-gen"
import { FunctionComponent, useState } from "react"
import { Controller, useFormContext } from "react-hook-form"
import RenderList from "@components/render-list"
import Render from "@components/render"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import { AdvancedImage } from "@cloudinary/react"
import useImage from "@hooks/use-image"

type Props = {
    label: string
    name: string
}

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
    const { control } = useFormContext()
    const { images, handleUpload } = useImage()

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
                                        handleUpload(files)
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
