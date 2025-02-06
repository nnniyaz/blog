import { MlString } from '../base/mlString/mlString'

export type Author = {
    id: string
    firstName: MlString
    lastName: MlString
    avatarUri: string
    createdAt: string
    updatedAt: string
    version: number
}
