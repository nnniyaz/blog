import { MlString } from '@domain/base/ml-string'

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
