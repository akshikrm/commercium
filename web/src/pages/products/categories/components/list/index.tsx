import { ADMIN_PATHS } from "@/paths"
import RenderList from "@components/render-list"
import { Box, Button, Card, Stack, Typography } from "@mui/material"
import { FunctionComponent } from "react"
import { Link } from "react-router"

type Props = {
    categories?: ProductCategory[]
    onDelete: (id: number) => void
}

const List: FunctionComponent<Props> = ({ categories: list, onDelete }) => {
    return (
        <Stack>
            <RenderList
                list={list}
                render={category => {
                    return (
                        <Card key={category.id}>
                            <Stack
                                direction='row'
                                alignItems='center'
                                justifyContent='space-between'
                            >
                                <Box>
                                    <Typography variant='h6'>
                                        {category.name}
                                    </Typography>
                                    <Typography variant='body1'>
                                        {category.description}
                                    </Typography>
                                </Box>
                                <Stack spacing={2} direction='row'>
                                    <Button
                                        color='warning'
                                        component={Link}
                                        to={ADMIN_PATHS.products.categories.edit(
                                            category.id
                                        )}
                                    >
                                        edit
                                    </Button>
                                    <Button
                                        color='error'
                                        onClick={() => onDelete(category.id)}
                                    >
                                        delete
                                    </Button>
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
