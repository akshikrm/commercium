import { useMemo } from "react"

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

export default useGetStatusColor
