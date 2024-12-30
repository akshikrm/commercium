import { Typography } from "@mui/material"

const ItemLabel = ({ label }: { label: string }) => {
    return (
        <Typography
            variant='subtitle2'
            color='textSecondary'
            sx={{
                textTransform: "capitalize"
            }}
        >
            {label}
        </Typography>
    )
}

export default ItemLabel
