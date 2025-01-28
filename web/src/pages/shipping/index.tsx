import { order } from "@api"
import HeaderBreadcrumbs from "@components/header"
import { useQuery } from "@tanstack/react-query"

import Table from "@mui/material/Table"
import TableBody from "@mui/material/TableBody"
import TableCell from "@mui/material/TableCell"
import TableContainer from "@mui/material/TableContainer"
import TableHead from "@mui/material/TableHead"
import TableRow from "@mui/material/TableRow"

import RenderList from "@components/render-list"
import Paper from "@mui/material/Paper"
import { DATE_VIEW_FORMAT } from "@config"
import dayjs from "dayjs"
import { Currency } from "@components/prefix"
import { Button, ButtonGroup, Menu, MenuItem } from "@mui/material"
import RenderIcon from "@components/render-icon"
import icons from "@/icons"
import { useMemo, useState } from "react"

const Shipping = () => {
    const query = useQuery({
        initialData: [],
        queryKey: ["shippingInformation"],
        queryFn: async () => await order.getShippingInformation()
    })

    return (
        <>
            <HeaderBreadcrumbs
                heading='Shipping'
                links={[{ label: "Home", href: "/" }, { label: "Shipping" }]}
            />
            <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label='simple table'>
                    <TableHead>
                        <TableRow>
                            <TableCell>No</TableCell>
                            <TableCell>Username</TableCell>
                            <TableCell>Shipping Status</TableCell>
                            <TableCell>Product</TableCell>
                            <TableCell>Price</TableCell>
                            <TableCell>Purchased On</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <RenderList
                            list={query.data}
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
                                        <TableCell>1</TableCell>
                                        <TableCell>{row.user.name}</TableCell>
                                        <TableCell>
                                            <StatusButton status={row.status} />
                                        </TableCell>
                                        <TableCell>
                                            {row.product.name}
                                        </TableCell>
                                        <TableCell>
                                            <Currency amount={row.amount} />
                                        </TableCell>
                                        <TableCell>
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
            </TableContainer>{" "}
        </>
    )
}

const StatusButton = ({ status }: { status: ShippingStatus }) => {
    const [anchorEl, setAnchor] = useState<Element | null>(null)
    const handleClose = () => {
        setAnchor(null)
    }

    const color = useMemo(() => {
        switch (status) {
            case "delivered": {
                return "success"
            }
            case "in-transit": {
                return "primary"
            }
            case "pending": {
                return "warning"
            }
        }
    }, [status])

    console.log(color)

    return (
        <>
            <ButtonGroup>
                <Button variant='outlined' size='small' color={color}>
                    {status}
                </Button>
                <Button
                    variant='outlined'
                    size='small'
                    color={color}
                    onClick={e => {
                        setAnchor(e.currentTarget)
                    }}
                >
                    <RenderIcon icon={icons.arrowDown} />
                </Button>
            </ButtonGroup>
            <Menu
                id='basic-menu'
                anchorEl={anchorEl}
                open={Boolean(anchorEl)}
                onClose={handleClose}
                MenuListProps={{
                    "aria-labelledby": "basic-button"
                }}
            >
                <MenuItem
                    onClick={handleClose}
                    selected={status === "delivered"}
                >
                    Delivered
                </MenuItem>
                <MenuItem
                    onClick={handleClose}
                    selected={status === "in-transit"}
                >
                    In Transit
                </MenuItem>
                <MenuItem onClick={handleClose} selected={status === "pending"}>
                    Pending
                </MenuItem>
            </Menu>
        </>
    )
}

export default Shipping
