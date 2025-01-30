import { Box, IconButton } from "@mui/material"
import { CloudinaryImage } from "@cloudinary/url-gen"
import { useState } from "react"
import Render from "@components/render"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import { AdvancedImage } from "@cloudinary/react"

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

export default ImageMiniPreview
