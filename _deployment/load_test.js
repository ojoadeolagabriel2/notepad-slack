import{ sleep } from 'k6';
import http from 'k6/http';

export let options = {
    duration : '10s',
    vus : 50,
};

export default function() {
    http.get('http://localhost:12345/liveliness');
}