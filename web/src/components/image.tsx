import { Cloudinary } from "@cloudinary/url-gen"
import { scale } from "@cloudinary/url-gen/actions/resize"
import { AdvancedImage } from "@cloudinary/react"
import { Box, IconButton, Stack } from "@mui/material"
import RenderIcon from "./render-icon"
import icons from "@/icons"

type Props = {
    publicID: string
    onClick?: () => void
}

const Image = ({ publicID, onClick }: Props) => {
    const cld = new Cloudinary({ cloud: { cloudName: "commercium" } })
    const img = cld.image(publicID).resize(scale().width(100).height(100))

    return (
        <Stack justifyContent='flex-end'>
            {/* <Box sx={{ textAlign: "right" }}> */}
            {/*     <IconButton size='small' onClick={onClick}> */}
            {/*         <RenderIcon icon={icons.close} /> */}
            {/*     </IconButton> */}
            {/* </Box> */}
            <AdvancedImage cldImg={img} />
        </Stack>
    )
}

export default Image
