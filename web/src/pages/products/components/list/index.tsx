import { Box, Button, Card, Stack, Typography } from "@mui/material"
import RenderList from "@/components/render-list"
import { Link } from "react-router"
import { FunctionComponent } from "react"
import { ADMIN_PATHS } from "@/paths"
import { BASE_URL_FILE } from "@config"

type Props = {
    products?: Product[]
    onDelete: (id: number) => void
}
const List: FunctionComponent<Props> = ({ products, onDelete }) => {
    return (
        <Stack>
            <RenderList
                list={products}
                render={product => {
                    const { id, image, name, description } = product
                    const completeImageURI = [BASE_URL_FILE, image].join("/")
                    return (
                        <Card key={id}>
                            <Stack direction='row' alignItems='center'>
                                <img
                                    src={completeImageURI}
                                    width={50}
                                    height={50}
                                />

                                <Stack
                                    direction='row'
                                    alignItems='center'
                                    width='100%'
                                    justifyContent='space-between'
                                >
                                    <Box>
                                        <Typography variant='h6'>
                                            {name}
                                        </Typography>
                                        <Typography variant='body1'>
                                            {description}
                                        </Typography>
                                    </Box>
                                    <Stack spacing={2} direction='row'>
                                        <Button
                                            color='warning'
                                            component={Link}
                                            to={ADMIN_PATHS.products.edit(id)}
                                        >
                                            edit
                                        </Button>
                                        <Button
                                            color='error'
                                            onClick={() => onDelete(id)}
                                        >
                                            delete
                                        </Button>
                                    </Stack>
                                </Stack>
                            </Stack>
                        </Card>
                    )
                }}
            />
        </Stack>
    )
}

export default List
