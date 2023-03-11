import http from "k6/http";
import { sleep, check } from "k6";

export let options = {
    vus: 10, // virtual users
    duration: '30s',
    rps: 100 // requests per second
};

export default function () {
    let data = [{source: 'GSO', destination: 'IND'}, {source: 'ATL', destination: 'GSO'}];

    let res = http.post("http://localhost:8080/track", JSON.stringify(data), {
        headers: { 'Content-Type': 'application/json' },
    });
    check(res, {
        'response code was 200': (res) => res.status === 200,
    });

    sleep(1);
}