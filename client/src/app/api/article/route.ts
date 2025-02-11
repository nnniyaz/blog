import { Lang } from '@domain/base/ml-string'
import { getArticleDb } from '@lib/mongo/mongo'
import { parseIntQueryParam } from '@lib/utils/parse-query-params'
import { NewBad, NewError, NewInternal } from '@lib/web/error'
import { NewSuccess } from '@lib/web/success'
import { NextRequest } from 'next/server'

export async function GET(req: NextRequest) {
    try {
        const searchParams = req.nextUrl.searchParams
        const offset = parseIntQueryParam(searchParams.get('offset'))
        const limit = parseIntQueryParam(searchParams.get('limit'))
        const search = searchParams.get('search') || ''

        const filter: {
            [key: string]: {
                [key: string]: { $regex: string; $options: string }
            }[]
        } = {}

        if (!search) {
            filter['$or'] = []
            Object.keys(Lang).forEach((key) => {
                filter['$or'].push({
                    [`title.${key}`]: { $regex: search, $options: 'i' },
                })
            })
        }

        const { collection, error } = await getArticleDb()
        if (error instanceof Error) {
            return NewInternal(req, error)
        }

        const cur = collection!.find(filter, { limit: limit, skip: offset })
        const articles = await cur.toArray()
        return NewSuccess(req, articles)
    } catch (error: unknown) {
        console.error(error)
        if (typeof error === 'string') return NewError(req, new Error(error))
        else if (error instanceof Error) return NewError(req, error)
        else return NewError(req, new Error('Unknown error'))
    }
}
