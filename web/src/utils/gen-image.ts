import { byRadius } from "@cloudinary/url-gen/actions/roundCorners"
import { scale } from "@cloudinary/url-gen/actions/resize"
import { Cloudinary, CloudinaryImage } from "@cloudinary/url-gen"

export const genImageFromPublicID = (publicId: string): CloudinaryImage => {
    const cld = new Cloudinary({ cloud: { cloudName: "commercium" } })
    return cld
        .image(publicId)
        .resize(scale().width(100).height(100))
        .roundCorners(byRadius(15))
}
