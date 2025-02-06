import { AxiosResponse } from 'axios'

import { $api, ApiRoutes, Request } from '../index'
import { ErrorResponse, SuccessResponse } from '../response/response'

type AuthServiceLoginIn = {
    email: string
    password: string
}

export class AuthService {
    static async login(
        request: Request<AuthServiceLoginIn>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.AUTH_LOGIN,
            request.body,
        )
    }

    static async logout(
        request: Request<null>,
    ): Promise<AxiosResponse<SuccessResponse<null> | ErrorResponse>> {
        return $api(request.lang).post<SuccessResponse<null> | ErrorResponse>(
            ApiRoutes.AUTH_LOGOUT,
            null,
        )
    }
}
