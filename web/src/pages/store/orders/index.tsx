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
import { useMemo } from "react"
import { Typography } from "@mui/material"
import parseToLocaleAmount from "@utils/convert-to-locale-amount"
import { order } from "@api"

const Orders = () => {
    const { data: orders } = useGetOrders()
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
                            <TableCell>Download</TableCell>
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
                                            <IconButton
                                                onClick={() =>
                                                    order.gerOrderByID(
                                                        row.transaction_id
                                                    )
                                                }
                                            >
                                                <RenderIcon
                                                    icon={icons.download}
                                                />
                                            </IconButton>
                                        </TableCell>
                                    </TableRow>
                                )
                            }}
                        />
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    )
}

const PurchaseItem = ({ products }: { products: OrderItems[] }) => {
    const [name, moreCount] = useMemo(() => {
        const [firstProduct, ...rest] = products
        return [firstProduct.name, rest.length]
    }, [products])

    return (
        <Typography>
            {name}
            <Typography variant='caption'>
                {moreCount > 0 ? `(+${moreCount}more)` : null}
            </Typography>
        </Typography>
    )
}

export default Orders
