import { Cloudinary } from "@cloudinary/url-gen"
import { scale } from "@cloudinary/url-gen/actions/resize"
import { AdvancedImage } from "@cloudinary/react"
import { Box, CircularProgress } from "@mui/material"
import Render from "@components/render"

const Image = ({ preview }: { preview: Preview }) => {
    const { publicID, status } = preview
    const cld = new Cloudinary({ cloud: { cloudName: "commercium" } })
    const img = cld.image(publicID).resize(scale().width(100).height(100))

    return (
        <Box>
            <Render
                when={status === "pending"}
                show={<CircularProgress variant='indeterminate' />}
                otherwise={<AdvancedImage cldImg={img} />}
            />
        </Box>
    )
}

export default Image
