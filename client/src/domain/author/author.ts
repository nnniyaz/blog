import { MlString } from '@domain/base/ml-string'

export type Author = {
    _id: string
    firstName: MlString
    lastName: MlString
    avatarUri: string
    createdAt: string
    updatedAt: string
    version: number
}
