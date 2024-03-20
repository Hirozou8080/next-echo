"use client";
import { TextField, Typography } from "@mui/material";
export default function Page() {
  const response = fetch("http://127.0.0.2:8080", {
    method: "GET",
    cache: "no-cache",
    mode: "cors",
  }).then((res) => res.json());
  console.log(response);
  console.log("Hello, world!");
  return (
    <div>
      <h1>Login</h1>
      <Typography>ID</Typography>
      <TextField name="id" />
      <Typography>PASS</Typography>
      <TextField name="pass" />
    </div>
  );
}
