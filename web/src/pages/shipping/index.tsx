import { order } from "@api"
import HeaderBreadcrumbs from "@components/header"
import { useMutation, useQuery, useQueryClient } from "@tanstack/react-query"
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
import toast from "react-hot-toast"

const Shipping = () => {
    const queryClient = useQueryClient()
    const query = useQuery({
        initialData: [],
        queryKey: ["shippingInformationList"],
        queryFn: async () => await order.getShippingInformation()
    })

    console.log(query.data)
    const mutation = useMutation({
        mutationFn: async (payload: {
            orderID: number
            status: ShippingStatus
        }) => {
            const { orderID, status } = payload
            return await order.updateShippingStatus(orderID, status)
        },
        onSuccess: async (data, vars) => {
            toast.success(data)
            queryClient.setQueryData(
                ["shippingInformationList"],
                (prevData: ShippingInformation[]) => {
                    const temp = [...prevData]
                    const itemIndex = temp.findIndex(
                        ({ id }) => id === vars.orderID
                    )
                    if (itemIndex > -1) {
                        const newElement = {
                            ...temp[itemIndex],
                            status: vars.status
                        }
                        temp.splice(itemIndex, 1, newElement)
                    }

                    return temp
                }
            )
        }
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
                            <TableCell>Product</TableCell>
                            <TableCell>Quantity</TableCell>
                            <TableCell>Price</TableCell>
                            <TableCell>Shipping Status</TableCell>
                            <TableCell>Purchased On</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        <RenderList
                            list={query.data}
                            render={(row, i) => {
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
                                        <TableCell>
                                            {row.user.first_name}&nbsp;
                                            {row.user.last_name}
                                        </TableCell>
                                        <TableCell>
                                            {row.product.name}
                                        </TableCell>
                                        <TableCell>{row.quantity}</TableCell>
                                        <TableCell>
                                            <Currency amount={row.amount} />
                                        </TableCell>
                                        <TableCell>
                                            <StatusButton
                                                status={row.status}
                                                handleUpdate={status => {
                                                    mutation.mutate({
                                                        orderID: row.id,
                                                        status
                                                    })
                                                }}
                                            />
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
            </TableContainer>
        </>
    )
}

const useGetStatusColor = (status: ShippingStatus) => {
    return useMemo(() => {
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
}

const StatusButton = ({
    status,
    handleUpdate
}: {
    status: ShippingStatus
    handleUpdate: (status: ShippingStatus) => void
}) => {
    const [anchorEl, setAnchor] = useState<Element | null>(null)

    const handleChange = (newStatus?: ShippingStatus) => {
        if (newStatus) {
            if (newStatus != status) {
                handleUpdate(newStatus)
            }
        }
        setAnchor(null)
    }
    const color = useGetStatusColor(status)

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
                onClose={() => handleChange()}
                MenuListProps={{
                    "aria-labelledby": "basic-button"
                }}
            >
                <MenuItem
                    onClick={() => handleChange("delivered")}
                    selected={status === "delivered"}
                >
                    Delivered
                </MenuItem>
                <MenuItem
                    onClick={() => handleChange("in-transit")}
                    selected={status === "in-transit"}
                >
                    In Transit
                </MenuItem>
                <MenuItem
                    onClick={() => handleChange("pending")}
                    selected={status === "pending"}
                >
                    Pending
                </MenuItem>
            </Menu>
        </>
    )
}

export default Shipping
