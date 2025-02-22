import { MlString } from '@domain/base/ml-string'

export type Contact = {
    id: string
    label: MlString
    link: string
    isDeleted: boolean
    createdAt: string
    updatedAt: string
}
