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
import { useNavigate } from "react-router"
import { USER_PATHS } from "@/paths"

const Orders = () => {
    const { data: orders } = useGetOrders()
    const navigate = useNavigate()
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
                            <TableCell>SI.No</TableCell>
                            <TableCell>Order ID</TableCell>
                            <TableCell>Items</TableCell>
                            <TableCell>Price</TableCell>
                            <TableCell>Purchased On</TableCell>
                            <TableCell>View</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <RenderList
                            list={orders}
                            render={(row, i: number) => {
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
                                        <TableCell>{i + 1}</TableCell>
                                        <TableCell>#{row.order_id}</TableCell>
                                        <TableCell>
                                            {row.products.length}
                                        </TableCell>
                                        <TableCell>
                                            <Currency>
                                                {row.purchase_price}
                                            </Currency>
                                        </TableCell>
                                        <TableCell>
                                            {dayjs(row.created_at).format(
                                                DATE_VIEW_FORMAT
                                            )}
                                        </TableCell>

                                        <TableCell>
                                            <IconButton
                                                onClick={() => {
                                                    navigate(
                                                        USER_PATHS.orders.view(
                                                            row.order_id
                                                        )
                                                    )
                                                }}
                                            >
                                                <RenderIcon icon={icons.view} />
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

export default Orders
