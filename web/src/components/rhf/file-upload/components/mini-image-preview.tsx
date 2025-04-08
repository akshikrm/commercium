import { Box, IconButton } from "@mui/material"
import { useState } from "react"
import Render from "@components/render"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import { AdvancedImage } from "@cloudinary/react"
import { genImageFromPublicID } from "@utils/gen-image"

type Props = {
    publicID: string
    handleDelete: (publicID: string) => void
}

const ImageMiniPreview = ({ publicID, handleDelete }: Props) => {
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
                when={publicID !== ""}
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
                                    <IconButton
                                        size='small'
                                        onClick={() => {
                                            handleDelete(publicID)
                                        }}
                                    >
                                        <RenderIcon icon={icons.delete} />
                                    </IconButton>
                                </Box>
                            }
                        />
                        <AdvancedImage
                            cldImg={genImageFromPublicID(publicID)}
                        />
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

export default ImageMiniPreview
