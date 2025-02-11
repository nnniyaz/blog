import { MlString } from '@domain/base/ml-string'

export type Bio = {
    _id: string
    bio: MlString
    active: boolean
    createdAt: string
    updatedAt: string
    version: number
}
