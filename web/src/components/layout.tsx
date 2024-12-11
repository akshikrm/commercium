import { ReactNode } from "react";

import {
  AppBar,
  Box,
  Button,
  Container,
  Toolbar,
  Typography,
} from "@mui/material";
import useLogout from "@hooks/auth/use-logout";
import useAuth from "@hooks/auth/use-auth";

const Layout = ({ children }: { children: ReactNode }) => {
  const logout = useLogout();
  const { user } = useAuth();

  const { first_name, last_name } = user;
  return (
    <>
      <AppBar position="static">
        <Container>
          <Toolbar variant="dense">
            <Typography variant="h6" color="inherit" component="div">
              Welcome back, {first_name} {last_name}
            </Typography>
            <Box
              sx={{
                flexGrow: 1,
                justifyContent: "flex-end",
                display: "flex",
              }}
            >
              <Button aria-label="menu" onClick={logout}>
                logout
              </Button>
            </Box>
          </Toolbar>
        </Container>
      </AppBar>
      <Container>{children}</Container>
    </>
  );
};

export default Layout;
