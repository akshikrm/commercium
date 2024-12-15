import HeaderBreadcrumbs from "@components/header";

const Purchase = () => {
	return (
		<HeaderBreadcrumbs
			heading="Purchase"
			links={[
				{ label: "Home", href: "/" },
				{ label: "Purchase", href: "/" },
			]}
		/>
	);
};

export default Purchase;
