import http from "k6/http";

export let options = {
    vus: 5,             // Virtual users
    duration: "5s",     // Duration
};

export default function() {
    http.get("http://host.docker.internal:9000/products");
};