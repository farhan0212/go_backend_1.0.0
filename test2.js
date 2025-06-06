import http from "k6/http";
import { check } from "k6";

export let options = {
  vus: 100,
  duration: "30s",
};

export default function () {
  const res = http.get("http://localhost:8080/api/users", {
    headers: {
      Authorization: "Bearer <TOKEN_KAMU>",
    },
  });
  check(res, {
    "status was 200": (r) => r.status === 200,
  });
}
