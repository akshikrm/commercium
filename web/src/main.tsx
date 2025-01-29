import { StrictMode } from "react"
import { createRoot } from "react-dom/client"
import { BrowserRouter, Navigate, Route, Routes } from "react-router"
import Login from "./pages/login"
import { Container, ThemeProvider } from "@mui/material"
import theme from "./theme"
import AddProduct from "./pages/products/add"
import EditProduct from "./pages/products/edit"
import { AuthGuard, AdminGuard, UserGuard } from "@components/guards.tsx"
import AddProductCategory from "./pages/products/categories/add/index.tsx"
import EditProductCategory from "./pages/products/categories/edit/index.tsx"
import { LocalizationProvider } from "@mui/x-date-pickers"
import { AdapterDayjs } from "@mui/x-date-pickers/AdapterDayjs"
import { QueryClient, QueryClientProvider } from "@tanstack/react-query"
import Products from "./pages/products/index.tsx"
import { Toaster } from "react-hot-toast"
import ProductCategoryList from "./pages/products/categories/index.tsx"
import "./index.css"
import Layout from "@components/layout.tsx"
import { AuthenticationProvider } from "./store/index.ts"
import Store from "./pages/store/index.tsx"
import Orders from "./pages/store/orders/index.tsx"
import Cart from "./pages/cart/index.tsx"
import { USER_PATHS } from "./paths.ts"
import OrderView from "./pages/store/orders/view/index.tsx"
import Shipping from "./pages/shipping/index.tsx"

const client = new QueryClient()

createRoot(document.getElementById("root")!).render(
    <StrictMode>
        <ThemeProvider theme={theme}>
            <Toaster
                position='top-right'
                toastOptions={{
                    style: {
                        fontFamily: "-apple-system",
                        fontSize: 15
                    }
                }}
            />

            <QueryClientProvider client={client}>
                <AuthenticationProvider>
                    <LocalizationProvider dateAdapter={AdapterDayjs}>
                        <BrowserRouter>
                            <Routes>
                                <Route
                                    path='auth'
                                    element={
                                        <Container maxWidth='sm'>
                                            <AuthGuard />
                                        </Container>
                                    }
                                >
                                    <Route path='login' element={<Login />} />
                                    <Route
                                        path='register'
                                        element={<h1>Register</h1>}
                                    />
                                </Route>

                                <Route
                                    path='/'
                                    element={
                                        <>
                                            <Layout>
                                                <UserGuard />
                                            </Layout>
                                        </>
                                    }
                                >
                                    <Route
                                        index
                                        element={
                                            <Navigate
                                                to={USER_PATHS.store.root}
                                            />
                                        }
                                    />
                                    <Route path='stores' element={<Store />} />
                                    <Route path='carts' element={<Cart />} />
                                    <Route path='orders'>
                                        <Route
                                            index={true}
                                            element={<Orders />}
                                        />
                                        <Route
                                            path=':id'
                                            element={<OrderView />}
                                        />
                                    </Route>
                                </Route>
                                <Route
                                    path='admin'
                                    element={
                                        <>
                                            <Layout>
                                                <AdminGuard />
                                            </Layout>
                                        </>
                                    }
                                >
                                    <Route path='orders' element={<Orders />} />
                                    <Route
                                        path='shipping'
                                        element={<Shipping />}
                                    />
                                    <Route path='products'>
                                        <Route index element={<Products />} />
                                        <Route
                                            path='add'
                                            element={<AddProduct />}
                                        />
                                        <Route
                                            path=':product_id'
                                            element={<EditProduct />}
                                        />
                                        <Route path='categories'>
                                            <Route
                                                index
                                                element={
                                                    <ProductCategoryList />
                                                }
                                            />
                                            <Route
                                                path='add'
                                                element={<AddProductCategory />}
                                            />
                                            <Route
                                                path=':category_id'
                                                element={
                                                    <EditProductCategory />
                                                }
                                            />
                                        </Route>
                                    </Route>
                                </Route>
                            </Routes>
                        </BrowserRouter>
                    </LocalizationProvider>
                </AuthenticationProvider>
            </QueryClientProvider>
        </ThemeProvider>
    </StrictMode>
)
