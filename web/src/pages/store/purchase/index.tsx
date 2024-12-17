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
import useGetPurchases from "@hooks/purchase/use-get-purchases"

const Purchase = () => {
    const { data: purchases } = useGetPurchases()
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
                            <TableCell>Name</TableCell>
                            <TableCell>Price</TableCell>
                            <TableCell>Purchased On</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <RenderList
                            list={purchases}
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
                                        <TableCell component='th' scope='row'>
                                            {i + 1}
                                        </TableCell>
                                        <TableCell component='th' scope='row'>
                                            #{row.order_id}
                                        </TableCell>
                                        <TableCell component='th' scope='row'>
                                            {row.product.name}
                                        </TableCell>
                                        <TableCell component='th' scope='row'>
                                            <Currency>
                                                {row.purchase_price}
                                            </Currency>
                                        </TableCell>
                                        <TableCell component='th' scope='row'>
                                            {dayjs(row.created_at).format(
                                                DATE_VIEW_FORMAT
                                            )}
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

export default Purchase
