const DENOMINATION = 100

export const convertToCommonAmount = (amount: number) =>
    (amount / DENOMINATION).toFixed(2)

export const Currency = ({ amount }: { amount: number }) => {
    return <>&#8377;{convertToCommonAmount(amount) || 0}</>
}
