import { MlString } from '../base/mlString/mlString'

export interface Book {
    id: string
    title: MlString
    author: MlString
    description: MlString
    coverUri: string
    eBookUri: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}
