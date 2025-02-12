import { Article } from '@domain/article/article'
import { Author } from '@domain/author/author'
import { Bio } from '@domain/bio/bio'
import { Book } from '@domain/book/book'
import { Contact } from '@domain/contact/contact'
import { Project } from '@domain/project/project'
import { Collection, MongoClient } from 'mongodb'

const ARTICLE_COLLECTION = 'articles'
const AUTHOR_COLLECTION = 'author'
const BIO_COLLECTION = 'bio'
const BOOK_COLLECTION = 'books'
const CONTACT_COLLECTION = 'contacts'
const PROJECT_COLLECTION = 'projects'

let MONGO_URI: string | null = null

let cachedClient: MongoClient | null = null

const cachedArticleColl: Collection<Article> | null = null
const cachedAuthorColl: Collection<Author> | null = null
const cachedBioColl: Collection<Bio> | null = null
const cachedBookColl: Collection<Book> | null = null
const cachedContactColl: Collection<Contact> | null = null
const cachedProjectColl: Collection<Project> | null = null

interface MongoConnect {
    client: MongoClient | null
    error: Error | null
}

export async function connectToDatabase(): Promise<MongoConnect> {
    try {
        if (!MONGO_URI) {
            MONGO_URI = process.env.MONGO_URI || ''
        }

        // check the MongoDB URI
        if (!MONGO_URI) {
            return {
                client: null,
                error: new Error('Define the MONGO_URI environmental variable'),
            }
        }

        // check the cached.
        if (cachedClient) {
            // load from cache
            return {
                client: cachedClient,
                error: null,
            }
        }

        // Connect to cluster
        const client: MongoClient = new MongoClient(MONGO_URI, {
            serverSelectionTimeoutMS: 100,
        })
        await client.connect()

        // set cache
        cachedClient = client

        return {
            client: cachedClient,
            error: null,
        }
    } catch (error: unknown) {
        console.error(error)
        return {
            client: null,
            error: new Error('Failed to connect to MongoDB'),
        }
    }
}

interface DbConnect<T> {
    collection: T | null
    error: Error | null
}

export async function getArticleDb(): Promise<
    DbConnect<typeof cachedArticleColl>
> {
    if (cachedArticleColl) {
        return { collection: cachedArticleColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(ARTICLE_COLLECTION),
        error: error,
    }
}

export async function getAuthorDb(): Promise<
    DbConnect<typeof cachedAuthorColl>
> {
    if (cachedAuthorColl) {
        return { collection: cachedAuthorColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(AUTHOR_COLLECTION),
        error: error,
    }
}

export async function getBioDb(): Promise<DbConnect<typeof cachedBioColl>> {
    if (cachedBioColl) {
        return { collection: cachedBioColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(BIO_COLLECTION),
        error: error,
    }
}

export async function getBookDb(): Promise<DbConnect<typeof cachedBookColl>> {
    if (cachedBookColl) {
        return { collection: cachedBookColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(BOOK_COLLECTION),
        error: error,
    }
}

export async function getContactDb(): Promise<
    DbConnect<typeof cachedContactColl>
> {
    if (cachedContactColl) {
        return { collection: cachedContactColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(CONTACT_COLLECTION),
        error: error,
    }
}

export async function getProjectDb(): Promise<
    DbConnect<typeof cachedProjectColl>
> {
    if (cachedProjectColl) {
        return { collection: cachedProjectColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return {
        collection: client!.db('main').collection(PROJECT_COLLECTION),
        error: error,
    }
}

export async function closeDatabase() {
    if (cachedClient) {
        await cachedClient.close()
    }
}
