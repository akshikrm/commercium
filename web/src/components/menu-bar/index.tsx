import { useTheme } from "@mui/material"
import { Menu, Sidebar } from "react-pro-sidebar"
import RenderList from "@components/render-list"
import Item from "./components/item"
import { ADMIN_MENUS, USER_MENUS } from "./paths"
import useIsUser from "@hooks/auth/use-is-user"

const MenuBar = () => {
    const theme = useTheme()

    const isUser = useIsUser()

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
                <RenderList
                    list={isUser ? USER_MENUS : ADMIN_MENUS}
                    render={menu => {
                        return <Item {...menu} />
                    }}
                />
            </Menu>
        </Sidebar>
    )
}

export default MenuBar
