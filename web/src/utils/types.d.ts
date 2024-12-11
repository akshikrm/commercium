type FailedResponse = {
  status: boolean;
  message: string;
  data?: ValidationErrors | null;
};

type Role = "admin" | "user";
type JWTPayload = {
  role: Role;
  sub: number;
};
