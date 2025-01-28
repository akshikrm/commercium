const DENOMINATION = 100

const showCommonAmount = (amount: number) => (amount / DENOMINATION).toFixed(2)

export const Currency = ({ amount }: { amount: number }) => {
    return <>&#8377;{showCommonAmount(amount) || 0}</>
}
