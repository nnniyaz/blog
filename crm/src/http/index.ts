import axios from 'axios'

import { Lang } from '@domain/base/ml-string'

export enum ApiRoutes {
    // Auth routes
    AUTH_LOGIN = '/auth/login',
    AUTH_LOGOUT = '/auth/logout',

    // Current user routes
    ME_GET = '/me',
    ME_UPDATE_EMAIL = '/me/email',
    ME_UPDATE_PASSWORD = '/me/password',

    // Author routes
    AUTHOR_GET_ALL = '/about/author',
    AUTHOR_GET_BY_ID = '/about/author/:id',
    AUTHOR_CREATE = '/about/author',
    AUTHOR_UPDATE = '/about/author/:id',
    AUTHOR_DELETE = '/about/author/:id',
    AUTHOR_RESTORE = '/about/author/restore/:id',

    // Bio routes
    BIO_GET_ALL = '/about/bio',
    BIO_GET_ACTIVE = '/about/bio/active',
    BIO_GET_BY_ID = '/about/bio/:id',
    BIO_CREATE = '/about/bio',
    BIO_UPDATE = '/about/bio/:id',
    BIO_DELETE = '/about/bio/:id',
    BIO_RESTORE = '/about/bio/restore/:id',

    // Contact routes
    CONTACT_GET_ALL = '/about/contact',
    CONTACT_GET_BY_ID = '/about/contact/:id',
    CONTACT_CREATE = '/about/contact',
    CONTACT_UPDATE = '/about/contact/:id',
    CONTACT_DELETE = '/about/contact/:id',
    CONTACT_RESTORE = '/about/contact/restore/:id',

    // Project routes
    PROJECT_GET_ALL = '/project',
    PROJECT_GET_BY_ID = '/project/:id',
    PROJECT_CREATE = '/project',
    PROJECT_UPDATE = '/project/:id',
    PROJECT_DELETE = '/project/:id',
    PROJECT_RESTORE = '/project/restore/:id',

    // Article routes
    ARTICLE_GET_ALL = '/article',
    ARTICLE_GET_BY_ID = '/article/:id',
    ARTICLE_CREATE = '/article',
    ARTICLE_UPDATE = '/article/:id',
    ARTICLE_DELETE = '/article/:id',
    ARTICLE_RESTORE = '/article/restore/:id',

    // Book routes
    BOOK_GET_ALL = '/book',
    BOOK_GET_BY_ID = '/book/:id',
    BOOK_CREATE = '/book',
    BOOK_UPDATE = '/book/:id',
    BOOK_DELETE = '/book/:id',
    BOOK_RESTORE = '/book/restore/:id',

    // User routes
    USER_GET_ALL = '/user',
    USER_GET_BY_ID = '/user/:id',
    USER_CREATE = '/user',
    USER_UPDATE_EMAIL = '/user/email/:id',
    USER_UPDATE_PASSWORD = '/user/password/:id',
    USER_UPDATE_ROLE = '/user/role/:id',
    USER_DELETE = '/user/:id',
    USER_RESTORE = '/user/restore/:id',
}

const validateStatus = (status: number) => {
    return (
        (status >= 200 && status < 401) ||
        (status > 401 && status !== 403 && status !== 404 && status < 500)
    )
}

export type Request<T> = {
    lang: Lang
    body: T
    controller?: AbortController
}

export const $api = (lang: Lang) =>
    axios.create({
        baseURL: import.meta.env.VITE_API_URI || '',
        withCredentials: true,
        headers: {
            'Content-Type': 'application/json',
            'Accept-Language': lang.toLowerCase(),
        },
        timeout: 120000,
        timeoutErrorMessage: 'Timeout error',
        validateStatus: validateStatus,
    })

export const $apiFormData = (lang: Lang) =>
    axios.create({
        baseURL: import.meta.env.VITE_API_URI || '',
        withCredentials: true,
        headers: {
            'Content-Type': 'multipart/form-data',
            'Accept-Language': lang.toLowerCase(),
        },
        timeout: 120000,
        timeoutErrorMessage: 'Timeout error',
        validateStatus: validateStatus,
    })
