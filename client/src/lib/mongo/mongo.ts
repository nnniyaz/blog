import { Collection, Db, MongoClient } from 'mongodb'

const ARTICLE_COLLECTION = 'articles'
const AUTHOR_COLLECTION = 'author'
const BIO_COLLECTION = 'bio'
const BOOK_COLLECTION = 'books'
const CONTACT_COLLECTION = 'contacts'
const PROJECT_COLLECTION = 'projects'
const SESSION_COLLECTION = 'sessions'
const USER_COLLECTION = 'users'

let MONGO_URI: string | null = null

let cachedClient: MongoClient | null = null

let cachedArticleColl: Collection | null = null
let cachedAuthorColl: Collection | null = null
let cachedBioColl: Collection | null = null
let cachedBookColl: Collection | null = null
let cachedContactColl: Collection | null = null
let cachedProjectColl: Collection | null = null
let cachedSessionColl: Collection | null = null
let cachedUserColl: Collection | null = null

interface MongoConnect {
    client: MongoClient | null
    error: Error | null
}

export async function connectToDatabase(): Promise<MongoConnect> {
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
    const client: MongoClient = new MongoClient(MONGO_URI)
    await client.connect()

    // set cache
    cachedClient = client

    return {
        client: cachedClient,
        error: null,
    }
}

interface DbConnect {
    collection: Collection | null
    error: Error | null
}

export async function getArticleDb(): Promise<DbConnect> {
    if (cachedArticleColl) {
        return { collection: cachedArticleColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(ARTICLE_COLLECTION), error: error }
}

export async function getAuthorDb(): Promise<DbConnect> {
    if (cachedAuthorColl) {
        return { collection: cachedAuthorColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(AUTHOR_COLLECTION), error: error }
}

export async function getBioDb(): Promise<DbConnect> {
    if (cachedBioColl) {
        return { collection: cachedBioColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(BIO_COLLECTION), error: error }
}

export async function getBookDb(): Promise<DbConnect> {
    if (cachedBookColl) {
        return { collection: cachedBookColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(BOOK_COLLECTION), error: error }
}

export async function getContactDb(): Promise<DbConnect> {
    if (cachedContactColl) {
        return { collection: cachedContactColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(CONTACT_COLLECTION), error: error }
}

export async function getProjectDb(): Promise<DbConnect> {
    if (cachedProjectColl) {
        return { collection: cachedProjectColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(PROJECT_COLLECTION), error: error }
}

export async function getSessionDb(): Promise<DbConnect> {
    if (cachedSessionColl) {
        return { collection: cachedSessionColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(SESSION_COLLECTION), error: error }
}

export async function getUserDb(): Promise<DbConnect> {
    if (cachedUserColl) {
        return { collection: cachedUserColl, error: null }
    }

    const { client, error } = await connectToDatabase()
    if (error instanceof Error) {
        return { collection: null, error: error }
    }
    return { collection: client!.db('main').collection(USER_COLLECTION), error: error }
}

export async function closeDatabase() {
    if (cachedClient) {
        await cachedClient.close()
    }
}
