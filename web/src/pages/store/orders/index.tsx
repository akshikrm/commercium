import HeaderBreadcrumbs from "@components/header"
import Table from "@mui/material/Table"
import TableBody from "@mui/material/TableBody"
import TableCell from "@mui/material/TableCell"
import TableContainer from "@mui/material/TableContainer"
import TableHead from "@mui/material/TableHead"
import TableRow from "@mui/material/TableRow"
import Paper from "@mui/material/Paper"
import dayjs from "dayjs"
import { DATE_VIEW_FORMAT } from "@config"
import RenderList from "@components/render-list"
import { Currency } from "@components/prefix"
import useGetOrders from "@hooks/orders/use-get-orders"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import IconButton from "@mui/material/IconButton"
import { useMemo, useState } from "react"
import { Box, Button, Chip, Popover, Stack, Typography } from "@mui/material"
import { order } from "@api"
import useGetStatusColor from "@hooks/shipping/use-get-status-color"
import Render from "@components/render"
import useCompleteTransaction from "@hooks/orders/use-complete-transaction"

const Orders = () => {
    const { data: orders, refetch } = useGetOrders()
    const [orderItems, setOrderItems] = useState<OrderItems[]>([])
    const [menuEl, setMenuEl] = useState<Element | null>(null)
    const completeOrder = useCompleteTransaction(async () => {
        await refetch()
    })

    const handleDownload = async (txnID: string) => {
        const data = await order.getOrderByID(txnID)
        window.open(data)
    }

    return (
        <>
            <HeaderBreadcrumbs
                heading='Purchase'
                links={[
                    { label: "Home", href: "/" },
                    { label: "Purchase", href: "/" }
                ]}
            />
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label='simple table'>
                    <TableHead>
                        <TableRow>
                            <TableCell>Invoice Number</TableCell>
                            <TableCell>Payment Status</TableCell>
                            <TableCell>Items</TableCell>
                            <TableCell>Price</TableCell>
                            <TableCell>Purchased On</TableCell>
                            <TableCell>Action</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <RenderList
                            list={orders}
                            render={row => {
                                return (
                                    <TableRow
                                        key={row.id}
                                        sx={{
                                            "&:last-child td, &:last-child th":
                                                {
                                                    border: 0
                                                }
                                        }}
                                    >
                                        <TableCell>
                                            #{row.invoice_number}
                                        </TableCell>
                                        <TableCell>
                                            {row.payment_status}
                                        </TableCell>
                                        <TableCell>
                                            <PurchaseItem
                                                products={row.products}
                                                onClickMoreItem={e => {
                                                    setMenuEl(e.currentTarget)
                                                    setOrderItems(row.products)
                                                }}
                                            />
                                        </TableCell>
                                        <TableCell>
                                            <Currency amount={row.total} />
                                        </TableCell>
                                        <TableCell>
                                            {dayjs(row.created_at).format(
                                                DATE_VIEW_FORMAT
                                            )}
                                        </TableCell>
                                        <TableCell>
                                            <Render
                                                when={
                                                    row.payment_status ===
                                                    "completed"
                                                }
                                                show={
                                                    <IconButton
                                                        onClick={() =>
                                                            handleDownload(
                                                                row.transaction_id
                                                            )
                                                        }
                                                    >
                                                        <RenderIcon
                                                            icon={
                                                                icons.download
                                                            }
                                                        />
                                                    </IconButton>
                                                }
                                                otherwise={
                                                    <Button
                                                        onClick={() => {
                                                            completeOrder(
                                                                row.transaction_id
                                                            )
                                                        }}
                                                    >
                                                        complete
                                                    </Button>
                                                }
                                            />
                                        </TableCell>
                                    </TableRow>
                                )
                            }}
                        />
                    </TableBody>
                </Table>
            </TableContainer>
            <Popover
                open={Boolean(menuEl)}
                anchorEl={menuEl}
                onClose={() => {
                    setMenuEl(null)
                    setTimeout(() => {
                        setOrderItems([])
                    }, 300)
                }}
                anchorOrigin={{
                    vertical: "bottom",
                    horizontal: "center"
                }}
            >
                <RenderList
                    list={orderItems}
                    render={orderItem => {
                        return <OrderItem data={orderItem} />
                    }}
                />
            </Popover>
        </>
    )
}

const OrderItem = ({ data }: { data: OrderItems }) => {
    const color = useGetStatusColor(data.shipping_status)
    return (
        <Box
            key={data.id}
            sx={{ borderBottom: "1px solid #c9c9c9", padding: 2 }}
        >
            <Typography variant='body1'>{data.name}</Typography>
            <Chip
                label={data.shipping_status}
                color={color}
                size='small'
                variant='outlined'
            />
            <Stack direction='row' mt={2}>
                <Typography variant='body2' color='textSecondary'>
                    Price:&nbsp;
                    <Currency amount={data.price} />
                </Typography>
                <Typography variant='body2' color='textSecondary'>
                    Quantity:&nbsp;{data.quantity}
                </Typography>
            </Stack>
        </Box>
    )
}

const PurchaseItem = ({
    products,
    onClickMoreItem
}: {
    products: OrderItems[]
    onClickMoreItem: (
        target: React.MouseEvent<HTMLSpanElement, MouseEvent>
    ) => void
}) => {
    const [name, moreCount] = useMemo(() => {
        const [firstProduct, ...rest] = products
        return [firstProduct.name, rest.length]
    }, [products])

    return (
        <Typography>
            {name}
            <Typography
                variant='caption'
                color='textSecondary'
                sx={{ cursor: "pointer" }}
                onClick={onClickMoreItem}
            >
                {moreCount > 0 ? `(+${moreCount}more)` : null}
            </Typography>
        </Typography>
    )
}

export default Orders
