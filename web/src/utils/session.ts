export const setSession = (token: string): Role => {
    const { role } = decodeJWT(token)
    localStorage.setItem("auth-token", token)
    localStorage.setItem("role", role)
    return role
}

export const clearSession = (): void => {
    localStorage.removeItem("auth-token")
    localStorage.removeItem("role")
}

export const getToken = (): string => {
    const authToken = localStorage.getItem("auth-token")

    if (authToken) {
        return authToken
    }
    return ""
}

export const decodeJWT = (token: string): JWTPayload => {
    const payload = token.split(".")[1].replace(/-/g, "+").replace(/-/g, "/")
    return JSON.parse(
        decodeURIComponent(
            window
                .atob(payload)
                .split("")
                .map(c => {
                    return "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2)
                })
                .join("")
        )
    )
}
