import HeaderBreadcrumbs from "@components/header"

const Subscriptions = () => {
    return (
        <>
            <HeaderBreadcrumbs
                heading='Subscription'
                links={[
                    { label: "home", href: "/" },
                    { label: "subscriptions" }
                ]}
            />
        </>
    )
}

export default Subscriptions
