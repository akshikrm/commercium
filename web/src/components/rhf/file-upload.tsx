import { Box, Button, Stack, TextField, Typography } from "@mui/material"
import { FunctionComponent, useEffect, useRef } from "react"
import { Controller, useFormContext } from "react-hook-form"
import RenderList from "@components/render-list"
import Image from "@components/image"
import useUploadImage from "@hooks/use-upload-image"

type Props = {
    label: string
    name: string
}

const RHFImageUpload: FunctionComponent<Props> = ({ name, label }) => {
    const { control, watch, setValue, getValues } = useFormContext()
    const mutation = useUploadImage()
    const images: string[] = watch("image")

    const { status, data } = mutation

    useEffect(() => {
        if (status === "success") {
            setValue("image", [...getValues("image"), data.public_id])
        }
    }, [status, data])

    return (
        <>
            <Box
                sx={{
                    cursor: "pointer",
                    border: "2px dashed grey",
                    borderRadius: "10px",
                    padding: "10px 10px",
                    fontSize: "16px",
                    textAlign: "center"
                }}
                onClick={() => {
                    const upload: HTMLInputElement | null =
                        document.querySelector("input[type='file']")

                    if (upload) {
                        upload.click()
                    }
                }}
            >
                <Typography variant='h6' color='textSecondary'>
                    Drag and drop an image to upload
                </Typography>
            </Box>
            <Stack direction='row'>
                <RenderList
                    list={images}
                    render={image => (
                        <Image
                            key={image}
                            publicID={image}
                            onClick={() => {
                                const currentState: string[] =
                                    getValues("image")
                                const selectedIndex = currentState.findIndex(
                                    (cur: string) => {
                                        return cur === image
                                    }
                                )
                                currentState.splice(selectedIndex, 1)
                                setValue("image", currentState)
                            }}
                        />
                    )}
                />
            </Stack>
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
                                        mutation.mutate(files)
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
