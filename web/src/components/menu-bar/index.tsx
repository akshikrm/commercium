import { useTheme } from "@mui/material"
import { Menu, Sidebar } from "react-pro-sidebar"
import RenderList from "@components/render-list"
import Item from "./components/item"
import menus from "./paths"

const MenuBar = () => {
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
                <RenderList
                    list={menus}
                    render={menu => {
                        return <Item {...menu} />
                    }}
                />
            </Menu>
        </Sidebar>
    )
}

export default MenuBar
