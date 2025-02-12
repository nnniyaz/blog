import axios from 'axios'

export enum ApiRoutes {
    ARTICLE = '/api/article',
    AUTHOR = '/api/author',
    BIO = '/api/bio',
    BOOK = '/api/book',
    CONTACT = '/api/contact',
    PROJECT = '/api/project',
}

const validateStatus = (status: number) => {
    return (
        (status >= 200 && status < 401) ||
        (status > 401 && status !== 403 && status !== 404 && status < 500)
    )
}

export const $api = axios.create({
    baseURL: process.env.CLIENT_URI || '',
    withCredentials: true,
    headers: {
        'Content-Type': 'application/json',
    },
    timeout: 120000,
    timeoutErrorMessage: 'Timeout error',
    validateStatus: validateStatus,
})
