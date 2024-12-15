import { zodResolver } from "@hookform/resolvers/zod"
import { useEffect } from "react"
import { useForm } from "react-hook-form"
import { z } from "zod"

const loginSchema = z.object({
    email: z.string().min(1, { message: "email is required" }).email({
        message: "please check your email"
    }),
    password: z.string().min(1, { message: "password is required" })
})

const loginDefaultValues: LoginRequest = {
    email: "",
    password: ""
}

const useLoginForm = (error: ValidationErrors) => {
    const methods = useForm<LoginRequest>({
        defaultValues: loginDefaultValues,
        resolver: zodResolver(loginSchema)
    })

    const { setError } = methods
    useEffect(() => {
        if (error) {
            Object.entries(error).forEach(([k, v]) => {
                setError(k, {
                    message: v
                })
            })
        }
    }, [error, setError])

    return methods
}

export default useLoginForm
