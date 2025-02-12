import { getContactDb } from '@lib/mongo/mongo'
import { NewError, NewInternal } from '@lib/web/error'
import { NewSuccess } from '@lib/web/success'
import { NextRequest } from 'next/server'

export async function GET(req: NextRequest) {
    try {
        const { collection, error } = await getContactDb()
        if (error instanceof Error) {
            return NewInternal(req, error)
        }

        const cur = collection!.find()
        const contacts = await cur.toArray()
        return NewSuccess(req, contacts)
    } catch (error: unknown) {
        console.error(error)
        if (typeof error === 'string') return NewError(req, new Error(error))
        else if (error instanceof Error) return NewError(req, error)
        else return NewError(req, new Error('Unknown error'))
    }
}
