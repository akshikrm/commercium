import { Typography, useTheme } from "@mui/material"
import { Menu, MenuItem, Sidebar, SubMenu } from "react-pro-sidebar"
import { Link, useLocation } from "react-router"

const MenuBar = () => {
    const { pathname } = useLocation()
    const theme = useTheme()

    return (
        <Sidebar>
            <Menu
                menuItemStyles={{
                    button: test => {
                        const { active } = test
                        return {
                            color: active
                                ? theme.palette.primary.contrastText
                                : "#d359ff",
                            backgroundColor: active
                                ? theme.palette.primary.main
                                : undefined,

                            "&:hover": {
                                backgroundColor: theme.palette.primary.light
                            }
                        }
                    }
                }}
            >
                <MenuItem>
                    <MenuItemLabel label='Dashboard' />
                </MenuItem>
                <SubMenu label={<MenuItemLabel label='e-commerce' />}>
                    <MenuItem
                        active={USER_PATHS.store.root == pathname}
                        component={<Link to={USER_PATHS.store.root} />}
                    >
                        <MenuItemLabel label='product' />
                    </MenuItem>
                    <MenuItem
                        active={USER_PATHS.orders.root == pathname}
                        component={<Link to={USER_PATHS.orders.root} />}
                    >
                        <MenuItemLabel label='orders' />
                    </MenuItem>
                </SubMenu>
                <MenuItem>
                    <MenuItemLabel label='Reports' />
                </MenuItem>
            </Menu>
        </Sidebar>
    )
}

const MenuItemLabel = ({ label }: { label: string }) => {
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


export defult MenuBar
