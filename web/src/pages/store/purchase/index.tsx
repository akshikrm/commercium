import HeaderBreadcrumbs from "@components/header"
import RHFProvider from "@components/rhf/provider"
import RHFTextField from "@components/rhf/text-field"
import { useForm } from "react-hook-form"

const Purchase = () => {
    const methods = useForm()
    const onSubmit = async (inputData: any) => {
        console.log(inputData)
    }
    return (
        <>
            <HeaderBreadcrumbs
                heading='Purchase'
                links={[
                    { label: "Home", href: "/" },
                    { label: "Purchase", href: "/" }
                ]}
            />
            <RHFProvider methods={methods} onSubmit={onSubmit}>
                <RHFTextField name='test' label='Test' />
            </RHFProvider>
        </>
    )
}

export default Purchase
