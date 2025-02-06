import { MlString } from '../base/mlString/mlString'

export type Article = {
    id: string
    title: MlString
    content: MlString
    isDeleted: boolean
    createdAt: string
    updatedAt: string
    version: number
}
