import http from "k6/http";

export let options = {
    vus: 10,             // Virtual users
    duration: "10s",     // Duration
};

export default function() {
    http.get("http://host.docker.internal:9000");
};