import { MlString } from '@domain/base/ml-string'

export type Book = {
    _id: string
    title: MlString
    author: MlString
    description: MlString
    coverUri: string
    eBookUri: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}

export type BooksMany = {
    books: Book[]
    count: number
}
