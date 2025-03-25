import { AxiosResponse } from 'axios'

import { User } from '@domain/user/user.ts'
import { $api, ApiRoutes, Request } from '@http/index.ts'
import { ErrorResponse, SuccessResponse } from '@http/response/response.ts'

type CurrentUserServiceUpdateEmailIn = {
    id: string
    email: string
}

type CurrentUserServiceUpdatePasswordIn = {
    id: string
    password: string
}

export class CurrentUserService {
    static async getCurrentUser(
        request: Request<null>,
    ): Promise<AxiosResponse<SuccessResponse<User> | ErrorResponse>> {
        return $api(request.lang).get<SuccessResponse<User> | ErrorResponse>(
            ApiRoutes.ME_GET,
            { signal: request.controller?.signal },
        )
    }

    static async updateEmail(
        request: Request<CurrentUserServiceUpdateEmailIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ME_UPDATE_EMAIL,
            null,
            { signal: request.controller?.signal },
        )
    }

    static async updatePassword(
        request: Request<CurrentUserServiceUpdatePasswordIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.ME_UPDATE_PASSWORD,
            null,
            { signal: request.controller?.signal },
        )
    }
}
