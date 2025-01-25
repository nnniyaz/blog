import {AxiosResponse} from 'axios'

import {MlString} from '../../domain/base/mlString/mlString'
import {Contact} from '../../domain/contact/contact'
import {$api, ApiRoutes, Request} from '../index'
import {ErrorResponse, SuccessResponse} from '../response/response'

type ContactServiceGetAllOut = {
    count: number
    contacts: Contact[]
}

type ContactServiceGetByIdIn = {
    id: string
}

type ContactServiceCreateIn = {
    label: MlString
    link: string
}

type ContactServiceUpdateIn = {
    id: string
    label: MlString
    link: string
}

type ContactServiceDeleteIn = {
    id: string
}

type ContactServiceRestoreIn = {
    id: string
}

export class ContactService {
    static async getAll(request: Request<null>): Promise<AxiosResponse<SuccessResponse<ContactServiceGetAllOut> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<ContactServiceGetAllOut> | ErrorResponse>(ApiRoutes.CONTACT_GET_ALL, null, {signal: request.controller?.signal})
    }

    static async getById(request: Request<ContactServiceGetByIdIn>): Promise<AxiosResponse<SuccessResponse<Contact> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<Contact> | ErrorResponse>(ApiRoutes.CONTACT_GET_BY_ID.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async create(request: Request<ContactServiceCreateIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.CONTACT_CREATE, null, {signal: request.controller?.signal})
    }

    static async update(request: Request<ContactServiceUpdateIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.CONTACT_UPDATE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async delete(request: Request<ContactServiceDeleteIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.CONTACT_DELETE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }

    static async restore(request: Request<ContactServiceRestoreIn>): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(ApiRoutes.CONTACT_RESTORE.replace(':id', request.body.id), null, {signal: request.controller?.signal})
    }
}