import HeaderBreadcrumbs from "@components/header"

const Shipping = () => {
    return (
        <>
            <HeaderBreadcrumbs
                heading='Shipping'
                links={[{ label: "Home", href: "/" }, { label: "Shipping" }]}
            />
        </>
    )
}

export default Shipping
