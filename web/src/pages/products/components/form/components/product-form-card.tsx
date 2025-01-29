import { Card, Typography } from "@mui/material"
import { ReactNode } from "react"

const ProductFormCard = ({
    children,
    title
}: {
    title: string
    children: ReactNode
}) => {
    return (
        <Card sx={{ height: "100%" }}>
            <Typography
                variant='subtitle2'
                color='textSecondary'
                sx={{ marginBottom: 1 }}
            >
                {title}
            </Typography>
            {children}
        </Card>
    )
}

export default ProductFormCard
