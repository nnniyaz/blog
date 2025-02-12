import { getBioDb } from '@lib/mongo/mongo'
import { NewError, NewInternal, NewNotFound } from '@lib/web/error'
import { NewSuccess } from '@lib/web/success'
import { NextRequest } from 'next/server'

export async function GET(req: NextRequest) {
    try {
        const { collection, error } = await getBioDb()
        if (error instanceof Error) {
            return NewInternal(req, error)
        }

        const bio = collection!.findOne({ active: true })
        if (!bio) {
            return NewNotFound(req, new Error('Not found'))
        }
        return NewSuccess(req, bio)
    } catch (error: unknown) {
        console.error(error)
        if (typeof error === 'string') return NewError(req, new Error(error))
        else if (error instanceof Error) return NewError(req, error)
        else return NewError(req, new Error('Unknown error'))
    }
}
