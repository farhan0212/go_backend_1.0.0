import http from "k6/http";
import { check } from "k6";
import { uuidv4 } from "https://jslib.k6.io/k6-utils/1.4.0/index.js";

export let options = {
  vus: 50, // virtual users
  duration: "30s",
};

export default function () {
  // Generate unique name and email per request
  const name = `user-${uuidv4()}`;
  const email = `${name}@example.com`;

  const payload = JSON.stringify({
    name: name,
    email: email,
  });

  const headers = {
    "Content-Type": "application/json",
  };

  const res = http.post("http://localhost:8080/users", payload, { headers });

  check(res, {
    "status is 201 or 200": (r) => r.status === 201 || r.status === 200,
  });
}
