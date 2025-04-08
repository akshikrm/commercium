import { Box, Button, Stack, TextField } from "@mui/material"
import { FunctionComponent } from "react"
import { Controller, useFormContext } from "react-hook-form"
import RenderList from "@components/render-list"
import useImage from "@hooks/use-image"
import ImageMiniPreview from "./components/mini-image-preview"

type Props = {
    label: string
    name: string
}

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
    const { control } = useFormContext()
    const { images, handleUpload, handleDelete } = useImage()

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
                        return (
                            <ImageMiniPreview
                                publicID={image.publicID}
                                handleDelete={handleDelete}
                            />
                        )
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

export default RHFImageUpload
