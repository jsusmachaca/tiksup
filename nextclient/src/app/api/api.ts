import axios from 'axios';

// Configura Axios para incluir el token en las solicitudes
const api = axios.create({
    baseURL: 'http://localhost:3000',
});

api.interceptors.request.use(
    config => {
        const token = localStorage.getItem('token'); // O donde sea que almacenes el token
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

export default api;
