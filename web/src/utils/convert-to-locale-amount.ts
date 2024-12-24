const parseToLocaleAmount = (amount: string): string => {
    const parsed = parseFloat(amount)
    if (!isNaN(parsed)) {
        const formatter = new Intl.NumberFormat("en-IN", {
            style: "decimal",
            minimumFractionDigits: 2,
            maximumFractionDigits: 2
        })
        return formatter.format(parsed / 100)
    }

    return ""
}

export default parseToLocaleAmount
