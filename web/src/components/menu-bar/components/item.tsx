import { MenuItem, SubMenu } from "react-pro-sidebar"
import { Link, useLocation } from "react-router"
import RenderList from "@components/render-list"
import useHandleOpenStatus from "../hooks/use-handle-open-status"
import ItemLabel from "./item-label"

const Item = (props: Paths) => {
    const { label, path, children } = props
    const { pathname } = useLocation()
    const [open, toggleOpen] = useHandleOpenStatus(children)

    if (children?.length) {
        return (
            <SubMenu
                label={<ItemLabel label={label} />}
                open={open}
                onClick={toggleOpen}
            >
                <RenderList
                    list={children}
                    render={item => <Item {...item} />}
                />
            </SubMenu>
        )
    }

    return (
        <MenuItem active={pathname === path} component={<Link to={path} />}>
            <ItemLabel label={label} />
        </MenuItem>
    )
}

export default Item
