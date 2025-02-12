import { MlString } from '@domain/base/ml-string'

export type Contact = {
    _id: string
    label: MlString
    link: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}

export type ContactsMany = {
    contacts: Contact[]
    count: number
}
